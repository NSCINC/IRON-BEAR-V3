Aqui está o código convertido para Go, adaptado conforme solicitado:

```go
// Package crypto fornece funções para criptografia e assinatura digital.
package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

// GenerateKeyPair gera um par de chaves usando a curva elíptica especificada.
func GenerateKeyPair(curve elliptic.Curve) (*ecdsa.PrivateKey, error) {
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// GenerateSharedSecret gera o segredo compartilhado usando a chave privada e a chave pública.
func GenerateSharedSecret(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	x, _ := privateKey.Curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	return x.Bytes(), nil
}

func main() {
	// Gerar par de chaves
	curve := elliptic.P256() // Curva a ser utilizada (P-256)
	privateKey, err := GenerateKeyPair(curve)
	if err != nil {
		fmt.Println("Erro ao gerar o par de chaves:", err)
		return
	}

	// Gerar a chave pública
	publicKey := &privateKey.PublicKey

	// Gerar o segredo compartilhado
	sharedSecret, err := GenerateSharedSecret(privateKey, publicKey)
	if err != nil {
		fmt.Println("Erro ao gerar o segredo compartilhado:", err)
		return
	}

	// Exibir o segredo compartilhado em formato hexadecimal
	fmt.Printf("Segredo compartilhado (em hexadecimal): %x\n", sharedSecret)
}
```

Este código em Go realiza as mesmas operações que o código Ruby fornecido: gera um par de chaves usando a curva elíptica P-256 e gera um segredo compartilhado usando a chave privada e a chave pública. Por fim, exibe o segredo compartilhado em formato hexadecimal.