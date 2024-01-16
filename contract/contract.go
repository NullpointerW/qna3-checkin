package contract

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/NullpointerW/ethereum-wallet-tool/pkg/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var (
	Address       = "0xb342e7d33b806544609370271a8d074313b7bc30"
	CheckinMethod = "0xe95a644f0000000000000000000000000000000000000000000000000000000000000001"
	//ClaimMethod   = "0x624f82f5000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000111bcf00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000041%s00000000000000000000000000000000000000000000000000000000000000"
)

func CallCheckin(pk string, rpc *ethclient.Client) (common.Hash, error) {
	data := CKBuffer.Get()
	defer CKBuffer.Put(data)
	txHash, err := tx.Transfer(pk, Address, "0", data, rpc)
	if err != nil {
		return common.Hash{}, err
	}
	//fmt.Println(txHash.String())
	_, err = tx.WaitForTransactionConfirmation(rpc, txHash)
	return txHash, err
}

func CallClaim(pk string, amt, nonce int, sign string, rpc *ethclient.Client) (common.Hash, error) {
	buffer := CLBuffer.Get()
	defer CLBuffer.Put(buffer)
	data, err := buffer.Fill(uint32(amt), uint32(nonce), sign)
	if err != nil {
		return common.Hash{}, err
	}
	//hb, err := BuildClaimMethod(amt, nonce, sign)
	if err != nil {
		return common.Hash{}, err
	}
	txHash, err := tx.Transfer(pk, Address, "0", data, rpc)
	if err != nil {
		return common.Hash{}, err
	}
	//fmt.Println(txHash.String())
	_, err = tx.WaitForTransactionConfirmation(rpc, txHash)
	return txHash, err
}

// SignMessage 使用私钥对消息进行签名
func SignMessage(privateKeyHex, msg string) (signatureHex, address string, err error) {
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
