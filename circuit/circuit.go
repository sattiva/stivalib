package circuit

import (
	"errors"
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	HalfOpen
	Open
)

type Breaker struct {
	mu       sync.Mutex
	state    State
	fails    int
	thresh   int
	cooldown time.Duration
	lastFail time.Time
}

func New(threshold int, cooldown time.Duration) *Breaker {
	return &Breaker{thresh: threshold, cooldown: cooldown}
}

func (b *Breaker) Execute(fn func() error) error {
	b.mu.Lock()
	if b.state == Open {
		if time.Since(b.lastFail) > b.cooldown {
			b.state = HalfOpen
		} else {
			b.mu.Unlock()
			return errors.New("circuit open")
		}
	}
	b.mu.Unlock()

	err := fn()

	b.mu.Lock()
	defer b.mu.Unlock()
	if err != nil {
		b.fails++
		b.lastFail = time.Now()
		if b.fails >= b.thresh {
			b.state = Open
		}
		return err
	}

	b.fails = 0
	b.state = Closed
	return nil
}
