Aqui está o código convertido para Go, com a adaptação para usar a estrutura de chaves unitárias de Keccak:

```go
package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/consensys/gnark-crypto/ecc/bn256/fr"
	"github.com/consensys/gnark-crypto/ecc/bn256/g1"
	"github.com/consensys/gnark-crypto/ecc/bn256/g2"
	"github.com/stretchr/testify/assert"
)

func TestKeyGeneration(t *testing.T) {
	users := 10
	keys := make([][]byte, users)

	for i := 0; i < users; i++ {
		key := make([]byte, 32)
		hash := sha3.NewLegacyKeccak256()
		hash.Write([]byte(i))
		hash.Sum(key[:0])
		keys[i] = key
	}

	// Check if all keys are unique
	assert.Equal(t, users, len(unique(keys)), "Generated keys are not unique.")
}

func unique(slice [][]byte) [][]byte {
	keys := make(map[string]bool)
	list := make([][]byte, 0)

	for _, entry := range slice {
		key := crypto.Keccak256Hash(entry).Hex()
		if _, value := keys[key]; !value {
			keys[key] = true
			list = append(list, entry)
		}
	}

	return list
}

func TestPairings(t *testing.T) {
	params := bn256Init()
	a1 := params.G1.ScalarBaseMult(new(big.Int).SetInt64(1))
	a2 := params.G1.ScalarBaseMult(new(big.Int).SetInt64(2))
	a37 := params.G1.ScalarBaseMult(new(big.Int).SetInt64(37))
	an1 := params.G1.ScalarBaseMult(params.G1Order)

	b0 := params.G2.ScalarBaseMult(new(big.Int).SetInt64(0))
	b1 := params.G2.ScalarBaseMult(new(big.Int).SetInt64(1))
	b2 := params.G2.ScalarBaseMult(new(big.Int).SetInt64(2))
	b27 := params.G2.ScalarBaseMult(new(big.Int).SetInt64(27))
	b999 := params.G2.ScalarBaseMult(new(big.Int).SetInt64(999))
	bn1 := params.G2.ScalarBaseMult(params.G2Order)

	p1 := bn256Pair(a1, b1)
	pn1 := bn256Pair(a1, bn1)
	np1 := bn256Pair(an1, b1)
	assert.Equal(t, pn1.String(), np1.String(), "Pairing mismatch: e(a, -b) != e(-a, b)")
	assert.True(t, bn256PairingCheck(params, []g1.PointG1{a1, an1}, []g2.PointG2{b1, b1}), "MultiAte check gave false negative!")

	p0 := newGT().Add(p1, pn1)
	p0_2 := bn256Pair(a1, b0)
	assert.Equal(t, p0.String(), p0_2.String(), "Pairing mismatch: e(a, b) * e(a, -b) != 1")

	p0_3 := newGT().ScalarMult(p1, new(big.Int).Set(params.GTOrder))
	assert.Equal(t, p0.String(), p0_3.String(), "Pairing mismatch: e(a, b) has wrong order")

	p2 := bn256Pair(a2, b1)
	p2_2 := bn256Pair(a1, b2)
	p2_3 := newGT().ScalarMult(p1, new(big.Int).SetInt64(2))
	assert.Equal(t, p2.String(), p2_2.String(), "Pairing mismatch: e(a, b * 2) != e(a * 2, b)")
	assert.Equal(t, p2.String(), p2_3.String(), "Pairing mismatch: e(a, b * 2) != e(a, b) ** 2")
	assert.NotEqual(t, p2.String(), p1.String(), "Pairing is degenerate!")
	assert.False(t, bn256PairingCheck(params, []g1.PointG1{a1, a1}, []g2.PointG2{b1, b1}), "MultiAte check gave false positive!")

	p999 := bn256Pair(a37, b27)
	p999_2 := bn256Pair(a1, b999)
	assert.Equal(t, p999.String(), p999_2.String(), "Pairing mismatch: e(a * 37, b * 27) != e(a, b * 999)")
}

func bn256Init() *bn256Params {
	return &bn256Params{
		G1:       g1.New(),
		G2:       g2.New(),
		GT:       newGT(),
		G1Order:  fr.Modulus(),
		G2Order:  fr.Modulus(),
		GTOrder:  new(big.Int).SetInt64(21888242871839275222246405745257275088696311157297823662689037894645226208583),
		PairFunc: bn256Pair,
	}
}

func bn256Pair(a g1.PointG1, b g2.PointG2) *GT {
	return newGT().Pair(a, b)
}

type bn256Params struct {
	G1       *g1.G1Curve
	G2       *g2.G2Curve
	GT       *GT
	G1Order  *big.Int
	G2Order  *big.Int
	GTOrder  *big.Int
	PairFunc func(g1.PointG1, g2.PointG2) *GT
}

func bn256PairingCheck(params *bn256Params, a []g1.PointG1, b []g2.PointG2) bool {
	p := params.PairFunc(a[0], b[0])
	for i := 1; i < len(a); i++ {
		p.Mul(params.PairFunc(a[i], b[i]))
	}
	return p.Equal(params.GT.One())
}

func newGT() *GT {
	return &GT{
		fr.New().SetInt64(1),
	}
}

type GT struct {
	*fr.Element
}

func (p *GT) Pair(a g1.PointG1, b g2.PointG2) *GT {
	return p.Mul(newGT().ScalarProd(a.X(), b.Y()))
}

func (p *GT) Mul(x *GT) *GT {
	p.Element.Mul(p.Element, x.Element)
	return p
}

func (p *GT) Equal(x *GT) bool {
	return p.Element.Equal(x.Element)
}

func (p *GT) String() string {
	return p.Element.String()
}

func (p *GT) ScalarMult(x *GT, scalar *big.Int) *GT {
	p.Element.Exp(x.Element, scalar)
	return p
}

func (p *GT) ScalarProd(x *fr.Element, y *fr.Element) *GT {
	p.Element.Mul(x, y)
	return p
}
```

Este código Go executa os testes unitários originalmente