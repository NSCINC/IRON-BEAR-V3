Aqui está o código traduzido para Go:

```go
package main

import "math/big"

var (
	R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11, R12, R13, R14, R15, R16 *big.Int
	np                                                                     []*big.Int
)

func mul(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15 *big.Int) {
	r1 := new(big.Int).Set(R1)
	r5_r1 := new(big.Int).Mul(r1, R5)
	c0.Set(r5_r1)
	c1.Set(new(big.Int).Rsh(r5_r1, 64))

	r0 := new(big.Int).Mul(r1, R6)
	c1.Add(c1, r0)
	c2.Set(new(big.Int).Rsh(r0, 64))

	r0.Mul(r1, R7)
	c2.Add(c2, r0)
	c3.Set(new(big.Int).Rsh(r0, 64))

	r0.Mul(r1, R8)
	c3.Add(c3, r0)
	c4.Set(new(big.Int).Rsh(r0, 64))

	c4.Add(c4, big.NewInt(0))

	r2 := new(big.Int).Set(R2)
	r1_r2 := new(big.Int).Mul(r2, R5)
	c5.Set(r1_r2)
	c6.Set(new(big.Int).Rsh(r1_r2, 64))

	r0.Mul(r2, R6)
	c6.Add(c6, r0)
	c7.Set(new(big.Int).Rsh(r0, 64))

	r0.Mul(r2, R7)
	c6.Add(c6, r0)
	c8.Set(new(big.Int).Rsh(r0, 64))

	r0.Mul(r2, R8)
	c6.Add(c6, r0)
	c9.Set(new(big.Int).Rsh(r0, 64))

	c1.Add(c1, r1)
	c2.Add(c2, R26)
	c3.Add(c3, R27)
	c4.Add(c4, R29)
	c5.Add(c5, big.NewInt(0))

	r3 := new(big.Int).Set(R3)
	r1_r3 := new(big.Int).Mul(r3, R5)
	c10.Set(r1_r3)
	c11.Set(new(big.Int).Rsh(r1_r3, 64))

	r0.Mul(r3, R6)
	c11.Add(c11, r0)
	c12.Set(new(big.Int).Rsh(r0, 64))

	r0.Mul(r3, R7)
	c11.Add(c11, r0)
	c13.Set(new(big.Int).Rsh(r0, 64))

	r0.Mul(r3, R8)
	c11.Add(c11, r0)
	c14.Set(new(big.Int).Rsh(r0, 64))

	c2.Add(c2, r1)
	c3.Add(c3, R26)
	c4.Add(c4, R27)
	c5.Add(c5, R29)
	c6.Add(c6, big.NewInt(0))

	r4 := new(big.Int).Set(R4)
	r1_r4 := new(big.Int).Mul(r4, R5)
	c14.Add(c14, r1_r4)
	c15.Set(new(big.Int).Rsh(r1_r4, 64))

	r0.Mul(r4, R6)
	c15.Add(c15, r0)
	c16.Set(new(big.Int).Rsh(r0, 64))

	r0.Mul(r4, R7)
	c15.Add(c15, r0)
	c17.Set(new(big.Int).Rsh(r0, 64))

	r0.Mul(r4, R8)
	c15.Add(c15, r0)
	c18.Set(new(big.Int).Rsh(r0, 64))

	c3.Add(c3, r1)
	c4.Add(c4, R26)
	c5.Add(c5, R27)
	c6.Add(c6, R29)
	c7.Add(c7, big.NewInt(0))

	// Chamada da função para redução
	gfpReduce(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15)
}

func gfpReduce(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15 *big.Int) {
	r17 := np[0]
	r25 := np[1]
	r19 := np[2]
	r20 := np[3]

	r1 := R9
	r17_r1 := new(big.Int).Mul(r1, r17)
	r2 := new(big.Int).Rsh(r17_r1, 64)
	r2.Add(r2, new(big.Int).Mul(r1, r25))
	r3 := new(big.Int).Rsh(r2, 64)
	r3.Add(r3, new(big.Int).Mul(r1, r19))
	r4 := new(big.Int).Rsh(r3, 64)
	r4.Add(r4, new(big.Int).Mul(r1, r20))

	r21 := new(big.Int).Mul(r1, R5)
	r22 := new(big.Int).Rsh(r21, 64)
	r22.Add(r22, new(big.Int).Mul(r1, R6))
	r23 := new(big.Int).Rsh(r22, 64)
	r23.Add(r23, new(big.Int).Mul(r1, R7))

	r24 := new(big.Int).Rsh(r23, 64)
	r24.Add(r24, new(big.Int).Mul(r1, R8))
	r24.Add(r24, big.NewInt(0))

	r2.Add(r2, r21)
	r3.Add(r3, r22)
	r4.Add(r4, r23)

	r1 = R10
	r21 = new(big.Int).Mul(r1, R5)
	r22 = new(big.Int).Rsh(r21, 64)
	r22.Add(r22, new(big.Int).Mul(r1, R6))
	r23 = new(big.Int).Rsh(r22, 64)
	r23.Add(r23, new(big.Int).Mul(r1, R7))

	r24.Add(r24, new(big.Int).Rsh(r23, 64))
	r24.Add(r24, new(big.Int).Mul(r1, R8))
	r24.Add(r24, big.NewInt(0))

	r2.Add(r2, r1)
	r3.Add(r3, r21)
	r4.Add(r4, r22)

	r1 = R11
	r21 = new(big.Int).Mul(r1, R5)
	r22 = new(big.Int).Rsh(r21, 64)
	r22.Add(r22, new(big.Int).Mul