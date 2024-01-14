package contract

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

func padLeftZeroes(s string, length int) string {
	for len(s) < length {
		s = "0" + s
	}
	return s
}
func BuildClaimMethod(amt, nonce int, sign string) ([]byte, error) {
	methodID := "624f82f5"
	p1 := padLeftZeroes(big.NewInt(int64(amt)).Text(16), 64)
	p2 := padLeftZeroes(big.NewInt(int64(nonce)).Text(16), 64)
	p3 := "0000000000000000000000000000000000000000000000000000000000000060" // Offset
	p4 := "0000000000000000000000000000000000000000000000000000000000000041" // Length of the data
	p := "00000000000000000000000000000000000000000000000000000000000000"
	method := methodID + p1 + p2 + p3 + p4 + strings.TrimPrefix(sign, "0x") + p
	fmt.Println(method)

	return hex.DecodeString(method)
}
