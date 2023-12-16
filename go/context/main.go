package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	context_cancel "github.com/Torafugu2929/playground/go/context/cancel"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// child
	go context_cancel.PrintElaspedTimeUntilCancellation(ctx, "child")

	wg.Add(1)

	<-quit

	fmt.Println("[parent] Ctrl+C pressed")
	cancel()
	time.Sleep(1 * time.Second)
}
