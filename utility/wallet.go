package utility

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

/*func main() {
	fmt.Println(verifySig(
		"0x829814B6E4dfeC4b703F2c6fDba28F1724094D11",
		"0x53edb561b0c1719e46e1e6bbbd3d82ff798762a66d0282a9adf47a114e32cbc600c248c247ee1f0fb3a6136a05f0b776db4ac82180442d3a80f3d67dde8290811c",
		[]byte("hello"),
	))
}*/

func EthVerifySignature(address string, signature string, msg []byte) bool {
	fromAddr := common.HexToAddress(address)

	sig := hexutil.MustDecode(signature)
	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 27 && sig[64] != 28 {
		return false
	}
	sig[64] -= 27

	pubKey, err := crypto.SigToPub(ethSignatureHash(msg), sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return fromAddr == recoveredAddr
}

// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L404
// signHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calculated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func ethSignatureHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func VerifyWalletSignature(chain string, address string, signature string, msg []byte) bool {
	if chain == "bsc" {
		return EthVerifySignature(address, signature, msg)
	}
	return false
}

func BigIntDataToString(data []byte) string {
	var tmp big.Int
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return ""
	}
	return tmp.String()
}
