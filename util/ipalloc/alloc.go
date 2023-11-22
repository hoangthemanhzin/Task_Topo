package ipalloc

import (
	"net"
)

type idPool struct {
	l    int64
	u    int64
	used map[int64]bool
}

func newIdPool(l int64, u int64) idPool {
	return idPool{
		l:    l,
		u:    u,
		used: make(map[int64]bool),
	}
}

func (p *idPool) allocate() (id int64) {
	for id = p.l; id <= p.u; id++ {
		if _, ok := p.used[id]; !ok {
			p.used[id] = true
			return
		}
	}
	id = p.l - 1
	return
}

func (i *idPool) release(id int64) {
	delete(i.used, id)
}

type IpAllocator struct {
	cidr *net.IPNet
	pool idPool
}

func New(cidr *net.IPNet) *IpAllocator {
	//calculate number of mask bits
	var bits int
	for _, b := range cidr.Mask {
		for ; b != 0; b /= 2 {
			if b%2 != 0 {
				bits++
			}
		}
	}

	return &IpAllocator{
		cidr: cidr,
		pool: newIdPool(1, 1<<int64(bits)-2),
	}
}

// Allocate will allocate the IP address and returns it
func (a *IpAllocator) Allocate() (ip net.IP) {
	if id := a.pool.allocate(); id < a.pool.l {
		return
	} else {
		return offsetIp(a.cidr.IP, int(id))
	}
}

func (a *IpAllocator) Release(ip net.IP) {
	id := ipOffset(ip, a.cidr.IP)
	a.pool.release(int64(id))
}

// return new Ip by offseting a base Ip
func offsetIp(base net.IP, offset int) (ip net.IP) {
	ip = make(net.IP, len(base))
	copy(ip, base)

	var carry int
	for i := len(ip) - 1; i >= 0; i-- {
		if offset == 0 {
			break
		}

		val := int(ip[i]) + carry + offset%256
		ip[i] = byte(val % 256)
		carry = val / 256

		offset /= 256
	}

	return
}

// difference between 2 Ip addresses
func ipOffset(in, base net.IP) (offset int) {
	exp := 1
	for i := len(base) - 1; i >= 0; i-- {
		offset += int(in[i]-base[i]) * exp
		exp *= 256
	}
	return
}
