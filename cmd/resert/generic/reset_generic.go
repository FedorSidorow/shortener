package generic

import (
	"sync"
)

type Resetter interface {
	Reset()
}

type Pool[T Resetter] struct {
	items []T
	mu    sync.Mutex
}

func New[T Resetter](item T) *Pool[T] {
	p := &Pool[T]{}
	// Always store the initial item
	p.items = append(p.items, item)
	return p
}

func (p *Pool[T]) Get() T {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.items) == 0 {
		// Return zero value
		var zero T
		return zero
	}

	// Get last item (LIFO like sync.Pool)
	item := p.items[len(p.items)-1]
	p.items = p.items[:len(p.items)-1]
	return item
}

func (p *Pool[T]) Put(item T) {
	item.Reset()

	p.mu.Lock()
	defer p.mu.Unlock()
	p.items = append(p.items, item)
}
