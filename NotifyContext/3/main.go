package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

// ゴルーチンで5つのリクエストを送る
func main() {
	var wg sync.WaitGroup

	repeat := 5

	wg.Add(repeat)

	for i := 0; i < repeat; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Request%d\n", n)

			url := "http://localhost:8080"
			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close()

			io.Copy(os.Stdout, resp.Body)
		}(i)
	}

	wg.Wait()
	fmt.Println("Finish")
}
