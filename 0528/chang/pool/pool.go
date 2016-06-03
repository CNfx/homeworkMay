package pool

import (
	"errors"
	"sync"
)

var (
	ErrPoolFull  = errors.New("pool full")
	ErrPoolEmpty = errors.New("pool empty")
)

type Object struct {
	Data int
}

type Pool struct {
	data   []*Object
	header int
	lock   sync.Mutex
}

func NewPool(cap int) *Pool {
	pool := &Pool{
		data: make([]*Object, cap),
	}

	pool.header = 0
	return pool
}

func (p *Pool) Push(obj *Object) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.header == 0 {
		return ErrPoolFull
	}

	p.data[p.header-1] = obj
	p.header--
	return nil
}

func (p *Pool) Pop() (*Object, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.header == cap(p.data) {
		return nil, ErrPoolEmpty
	}

	v := p.data[p.header]
	p.header++
	return v, nil
}
