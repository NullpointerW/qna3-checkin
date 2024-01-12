package contract

import (
	"fmt"
	"testing"
)

func TestSignAddress(t *testing.T) {
	signed, addr, err := SignAddress("test", "AI + DYOR = Ultimate Answer to Unlock Web3 Universe")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sign:", signed, "addr", addr)
}
