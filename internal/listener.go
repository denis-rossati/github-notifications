package internal

import (
	"context"
	"fmt"
	"github.com/gen2brain/beeep"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type lastUpdate struct {
	mu   sync.Mutex
	time *time.Time
}

var lastUpdateTime lastUpdate

var sem = make(chan struct{}, 5)

func Listen(ctx context.Context, token string) {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	log.Println("Listening for notifications")

	for {
		select {
		case <-ticker.C:
			log.Println("Requesting notifications")
			lastUpdateTime.mu.Lock()
			since := lastUpdateTime.time
			lastUpdateTime.mu.Unlock()

			select {
			case sem <- struct{}{}:
				go func() {
					defer func() { <-sem }()
					processNotifications(ctx, token, since)
				}()
			default:
				log.Println("Skipping notification check")
			}
		case <-ctx.Done():
			log.Println("Listener stopping due to context cancellation")
			return
		}
	}
}

func processNotifications(ctx context.Context, token string, since *time.Time) {
	notifications, err := getNotifications(ctx, token, since)

	if err != nil {
		log.Printf("Error getting notifications: %v", err)
	}

	if len(notifications) == 0 {
		return
	}

	lastUpdateTime.mu.Lock()
	defer lastUpdateTime.mu.Unlock()

	for i, n := range notifications {
		if i == 0 {
			earliestNotification, err := time.Parse("2006-01-02T15:04:05Z07:00", n.UpdatedAt)

			if err == nil && (lastUpdateTime.time == nil || lastUpdateTime.time.Before(earliestNotification)) {
				// Add one second to prevent getting the same latest notification over and over
				t := earliestNotification.Add(time.Second * 1)

				lastUpdateTime.time = &t
			}
		}

		subtitle := fmt.Sprintf(
			"%s at %s",
			n.Reason,
			n.Repository.FullName,
		)

		execPath, err := os.Executable()

		if err != nil {
			log.Printf("Error getting executable path: %v", err)
			execPath = "."
		}

		dir := filepath.Dir(execPath)
		p := filepath.Join(dir, "assets", "github.png")

		err = beeep.Notify(
			n.Subject.Title,
			subtitle,
			p,
		)

		if err != nil {
			log.Printf("Error sending notification: %s", err)
		}
	}
}
