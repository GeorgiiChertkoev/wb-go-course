package builtins

import (
	"context"
	"errors"
	"fmt"
)

type PrintWorkingDir struct {
	baseCmd
	args     []string
	ctx      context.Context
	doneFunc context.CancelFunc
}

func NewPrintWorkingDir(ctx context.Context, args ...string) *PrintWorkingDir {
	newCtx, done := context.WithCancel(ctx)
	return &PrintWorkingDir{
		args:     args,
		ctx:      newCtx,
		doneFunc: done,
	}
}

func (cd *PrintWorkingDir) Start() error {
	if len(cd.args) != 0 {
		return errors.New("pwd takes exactly 0 arguments")
	}

	go func() {
		defer cd.doneFunc()
		fmt.Fprintln(cd.stdout, cd.dir)
	}()

	return nil
}

func (cd *PrintWorkingDir) Wait() error {
	<-cd.ctx.Done()
	return cd.baseCmd.Wait()
}
