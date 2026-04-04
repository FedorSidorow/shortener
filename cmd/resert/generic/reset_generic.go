package generic

type Resetter interface {
	Reset()
}

type Pool[T Resetter] struct {
	item T
}

func New[T Resetter](item T) *Pool[T] {
	return &Pool[T]{item: item}
}

func (p *Pool[T]) Get() T {
	return p.item
}

func (p *Pool[T]) Put(item T) {
	item.Reset()
}
