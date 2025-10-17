package builtins

import (
	"context"
	"io"
	"os/exec"
	"runtime"
)

type Processes struct {
	*exec.Cmd
}

func NewProcesses(ctx context.Context, args ...string) *Processes {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin": // macOS
		cmd = exec.CommandContext(ctx, "ps", "-e", "-o", "pid,comm")
	case "windows":
		cmd = exec.CommandContext(ctx, "tasklist")
	default:
		cmd = exec.CommandContext(ctx, "ps")
	}

	return &Processes{
		Cmd: cmd,
	}
}

func (ps *Processes) SetDir(path string) {
	ps.Cmd.Dir = path
}

func (ps *Processes) SetStdout(w io.Writer) {
	ps.Stdout = w
}

func (ps *Processes) SetStdin(r io.Reader) {
	ps.Stdin = r
}
