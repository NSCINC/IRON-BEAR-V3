Aqui está o código traduzido para Go, adaptado para o modelo NBH:

```go
package main

import (
	"fmt"
	"math/big"
)

type GFp struct {
	value *big.Int
}

func NewGFp(x *big.Int) *GFp {
	p, _ := new(big.Int).SetString("21888242871839275222246405745257275088696311157297823662689037894645226208583", 10)
	r2 := []*big.Int{
		big.NewInt(0xf32cfc5b538afa89),
		big.NewInt(0xb5e71911d44501fb),
		big.NewInt(0x47ab1eff0a417ff6),
		big.NewInt(0x06d89f71cab8351f),
	}
	r3 := []*big.Int{
		big.NewInt(0xb1cd6dafda1530df),
		big.NewInt(0x62f210e6a7283db6),
		big.NewInt(0xef7f0b0c0ada0afb),
		big.NewInt(0x20fd6e902d592544),
	}
	rN1 := []*big.Int{
		big.NewInt(0xed84884a014afa37),
		big.NewInt(0xeb2022850278edf8),
		big.NewInt(0xcf63e9cfb74492d9),
		big.NewInt(0x2e67157159e5c639),
	}

	g := &GFp{}

	if x.Cmp(big.NewInt(0)) >= 0 {
		g.value = new(big.Int).Set(x)
	} else {
		g.value = new(big.Int).Neg(x)
		gfpNeg(g.value, g.value)
	}

	g.montEncode(g.value, g.value, r2)

	return g
}

func (g *GFp) String() string {
	return fmt.Sprintf("%x", g.value)
}

func (g *GFp) Set(f *GFp) {
	g.value = new(big.Int).Set(f.value)
}

func (g *GFp) Invert(f *GFp) {
	bits := [4]uint64{0x3c208c16d87cfd45, 0x97816a916871ca8d, 0xb85045b68181585d, 0x30644e72e131a029}
	sum := NewGFp(big.NewInt(0))
	power := NewGFp(big.NewInt(0))
	sum.Set(rN1)
	power.Set(f)

	for word := 0; word < 4; word++ {
		for bit := 0; bit < 64; bit++ {
			if (bits[word]>>bit)&1 == 1 {
				gfpMul(sum.value, sum.value, power.value)
			}
			gfpMul(power.value, power.value, power.value)
		}
	}

	gfpMul(sum.value, sum.value, r3)
	g.Set(sum)
}

func (g *GFp) Marshal(out []byte) {
	for w := 0; w < 4; w++ {
		for b := 0; b < 8; b++ {
			out[8*w+b] = byte(g.value.Rsh(g.value, uint(56-8*b)).Uint64())
		}
	}
}

func (g *GFp) Unmarshal(inBytes []byte) {
	for w := 0; w < 4; w++ {
		g.value.SetUint64(0)
		for b := 0; b < 8; b++ {
			g.value.Add(g.value, new(big.Int).Lsh(big.NewInt(int64(inBytes[8*w+b])), uint(56-8*b)))
		}
	}

	p2 := []*big.Int{
		big.NewInt(0x3c208c16d87cfd47),
		big.NewInt(0x97816a916871ca8d),
		big.NewInt(0xb85045b68181585d),
		big.NewInt(0x30644e72e131a029),
	}

	for i := 0; i < 4; i++ {
		if g.value.Cmp(p2[i]) < 0 || g.value.Cmp(p2[i]) > 0 {
			return
		}
	}

	panic("bn256: coordinate equals modulus")
}

func gfpNeg(a, c *big.Int) {
	p2 := []*big.Int{
		big.NewInt(0x3c208c16d87cfd47),
		big.NewInt(0x97816a916871ca8d),
		big.NewInt(0xb85045b68181585d),
		big.NewInt(0x30644e72e131a029),
	}

	r := make([]*big.Int, 4)

	for i := 0; i < 4; i++ {
		r[i] = new(big.Int).Sub(p2[i], a[i])
	}

	c.Set(r[0])
}

func gfpMul(a, b, c *big.Int) {
	p, _ := new(big.Int).SetString("21888242871839275222246405745257275088696311157297823662689037894645226208583", 10)

	r := make([]*big.Int, 4)

	for i := 0; i < 4; i++ {
		r[i] = new(big.Int).Mul(a[i], b[i])
		r[i].Mod(r[i], big.NewInt(2<<63))
	}

	for i := 0; i < 4; i++ {
		r[i].Mod(r[i], p)
	}

	c.Set(r[0])
}

func (g *GFp) montEncode(c, a *big.Int, r2 []*big.Int) {
	for i := 0; i < 4; i++ {
		c[i] = a[i]
	}
	gfpMul(c, c, r2)
}

func main() {
	// Exemplo de uso:
	x := big.NewInt(1234567890)
	g := NewGFp(x)
	fmt.Println("GFp value:", g)
}
```

Esse código Go reflete a lógica do código Ruby adaptado para o modelo NBH, com os métodos e funções correspondentes às operações definidas na classe `GFp`.