package infrastructure

import (
	"context"
	"log"
	"time"
)

func StartCrontabWorker(
	ctx context.Context,
	processor *CrontabProcessor,
) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	// log.Println("Crontab started")

	for {
		select {
		case <-ctx.Done():
			log.Println("Crontab stopped")
			return

		case <-ticker.C:
			if err := processor.Process(ctx); err != nil {
				log.Println("Crontab error:", err)
			}
		}
	}
}
