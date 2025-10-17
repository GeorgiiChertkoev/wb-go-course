package builtins

import (
	"io"
	"os"
)

type baseCmd struct {
	stdin  io.Reader
	stdout io.Writer
	dir    string
	piped  bool
}

func (b *baseCmd) StdoutPipe() (io.ReadCloser, error) {
	// os.Pipe() was chosen over io.Pipe() because
	// if no one reads from pipe, writing goroutine
	// blocks forever resulting in deadlock
	pr, pw, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	b.piped = true
	b.stdout = pw
	return pr, nil
}

func (b *baseCmd) SetStdout(w io.Writer) {
	b.stdout = w
}

func (b *baseCmd) SetStdin(r io.Reader) {
	b.stdin = r
}

func (c *baseCmd) SetDir(path string) {
	c.dir = path
}

// should be not blocking
func (b *baseCmd) Start() error {
	return nil
}

func (b *baseCmd) Wait() error {
	if b.piped {
		w, ok := b.stdout.(*os.File)
		if ok {
			return w.Close()
		}
	}
	return nil
}
