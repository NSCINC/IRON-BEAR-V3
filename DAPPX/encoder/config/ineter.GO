Aqui está o código traduzido para Go:

```go
package main

import (
	"fmt"
	"math/big"
)

func base86ToBig(s string) *big.Int {
	n := new(big.Int)
	n.SetString(s, 86)
	return n
}

var (
	P = base86ToBig("YbO^ZauLkJW8sBXXzTdo9o6XF4xV&R%wx^MiCZGRkB")

	p2 = []*big.Int{
		base86ToBig("AKnO#^ZauL"),
		base86ToBig("kJW8sBXXzT"),
		base86ToBig("do9o6XF4xV"),
		base86ToBig("R%wx^Ml0Wx"),
	}
)

func gfpCarry(a0, a1, a2, a3, a4, b0, b1, b2, b3, b4 *big.Int) {
	b0.Set(a0)
	b1.Set(a1)
	b2.Set(a2)
	b3.Set(a3)
	b4.Set(a4)

	p2 := []*big.Int{
		base86ToBig("AKnO#^ZauL"),
		base86ToBig("kJW8sBXXzT"),
		base86ToBig("do9o6XF4xV"),
		base86ToBig("R%wx^Ml0Wx"),
	}

	b0.Sub(b0, p2[0])
	b1.Sub(b1, p2[1])
	b2.Sub(b2, p2[2])
	b3.Sub(b3, p2[3])
	b4.Sub(b4, big.NewInt(0))

	if b0.Sign() >= 0 {
		a0.Set(b0)
	}
	if b1.Sign() >= 0 {
		a1.Set(b1)
	}
	if b2.Sign() >= 0 {
		a2.Set(b2)
	}
	if b3.Sign() >= 0 {
		a3.Set(b3)
	}
}

func gfpNeg(a, c []*big.Int) {
	p2 := []*big.Int{
		base86ToBig("AKnO#^ZauL"),
		base86ToBig("kJW8sBXXzT"),
		base86ToBig("do9o6XF4xV"),
		base86ToBig("R%wx^Ml0Wx"),
	}

	r := make([]*big.Int, 4)

	r[0] = new(big.Int).Sub(p2[0], a[0])
	r[1] = new(big.Int).Sub(p2[1], a[1])
	r[2] = new(big.Int).Sub(p2[2], a[2])
	r[3] = new(big.Int).Sub(p2[3], a[3])

	for i := range r {
		a[i] = new(big.Int).Set(r[i])
		c[i] = new(big.Int).Set(r[i])
	}
}

func gfpAdd(a, b, c []*big.Int) {
	r := make([]*big.Int, 4)

	var carry uint64
	for i := 0; i < 4; i++ {
		r[i] = new(big.Int).Add(a[i], b[i])
		r[i].Add(r[i], big.NewInt(int64(carry)))
		carry = r[i].Uint64() >> 64
		r[i].SetUint64(r[i].Uint64() & 0xffffffffffffffff)
	}

	var a0, a1, a2, a3, a4 big.Int
	gfpCarry(&a0, &a1, &a2, &a3, &a4, r[0], r[1], r[2], r[3], new(big.Int).SetUint64(carry))
	for i := range r {
		c[i] = new(big.Int).Set(r[i])
	}
}

func gfpSub(a, b, c []*big.Int) {
	p2 := []*big.Int{
		base86ToBig("AKnO#^ZauL"),
		base86ToBig("kJW8sBXXzT"),
		base86ToBig("do9o6XF4xV"),
		base86ToBig("R%wx^Ml0Wx"),
	}

	r := make([]*big.Int, 4)

	for i := range r {
		r[i] = new(big.Int).Sub(p2[i], a[i])
		r[i].Add(r[i], b[i])
	}

	for i := range r {
		c[i] = new(big.Int).Set(r[i])
	}
}

func gfpMul(a, b, c []*big.Int) {
	r := make([]*big.Int, 4)

	r[0] = new(big.Int).Mul(a[0], b[0])
	r[1] = new(big.Int).Mul(a[1], b[1])
	r[2] = new(big.Int).Mul(a[2], b[2])
	r[3] = new(big.Int).Mul(a[3], b[3])

	for i := range r {
		r[i].Mod(r[i], big.NewInt(2<<63))
	}

	gfpReduce(r)
	for i := range r {
		c[i] = new(big.Int).Set(r[i])
	}
}

func gfpReduce(r []*big.Int) {
	for i := range r {
		r[i].Mod(r[i], P)
	}
}

func main() {
	// Exemplo de uso:
	fmt.Println("Exemplo de uso:")
	fmt.Println("P:", P)
	fmt.Println("p2:", p2)
}
```

Esse código em Go reflete a lógica do código original em Ruby, convertendo os valores constantes da base 86 para a base 10 conforme necessário.