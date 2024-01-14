package contract

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
)

var (
	cmId = [4]byte{0x62, 0x4f, 0x82, 0xf5}
	p2   = [32]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x60,
	}
	p3 = [32]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x41}
	sfx = [31]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
)

type CMBuffer [228]byte

func (b *CMBuffer) Fill(arg0, arg1 uint32, arg2 string) ([]byte, error) {
	id := b[:4]
	pLen := 32
	copy(id, cmId[:])
	p0 := b[4 : pLen+4]
	p0 = p0[pLen-4:]
	binary.BigEndian.PutUint32(p0, arg0)
	p1 := b[pLen+4 : pLen+36]
	p1 = p1[pLen-4:]
	binary.BigEndian.PutUint32(p1, arg1)
	_p2 := b[pLen+36 : pLen*2+36]
	copy(_p2, p2[:])
	_p3 := b[pLen*2+36 : pLen*3+36]
	copy(_p3, p3[:])
	sign := b[pLen*3+36 : pLen*3+36+65]
	hb, err := hex.DecodeString(strings.TrimPrefix(arg2, "0x"))
	if err != nil {
		return nil, err
	}
	copy(sign, hb)
	_sfx := b[pLen*3+36+65 : pLen*3+36+65+31]
	copy(_sfx, sfx[:])
	return b[:], nil
}

// CKBuffer qna3 check-in hex-data buffer
var (
	CKBuffer = NewPool(CheckinMethod)
	CLBuffer = cmPool{p: sync.Pool{New: func() any {
		return new(CMBuffer)
	}}}
)

type Pool struct {
	p    sync.Pool
	data string
}

func NewPool(d string) *Pool {
	return &Pool{p: sync.Pool{New: func() any {
		h := strings.TrimPrefix(d, "0x")
		data, err := hex.DecodeString(h)
		if err != nil {
			panic(fmt.Sprintf("data string invalid:%s", err))
		}
		return data
	}}}
}

func (p *Pool) Get() []byte {
	return (p.p.Get()).([]byte)
}

func (p *Pool) Put(b []byte) {
	p.p.Put(b)
}

type cmPool struct {
	p sync.Pool
}

func (p *cmPool) Get() *CMBuffer {
	return (p.p.Get()).(*CMBuffer)
}

func (p *cmPool) Put(b *CMBuffer) {
	p.p.Put(b)
}
