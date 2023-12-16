package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	context_cancel "github.com/Torafugu2929/playground/go/context/cancel"
)

func main() {
	timeout := 10
	ctxInterrupt, cancelInterrupt := context.WithCancel(context.Background())
	ctxTimeout, cancelTimeout := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(timeout)*time.Second))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// child
	go context_cancel.PrintElaspedTimeUntilCancellation(ctxInterrupt, "childInterrupt")
	go context_cancel.PrintElaspedTimeUntilCancellation(ctxTimeout, "childTimeout")

	<-quit

	fmt.Println("[parent] Ctrl+C pressed")
	cancelInterrupt()
	cancelTimeout()
	time.Sleep(1 * time.Second)
}
