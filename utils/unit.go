package utils

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/params"
)

var (
	Wei   = big.NewInt(params.Wei)
	GWei  = big.NewInt(params.GWei)
	Ether = big.NewInt(params.Ether)
)

// EtherToWei converts Ether to Wei.
func EtherToWei(ether *big.Int) *big.Int {
	return new(big.Int).Mul(ether, big.NewInt(params.Ether))
}

// ToGWei converts Wei to GWei.
func ToGWei(wei *big.Int) *big.Int {
	return new(big.Int).Div(wei, GWei)
}

// ToEther converts Wei to Ether.
func ToEther(wei *big.Int) *big.Int {
	return new(big.Int).Div(wei, Ether)
}

// ToDigit separates by 3 digits.
func ToDigit(n *big.Int) string {
	arr := strings.Split(fmt.Sprintf("%d", n), "")
	cnt := len(arr) - 1
	res := ""
	i2 := 0
	for i := cnt; i >= 0; i-- {
		if i2 > 2 && i2%3 == 0 {
			res = fmt.Sprintf(",%s", res)
		}
		res = fmt.Sprintf("%s%s", arr[i], res)
		i2++
	}
	return res
}

// Adjust units and separate by 3 digits.
func FormatWei(amount *big.Int) string {
	if amount.Cmp(Ether) > 0 {
		return ToDigit(ToEther(amount)) + " OAS"
	} else if amount.Cmp(GWei) > 0 {
		return ToDigit(ToGWei(amount)) + " GWei"
	}
	return ToDigit(amount) + " Wei"
}
