package context_cancel

import (
	"context"
	"fmt"
	"time"
)

func PrintElaspedTimeUntilCancellation(ctx context.Context, routineName string) {
	start := time.Now()
	elasped := 0 * time.Second

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] Print elasped time cancelled\n", routineName)
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
