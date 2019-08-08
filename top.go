package topx

import "sync"

type TopX struct {
	mtx   sync.Mutex
	Count int
	items []uint64
}

func New(count int) *TopX {
	t := new(TopX)
	t.Count = count
	return t
}

func (t *TopX) Add(n uint64) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	for i := range t.items {
		if n > t.items[i] {
			t.items = append(t.items, 0)
			copy(t.items[i+1:], t.items[i:])
			t.items[i] = n
			if len(t.items) > t.Count {
				t.items = t.items[:t.Count]
			}
			return
		}
	}

	if len(t.items) < t.Count {
		t.items = append(t.items, n)
	}
}

func (t *TopX) Get() []uint64 {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	cpy := make([]uint64, len(t.items))
	copy(cpy, t.items)
	return cpy
}
