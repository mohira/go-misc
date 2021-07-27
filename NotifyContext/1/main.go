package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Graceful Shutdown
// https://play.golang.org/p/S-h80RaQ4UM
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second) // 重たい処理のシミュレーション
		fmt.Fprintln(w, "done")
	})

	// http.ListenAndServe() をゴルーチンで動かす
	go func() {
		fmt.Printf("server is running http://localhost%s\n", server.Addr)

		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()

	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
