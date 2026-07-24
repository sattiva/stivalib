package ratelimit

import (
	"sync"
	"time"
)

type Limiter struct {
	mu     sync.Mutex
	v      map[string]*bucket
	rate   float64
	burst  float64
	ttl    time.Duration
}

type bucket struct {
	toks float64
	last time.Time
}

func New(rate, burst float64, ttl time.Duration) *Limiter {
	l := &Limiter{
		v:     make(map[string]*bucket),
		rate:  rate,
		burst: burst,
		ttl:   ttl,
	}
	go l.cleanup()
	return l
}

func (l *Limiter) Allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	b, ok := l.v[key]
	if !ok {
		l.v[key] = &bucket{toks: l.burst - 1, last: now}
		return true
	}

	el := now.Sub(b.last).Seconds()
	b.last = now
	b.toks += el * l.rate
	if b.toks > l.burst {
		b.toks = l.burst
	}

	if b.toks >= 1 {
		b.toks -= 1
		return true
	}
	return false
}

func (l *Limiter) cleanup() {
	for range time.Tick(l.ttl) {
		l.mu.Lock()
		now := time.Now()
		for k, b := range l.v {
			if now.Sub(b.last) > l.ttl {
				delete(l.v, k)
			}
		}
		l.mu.Unlock()
	}
}
