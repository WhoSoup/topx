package topx

import (
	"testing"
)

func TestTopX_Add(t *testing.T) {
	top := New(5)

	for i := uint64(0); i < 500; i++ {
		top.Add(i)
	}

	r := top.Get()

	for i := uint64(0); i < 5; i++ {
		if r[i] != 499-i {
			t.Errorf("bad data at %d. got = %d, want = %d", i, r[i], 499-i)
		}
	}
}
