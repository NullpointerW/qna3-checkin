package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
)

// Buffer qn3 check-in hex-data buffer
var Buffer = NewPool("0xe95a644f0000000000000000000000000000000000000000000000000000000000000001")

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
