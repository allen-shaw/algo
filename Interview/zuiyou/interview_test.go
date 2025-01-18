package zuiyou

import (
	"fmt"
	"math/big"
	"slices"
	"testing"
)

var mapping = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// 二进制->可见字符串
func base58Encode(ss []byte) string {
	result := make([]byte, 0, len(ss)*2)

	n := new(big.Int)
	n.SetBytes(ss)

	base := big.NewInt(56)
	zero := big.NewInt(0)
	for n.Cmp(zero) > 0 {
		mod := big.NewInt(0)
		n.DivMod(n, base, mod)
		result = append(result, mapping[mod.Int64()])
	}

	slices.Reverse(result)
	return string(result)
}

func Test_base58Encoding(t *testing.T) {
	ans := base58Encode([]byte("Hello, Base58!"))
	fmt.Println(ans)
}
