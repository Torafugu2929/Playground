package context_cancel

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func PrintElaspedTimeUntilCancellation(ctx context.Context, routineName string) {
	start := time.Now()
	elasped := 0 * time.Second

	if deadline, ok := ctx.Deadline(); ok {
		fmt.Printf("[%s] Deadline: %s\n", routineName, deadline.Format(time.RFC3339))
	}

	for {
		select {
		// 親ルーチンによってcancel()が呼び出されたときに入るブロック。
		// select文は対象のチャネルに値を受信したかどうかを判定する
		case <-ctx.Done():
			err := ctx.Err()
			if errors.Is(err, context.Canceled) {
				fmt.Printf("[%s] Cancelled by intention\n", routineName)
			} else if errors.Is(err, context.DeadlineExceeded) {
				fmt.Printf("[%s] Cancelled by timeout\n", routineName)
			} else {
				fmt.Printf("[%s] Cancelled by unknown reason\n", routineName)
			}

			return

		default:
			now := time.Now()
			if now.Sub(start) > elasped+time.Second {
				fmt.Printf("[%s] Elasped time: %d[s]\n", routineName, elasped/time.Second)
				time.Sleep(time.Second)
				elasped += time.Second
			}

		}
	}
}
