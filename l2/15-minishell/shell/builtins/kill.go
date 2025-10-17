package builtins

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type KillCmd struct {
	baseCmd
	args     []string
	ctx      context.Context
	doneFunc context.CancelFunc
}

func NewKillCmd(ctx context.Context, args ...string) *KillCmd {
	newCtx, done := context.WithCancel(ctx)
	return &KillCmd{
		args:     args,
		ctx:      newCtx,
		doneFunc: done,
	}
}

func (cd *KillCmd) Start() error {
	if len(cd.args) != 1 {
		return errors.New("kill takes exactly 1 argument")
	}

	n, err := strconv.Atoi(cd.args[0])
	if err != nil {
		return fmt.Errorf("kill argument should be a number. got: %s, err: %v", cd.args[0], err)
	}

	go func() {
		defer cd.doneFunc()
		p, err := os.FindProcess(n)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error finding process with pid %d: %v", n, err)
		}
		err = p.Kill()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error killin process with pid %d: %v", n, err)
		}
	}()

	return nil
}

func (cd *KillCmd) Wait() error {
	<-cd.ctx.Done()
	return cd.baseCmd.Wait()
}
