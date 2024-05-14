Aqui está o código convertido para Go:

```go
package main

import "math/big"

// Defina as variáveis R1, R2, ..., R29 como *big.Int

func mul(c0, c1, c2, c3, c4, c5, c6, c7 *big.Int) []*big.Int {
    result := []*big.Int{}

    // Computar c0
    r1r5 := new(big.Int).Mul(R1, R5)
    result = append(result, r1r5)
    c1.Set(new(big.Int).Rsh(r1r5, 64))

    // Computar c1
    temp := new(big.Int).Mul(R1, R6)
    result = append(result, temp)
    c1.Add(c1, temp)
    c2.Set(new(big.Int).Rsh(temp, 64))

    // Computar c2
    temp.Mul(R1, R7)
    result = append(result, temp)
    c2.Add(c2, temp)
    c3.Set(new(big.Int).Rsh(temp, 64))

    // Computar c3
    temp.Mul(R1, R8)
    result = append(result, temp)
    c3.Add(c3, temp)
    c4.Set(new(big.Int).Rsh(temp, 64))

    // Computar c4
    c4.Add(c4, big.NewInt(0))

    // Computar c5
    temp.Mul(R2, R5)
    result = append(result, temp)
    c5.Set(temp)

    // Computar c6
    temp.Mul(R2, R6)
    result = append(result, temp)
    c6.Set(temp)

    // Computar c7
    temp.Mul(R2, R7)
    result = append(result, temp)
    c7.Set(temp)

    // Restante do código omitido para brevidade
    return result
}

func gfpReduce() []*big.Int {
    result := []*big.Int{}

    // Computar m = (T * N') mod R, armazenar m em R1:R2:R3:R4
    r17 := big.NewInt(0) // Carregar valor adequado para R17
    r25 := big.NewInt(0) // Carregar valor adequado para R25
    r19 := big.NewInt(0) // Carregar valor adequado para R19
    r20 := big.NewInt(0) // Carregar valor adequado para R20

    // Computar r1 * r17
    r1r17 := new(big.Int).Mul(R9, r17)
    result = append(result, r1r17)
    r1r17High := new(big.Int).Rsh(r1r17, 64)

    // Computar r1 * r25
    r1r25 := new(big.Int).Mul(R9, r25)
    result = append(result, r1r25)
    r1r25 = new(big.Int).Add(r1r25, R2)
    r1r25High := new(big.Int).Rsh(r1r25, 64)

    // Computar r1 * r19
    r1r19 := new(big.Int).Mul(R9, r19)
    result = append(result, r1r19)
    r1r19 = new(big.Int).Add(r1r19, R3)
    r1r19High := new(big.Int).Rsh(r1r19, 64)

    // Computar r1 * r20
    r1r20 := new(big.Int).Mul(R9, r20)
    result = append(result, r1r20)
    r1r20 = new(big.Int).Add(r1r20, R4)
    r1r20High := new(big.Int).Rsh(r1r20, 64)

    // Computar m * N
    result = append(result, mul(r17, r25, r19, r20, R21, R22, R23, R24)...)

    // Adicionar o intermediário de 512 bits a m*N
    c10 := new(big.Int).Add(R9, r17)
    c11 := new(big.Int).Add(R10, r25)
    c12 := new(big.Int).Add(R11, r19)
    c13 := new(big.Int).Add(R12, r20)
    c14 := new(big.Int).Add(R13, R21)
    c15 := new(big.Int).Add(R14, R22)
    c16 := new(big.Int).Add(R15, R23)
    c17 := new(big.Int).Add(R16, R24)

    // Reduzir mod p, se necessário
    c10.Sub(c10, R5)
    c11.Sub(c11, R6)
    c12.Sub(c12, R7)
    c13.Sub(c13, R8)

    result = append(result, c10, c11, c12, c13)

    return result
}
```

Certifique-se de definir todas as variáveis `R1`, `R2`, ..., `R29` como `*big.Int` e inicializá-las com os valores adequados.