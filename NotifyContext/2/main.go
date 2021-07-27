package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

// https://pkg.go.dev/os/signal#example-NotifyContext
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Ctrl+C で SIGINT を送信する → ctxがキャンセルされる
	// なにもせず3秒経過する → <-time.After のcaseが採用される
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("シグナルは送信されませんでしたよ")
	case <-ctx.Done():
		fmt.Println()
		fmt.Println(ctx.Err())
		stop()
	}
}
