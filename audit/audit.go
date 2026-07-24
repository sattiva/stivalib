package audit

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

type Entry struct {
	Index     uint64 `json:"index"`
	Timestamp string `json:"timestamp"`
	Actor     string `json:"actor"`
	Action    string `json:"action"`
	PrevHash  string `json:"prev_hash"`
	Hash      string `json:"hash"`
}

type Chain struct {
	mu     sync.RWMutex
	entries []Entry
}

func NewChain() *Chain {
	c := &Chain{entries: make([]Entry, 0)}
	c.Append("SYSTEM", "GENESIS_BLOCK")
	return c
}

func (c *Chain) Append(actor, action string) Entry {
	c.mu.Lock()
	defer c.mu.Unlock()

	idx := uint64(len(c.entries))
	prev := "0000000000000000000000000000000000000000000000000000000000000000"
	if idx > 0 {
		prev = c.entries[idx-1].Hash
	}

	ts := time.Now().UTC().Format(time.RFC3339Nano)
	raw := fmt.Sprintf("%d:%s:%s:%s:%s", idx, ts, actor, action, prev)
	h := sha256.Sum256([]byte(raw))

	e := Entry{
		Index:     idx,
		Timestamp: ts,
		Actor:     actor,
		Action:    action,
		PrevHash:  prev,
		Hash:      hex.EncodeToString(h[:]),
	}
	c.entries = append(c.entries, e)
	return e
}

func (c *Chain) Verify() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for i := 1; i < len(c.entries); i++ {
		curr := c.entries[i]
		prev := c.entries[i-1]

		if curr.PrevHash != prev.Hash {
			return false
		}

		raw := fmt.Sprintf("%d:%s:%s:%s:%s", curr.Index, curr.Timestamp, curr.Actor, curr.Action, curr.PrevHash)
		h := sha256.Sum256([]byte(raw))
		if hex.EncodeToString(h[:]) != curr.Hash {
			return false
		}
	}
	return true
}
