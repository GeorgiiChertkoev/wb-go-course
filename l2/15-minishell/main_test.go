package main_test

import (
	"bytes"
	"context"
	"minishell/shell"
	"runtime"
	"strings"
	"testing"
	"time"
)

func runShellInput(t *testing.T, input string) (stdout, stderr string) {
	t.Helper()

	in := bytes.NewBufferString(input)
	out := &bytes.Buffer{}
	errout := &bytes.Buffer{}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	s := shell.NewShellWithContext(ctx, "")
	s.Stdin = in
	s.Stdout = out
	s.Stderr = errout

	s.Start()

	return strings.TrimSpace(out.String()), strings.TrimSpace(errout.String())
}

func TestEcho(t *testing.T) {
	stdout, _ := runShellInput(t, "echo hello world\n")
	if stdout != "hello world" {
		t.Errorf("expected 'hello world', got '%s'", stdout)
	}
}

func TestPwdCd(t *testing.T) {
	tmp := t.TempDir()
	tmp2 := t.TempDir()
	stdout, _ := runShellInput(t, "cd "+tmp+"\npwd\n"+"cd "+tmp2+"\npwd\n")
	lines := strings.Split(stdout, "\n")
	if len(lines) < 2 {
		t.Fatalf("expected 2 lines, got %v", lines)
	}
	if strings.TrimSpace(lines[1]) != tmp2 {
		t.Errorf("expected dir=%s, got=%s", tmp, lines[1])
	}
}

func TestPipeline(t *testing.T) {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = `echo foo bar baz | findstr ba | find /C /V ""`
	} else {
		cmd = "echo foo bar baz | grep ba | wc -l\n"
	}

	stdout, _ := runShellInput(t, cmd)
	stdout = strings.TrimSpace(stdout)
	if stdout != "1" {
		t.Errorf("expected '1', got '%s'", stdout)
	}
}

func TestPs(t *testing.T) {
	stdout, _ := runShellInput(t, "ps\n")
	if !strings.Contains(stdout, "PID") && !strings.Contains(stdout, "COMMAND") {
		t.Errorf("expected ps output, got:\n%s", stdout)
	}
}
