package pool

import (
	"sync"
)

type BytePool struct {
	pool sync.Pool
	size int
}

func NewBytePool(bufferSize int) *BytePool {
	return &BytePool{
		size: bufferSize,
		pool: sync.Pool{
			New: func() interface{} {
				b := make([]byte, bufferSize)
				return &b
			},
		},
	}
}

func (p *BytePool) Get() *[]byte {
	return p.pool.Get().(*[]byte)
}

func (p *BytePool) Put(b *[]byte) {
	if b == nil || cap(*b) < p.size {
		return
	}
	*b = (*b)[:p.size]
	p.pool.Put(b)
}
