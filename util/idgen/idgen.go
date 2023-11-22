package idgen

type IdGenerator interface {
	Allocate() uint64
	Free(uint64)
}

type idGenerator struct {
	cur, lb, ub uint64
	values      map[uint64]bool //returned identities
}

func NewIdGenerator(lb, ub uint64) IdGenerator {
	return &idGenerator{
		lb:     lb,
		ub:     ub,
		values: make(map[uint64]bool),
		cur:    0,
	}
}

func (gen *idGenerator) Allocate() (id uint64) {
	if l := len(gen.values); l > 0 {
		for id, _ = range gen.values {
			delete(gen.values, id)
			break
		}
	} else {
		id = gen.cur
		if gen.cur++; gen.cur >= gen.ub {
			gen.cur = gen.lb
		}
	}
	return
}

func (gen *idGenerator) Free(id uint64) {
	gen.values[id] = true
}
