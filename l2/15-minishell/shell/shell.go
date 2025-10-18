package shell

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"minishell/shell/builtins"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
)

const exitCharacters = "\x04" // ctrl+D
const commandsBufferSize = 128

type Shell struct {
	Ctx            context.Context
	Dir            string
	Stdin          io.Reader
	Stdout         io.Writer
	Stderr         io.Writer
	commands       chan string
	cancelCommand  chan os.Signal
	shellCtxCancel context.CancelFunc
}

func NewShell(Dir string) *Shell {
	ctx, shellCtxCancel := context.WithCancel(context.Background())
	return &Shell{
		Ctx:            ctx,
		Dir:            Dir,
		Stdin:          os.Stdin,
		Stdout:         os.Stdout,
		Stderr:         os.Stderr,
		shellCtxCancel: shellCtxCancel,
	}
}

func NewShellWithContext(ctx context.Context, Dir string) *Shell {
	ctx, shellCtxCancel := context.WithCancel(ctx)
	return &Shell{
		Ctx:            ctx,
		Dir:            Dir,
		shellCtxCancel: shellCtxCancel,
	}
}

func (s *Shell) Start() {
	s.commands = make(chan string, commandsBufferSize)
	s.cancelCommand = make(chan os.Signal, 1)
	defer close(s.commands)
	defer close(s.cancelCommand)

	signal.Notify(s.cancelCommand, os.Interrupt)

	go s.consumeCommands()

	for {
		select {
		case <-s.Ctx.Done():
			return
		default:
		}
		scanner := bufio.NewScanner(s.Stdin)
		for scanner.Scan() {
			cmd := scanner.Text()
			if strings.ContainsAny(cmd, exitCharacters) || strings.TrimSpace(cmd) == "exit" { // extra way to exit
				s.shellCtxCancel()
				return
			}
			cmd = strings.TrimSpace(cmd)
			if cmd != "" {
				s.commands <- cmd
			}
		}
	}
}

// needed to be able to execute and read cmds separately
// to be able to cancel with ctrl+D
func (s *Shell) consumeCommands() {
	for cmd := range s.commands {
		select {
		case <-s.Ctx.Done():
			return
		case <-s.cancelCommand:
			// waste multiple cancels
		default:
		}
		ctx, cancel := context.WithCancel(s.Ctx)
		go func() {
			select {
			case <-ctx.Done():
			case <-s.cancelCommand:
				fmt.Printf("got cancel, stop \"%s\"\n", cmd)
				cancel()
			}
		}()
		err := s.execute(ctx, cmd)
		if err != nil {

			fmt.Fprintf(os.Stderr, "failed to execute %s with err: %v\n", cmd, err)
			fmt.Fprintf(s.Stderr, "failed to execute %s with err: %v\n", cmd, err)
		}
		cancel()
	}
}

// parses pipes and then calls runCommand
func (s *Shell) execute(ctx context.Context, commandLine string) (err error) {
	commands := strings.Split(commandLine, "|")

	cmds := make([]Command, 0)
	var prevStdout io.Reader = s.Stdin
	for i, command := range commands {
		cmd := s.createCmd(ctx, command, prevStdout)
		cmds = append(cmds, cmd)
		if i != len(commands)-1 {
			prevStdout, err = cmd.StdoutPipe()
		}
		if err != nil {
			return err
		}
	}
	cmds[len(cmds)-1].SetStdout(s.Stdout)

	for _, cmd := range cmds {
		cmd.SetDir(s.Dir)
		if err = cmd.Start(); err != nil {
			return err
		}
	}

	for i, cmd := range cmds {
		if err = cmd.Wait(); err != nil {
			return fmt.Errorf("error executing %s: %v", commands[i], err)
		}
	}
	return nil
}

// transforms command into either exec.Cmd or built-in command
// interface is used to hide difference and be able to use built-in as exec.Cmd
func (s *Shell) createCmd(ctx context.Context, command string, stdin io.Reader) Command {
	fields := strings.Fields(command)
	switch fields[0] {
	case "cd":
		return builtins.NewChangeDir(&s.Dir, fields[1:]...)
	case "pwd":
		return builtins.NewPrintWorkingDir(ctx, fields[1:]...)
	case "echo":
		return builtins.NewEcho(ctx, fields[1:]...)
	case "kill":
		return builtins.NewKillCmd(ctx, fields[1:]...)
	case "ps":
		return builtins.NewProcesses(ctx, fields[1:]...)
	}

	if runtime.GOOS == "windows" {
		fields = append([]string{"cmd", "/C"}, fields...)
	}
	cmd := exec.CommandContext(ctx, fields[0], fields[1:]...)
	cmd.Stdin = stdin

	// needed because all piped commands print in same
	// io.Writer causing data race
	cmd.Stderr = &syncWriter{w: s.Stderr}

	return &cmdWrapper{
		Cmd: cmd,
	}
}
