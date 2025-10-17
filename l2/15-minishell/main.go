package main

import (
	"minishell/shell"
)

func main() {
	// ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	// defer cancel()
	// s := shell.NewShellWithContext(ctx, "C:/")
	// var wg sync.WaitGroup
	// wg.Go(s.Start)
	// wg.Wait()
	s := shell.NewShell("C:\\Users\\geoch\\Downloads\\Telegram Desktop")
	s.Start()
}
