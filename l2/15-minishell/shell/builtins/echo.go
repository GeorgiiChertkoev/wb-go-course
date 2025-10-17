package builtins

import (
	"context"
	"fmt"
)

type Echo struct {
	baseCmd
	args     []string
	ctx      context.Context
	doneFunc context.CancelFunc
}

func NewEcho(ctx context.Context, args ...string) *Echo {
	newCtx, done := context.WithCancel(ctx)
	return &Echo{
		args:     args,
		ctx:      newCtx,
		doneFunc: done,
	}
}

func (cd *Echo) Start() error {
	go func() {
		defer cd.doneFunc()
		for _, arg := range cd.args {
			fmt.Fprintln(cd.stdout, arg) // behaves similar to windows
		}
	}()

	return nil
}

func (cd *Echo) Wait() error {
	<-cd.ctx.Done()
	return cd.baseCmd.Wait()
}
