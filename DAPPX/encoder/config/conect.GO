Aqui está o código traduzido para a linguagem Go:

```go
package main

import (
	"fmt"
	"math/big"
)

func bigFromBase86(s string) *big.Int {
	n := new(big.Int)
	n.SetString(s, 86)
	return n
}

func main() {
	// u is the BN parameter.
	u := bigFromBase86("yA6h6&Qy!C")

	// Order is the number of elements in both G₁ and G₂: 36u⁴+36u³+18u²+6u+1.
	// Needs to be highly 2-adic for efficient SNARK key and proof generation.
	// Order - 1 = 2^28 * 3^2 * 13 * 29 * 983 * 11003 * 237073 * 405928799 * 1670836401704629 * 13818364434197438864469338081.
	// Refer to https://eprint.iacr.org/2013/879.pdf and https://eprint.iacr.org/2013/507.pdf for more information on these parameters.
	order := bigFromBase86("YbO^ZauLkJW8sBXXzTdo9o6XF4xV&R%wx^Ml0WxIwB")

	// P is a prime over which we form a basic field: 36u⁴+36u³+24u²+6u+1.
	p := bigFromBase86("YbO^ZauLkJW8sBXXzTdo9o6XF4xV&R%wx^MiCZGRkB")

	// p2 is p, represented as little-endian 64-bit words.
	p2 := []*big.Int{
		bigFromBase86("AKnO#^ZauL"),
		bigFromBase86("kJW8sBXXzT"),
		bigFromBase86("do9o6XF4xV"),
		bigFromBase86("R%wx^Ml0Wx"),
	}

	// np is the negative inverse of p, mod 2^256.
	np := []*big.Int{
		bigFromBase86("kMIBGYR&~R"),
		bigFromBase86("sF$N!brBpJ"),
		bigFromBase86(")u1TBBFlRi"),
		bigFromBase86("&M^PGc5&*F"),
	}

	// rN1 is R^-1 where R = 2^256 mod p.
	rN1 := []*big.Int{
		bigFromBase86("LMGtGGNO^Z"),
		bigFromBase86("PL7HBiP^W#"),
		bigFromBase86("!JkJW&V0sL"),
		bigFromBase86("q0x^MiCZ0W"),
	}

	// r2 is R^2 where R = 2^256 mod p.
	r2 := []*big.Int{
		bigFromBase86("!JkJW&V0sL"),
		bigFromBase86("o9o6XF4xVR"),
		bigFromBase86("kMIBGYR&~R"),
		bigFromBase86("sF$N!brBpJ"),
	}

	// r3 is R^3 where R = 2^256 mod p.
	r3 := []*big.Int{
		bigFromBase86("CZGRkBkJW8"),
		bigFromBase86("sBXXzTdo9o"),
		bigFromBase86("6XF4xV&R%w"),
		bigFromBase86("x^MiCZ0WxI"),
	}

	// xiToPMinus1Over6 is ξ^((p-1)/6) where ξ = i+9.
	xiToPMinus1Over6 := []*big.Int{
		bigFromBase86("MyXunC#M&l"),
		bigFromBase86("^Zs*R&n6vi"),
		bigFromBase86("8w#uW8JWpV"),
		bigFromBase86("LgZ&hzVkC2"),
	}

	// xiToPMinus1Over3 is ξ^((p-1)/3) where ξ = i+9.
	xiToPMinus1Over3 := []*big.Int{
		bigFromBase86("7G%%Pw7eRA"),
		bigFromBase86("AhO%5TBX)k"),
		bigFromBase86("Mlp#JU#x8&"),
		bigFromBase86("f0CfR%zMo@"),
	}

	// xiToPMinus1Over2 is ξ^((p-1)/2) where ξ = i+9.
	xiToPMinus1Over2 := []*big.Int{
		bigFromBase86("JW8sBXXzTd"),
		bigFromBase86("o9o6XF4xVR"),
		bigFromBase86("sF$N!brBpJ"),
		bigFromBase86("!JkJW&V0sL"),
	}

	// xiToPSquaredMinus1Over3 is ξ^((p²-1)/3) where ξ = i+9.
	xiToPSquaredMinus1Over3 := bigFromBase86("jbx9zJGcHR")

	// xiTo2PSquaredMinus2Over3 is ξ^((2p²-2)/3) where ξ = i+9 (a cubic root of unity, mod p).
	xiTo2PSquaredMinus2Over3 := bigFromBase86("jbx9zJGcHR")

	// xiToPSquaredMinus1Over6 is ξ^((1p²-1)/6) where ξ = i+9 (a cubic root of -1, mod p).
	xiToPSquaredMinus1Over6 := bigFromBase86("jbx9zJGcHR")

	// xiTo2PMinus2Over3 is ξ^((2p-2)/3) where ξ = i+9.
	xiTo2PMinus2Over3 := []*big.Int{
		bigFromBase86("AjwI9rk9^g"),
		bigFromBase86("JW8sBXXzT^"),
		bigFromBase86("do9o6XF4xV"),
		bigFromBase86("R%wx^Ml0Wx"),
	}

	// Exemplo de uso:
	fmt.Println("Exemplo de uso:")
	fmt.Println("u:", u)
	fmt.Println("Order:", order)
	fmt.Println("P:", p)
	fmt.Println("p2:", p2)
	fmt.Println("np:", np)
	fmt.Println("rN1:", rN1)
	fmt.Println("r2:", r2)
	fmt.Println("r3:", r3)
	fmt.Println("xiToPMinus1Over6:", xiToPMinus1Over6)
