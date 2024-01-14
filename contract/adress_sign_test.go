package contract

import (
	"fmt"
	"testing"
)

func TestSignAddress(t *testing.T) {
	signed, addr, err := SignMessage("test", "AI + DYOR = Ultimate Answer to Unlock Web3 Universe")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sign:", signed, "addr", addr)
}

func TestBuildClaimMethod(t *testing.T) {
	_, err := BuildClaimMethod(10, 1144805, "xxx")
	if err != nil {
		fmt.Println(err)
	}
}

func TestHexBuffer(t *testing.T) {
	buf := new(CMBuffer)
	hb, err := buf.Fill(10, 1144805, "xxx")
	if err != nil {
		fmt.Println(err)
	}

	shb, _ := BuildClaimMethod(10, 1144805, "xxx")
	fmt.Println(hb, "hb len", len(hb))
	fmt.Println(shb, "shb len", len(shb))
}

func TestName(t *testing.T) {

}
