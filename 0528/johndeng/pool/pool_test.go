package pool

import "testing"

func TestGetPool(t *testing.T) {

	poolLen := 10
	pool := NewStringPool(poolLen)
	d := pool.Get()

	if len(d) == 0 {
		t.Error("Unexpected data.")
	}

	if len(pool.active) != 1 {
		t.Error("Unexpected active list")
	}

	if len(pool.idle) != (poolLen - 1) {
		t.Error("Unexpected idle list")
	}

}

func TestBackPool(t *testing.T) {

	poolLen := 10
	pool := NewStringPool(poolLen)
	d := pool.Get()
	pool.Back(d)

	if len(d) == 0 {
		t.Error("Unexpected data.")
	}

	if len(pool.active) != 0 {
		t.Error("Unexpected active list")
	}

	if len(pool.idle) != poolLen {
		t.Error("Unexpected idle list")
	}

}

func BenchmarkGetPool(b *testing.B) {

	pool := NewStringPool(10)

	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkBackPool(b *testing.B) {

	pool := NewStringPool(10)

	for i := 0; i < b.N; i++ {
		pool.Back(genString())
	}
}
