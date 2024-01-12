package contract

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

var (
	CheckinMethod = "0xe95a644f0000000000000000000000000000000000000000000000000000000000000001"
	ClaimMethod   = "0x624f82f5000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000111bcf00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000041%s00000000000000000000000000000000000000000000000000000000000000"
)

// SignAddress 使用私钥对以太坊地址进行签名
func SignAddress(privateKeyHex, msg string) (signatureHex, address string, err error) {
	// 将16进制的私钥字符串解码为ECDSA私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", "", fmt.Errorf("failed to decode private key: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	addr := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 对地址的字节进行Keccak256哈希，这是签名之前的标准步骤
	//addressHash := crypto.Keccak256Hash(addr.Bytes())

	prefixedHash := crypto.Keccak256Hash(
		[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(msg), msg)),
	)

	// 使用私钥对附加了前缀的消息哈希进行签名
	signature, err := crypto.Sign(prefixedHash.Bytes(), privateKey)

	// 签名地址哈希
	//signature, err := crypto.Sign(addressHash.Bytes(), privateKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign address: %v", err)
	}
	// https://stackoverflow.com/questions/69762108/implementing-ethereum-personal-sign-eip-191-from-go-ethereum-gives-different-s
	signature[64] += 27
	// 将签名字节转换为16进制字符串
	signatureHex = hexutil.Encode(signature)
	return signatureHex, addr.Hex(), nil
}
