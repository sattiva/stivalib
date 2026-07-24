package retry

import (
	"math/rand"
	"time"
)

func Do(fn func() error, maxRetries int, delay time.Duration) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		if err = fn(); err == nil {
			return nil
		}
		jitter := time.Duration(rand.Int63n(int64(delay / 2)))
		time.Sleep(delay + jitter)
		delay *= 2
	}
	return err
}
