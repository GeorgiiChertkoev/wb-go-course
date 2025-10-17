package builtins

import (
	"errors"
	"fmt"
	"os"
)

type ChangeDir struct {
	baseCmd
	err error
}

func NewChangeDir(shellDir *string, args ...string) *ChangeDir {
	var err error
	if len(args) != 1 {
		err = errors.New("cd takes exactly 1 argument")
		return &ChangeDir{
			err: err,
		}
	}

	stat, err := os.Stat(args[0])
	if err != nil || !stat.IsDir() {
		err = fmt.Errorf("cd argument is not a valid directory: %s", args[0])
		return &ChangeDir{
			err: err,
		}
	}
	*shellDir = args[0]

	return &ChangeDir{
		err: err,
	}
}

func (cd *ChangeDir) Start() error {
	return cd.err
}
