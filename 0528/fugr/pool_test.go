package structure

import (
	"testing"
	"unsafe"
)

func testCachePool(c Cache, convert func(unsafe.Pointer) Cache, t *testing.T) {
	var (
		first, last unsafe.Pointer
		pool        = NewPool(c, 1<<11)
	)
	for {
		c1, err := pool.Get(convert)
		if err != nil {
			t.Log(err)
			break
		}

		first = c1.Pointer()
		t.Logf("%p", first)

		if last != nil && first != last {
			t.Errorf("Unexpected,address from pool isnot the last one,%p!=%p", first, last)
		}

		c2, err := pool.Get(convert)
		if err != nil {
			t.Log(err)
			break
		}

		last = c2.Pointer()
		t.Logf("%p", last)

		if first == last {
			t.Errorf("Unexpected,got same address from pool,%p==%p", first, last)
		}

		// put back last to pool
		pool.Put(c2)
	}
}

func TestInts(t *testing.T) {
	a := Ints{}
	for i := range a {
		a[i] = i
	}
	a.Reset()
	for i := range a {
		if a[i] != 0 {
			t.Error(i, a[i], "!= 0")
		}
	}

	if sizeof := a.Sizeof(); sizeof != len(a)*8 {
		t.Errorf("Sizeof:want %d got %d", sizeof, len(a)*8)
	}

	testCachePool(&a, IntsConvert, t)
}

func TestComposite(t *testing.T) {
	i := 0x1000
	c := Composite{
		a:     0x1F,
		b:     0xEF0000FF,
		s:     "abcdefg",
		p:     &i,
		m:     make(map[string]string, 10),
		slice: make([]int, 10, 20),
		array: [size]int{1, 2, 3, 4, 5},
	}

	c.Reset()

	if c.a != 0 ||
		c.b != 0 ||
		c.s != "" ||
		c.p != nil ||
		c.m != nil ||
		c.slice != nil {
		t.Errorf("Unexpected,got %+v", c)
	}
	for i := range c.array {
		if c.array[i] != 0 {
			t.Error(i, c.array[i], "!= 0")
		}
	}

	if sizeof := c.Sizeof(); sizeof != 256 {
		t.Errorf("Sizeof:want %d got %d", 256, sizeof)
	}

	testCachePool(&c, CompositeConvert, t)
}
