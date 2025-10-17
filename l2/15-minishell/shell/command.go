package shell

import (
	"io"
	"os/exec"
)

// wrap up over exec.Cmd to add builtins
type Command interface {
	StdoutPipe() (io.ReadCloser, error)
	SetStdout(w io.Writer)
	SetStdin(r io.Reader)
	SetDir(path string)
	Start() error
	Wait() error
}

type cmdWrapper struct {
	*exec.Cmd
}

func (c *cmdWrapper) SetStdout(w io.Writer) {
	c.Stdout = w
}

func (c *cmdWrapper) SetStdin(r io.Reader) {
	c.Stdin = r
}

func (c *cmdWrapper) SetDir(path string) {
	c.Dir = path
}
