package structure

import (
	"errors"
	"unsafe"
)

var ErrPoolEmpty = errors.New("empty pool")

type Cache interface {
	Pointer() unsafe.Pointer
	// sizeof memory
	Sizeof() int
	// reset all elems tobe zero value
	Reset()
}

type Pool struct {
	head  unsafe.Pointer
	pools []byte
}

type list struct {
	next unsafe.Pointer
}

func NewPool(c Cache, cap int) *Pool {
	p := &Pool{
		pools: make([]byte, cap),
	}

	p.head = unsafe.Pointer(&p.pools[0])
	size := c.Sizeof()

	for i, cap := 0, len(p.pools); ; {
		current := (*list)(unsafe.Pointer(&p.pools[i]))
		next := i + size

		if next < cap {
			current.next = unsafe.Pointer(&p.pools[next])
		} else {
			current.next = nil
			break
		}
		i = next
	}

	return p
}

func (p *Pool) Get(convert func(unsafe.Pointer) Cache) (Cache, error) {
	if p.head == nil {
		return nil, ErrPoolEmpty
	}

	ptr := p.head

	node := (*list)(ptr)
	p.head = node.next

	return convert(ptr), nil
}

func (p *Pool) Put(c Cache) {
	c.Reset()

	head := p.head
	pointer := c.Pointer()
	p.head = pointer
	node := (*list)(pointer)
	node.next = head
}

const size = 1<<4 + 7

var (
	IntsConvert = func(ptr unsafe.Pointer) Cache {
		return (*Ints)(ptr)
	}

	CompositeConvert = func(ptr unsafe.Pointer) Cache {
		return (*Composite)(ptr)
	}
)

type Ints [size]int

func (a Ints) Sizeof() int {
	return int(unsafe.Sizeof(a))
}

func (a *Ints) Reset() {
	for i := range a {
		a[i] ^= a[i]
	}
}

func (a *Ints) Pointer() unsafe.Pointer {
	return unsafe.Pointer(a)
}

type Composite struct {
	a     byte
	b     int
	s     string
	p     *int
	m     map[string]string
	slice []int
	array [size]int
}

func (c Composite) Sizeof() int {
	return int(unsafe.Sizeof(c))
}

func (c *Composite) Reset() {
	c.a ^= c.a
	c.b ^= c.b
	for i := range c.array {
		c.array[i] ^= c.array[i]
	}
	c.s = ""
	c.p = nil
	c.slice = nil
	c.m = nil
}

func (c *Composite) Pointer() unsafe.Pointer {
	return unsafe.Pointer(c)
}
