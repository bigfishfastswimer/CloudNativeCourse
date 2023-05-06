package main

import (
	"context"
	"fmt"
	"time"
)

func publisher(ctx context.Context, ch chan int,  n int) {
	for i := 0; i<n; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Stop Producer due to main goroutine completed or context timeout")
			return
		default:
			ch <- i
			time.Sleep(time.Second)
		}
	}

}

func consumer(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop consumer due to main goroutine completed or context timeout")
			return
		case i:=<-ch:
			fmt.Printf("processed msg %d\n", i)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	queue := make(chan int, 10)
	go publisher(ctx, queue, 10)
	consumer(ctx, queue)
	fmt.Println("program finished")
}
