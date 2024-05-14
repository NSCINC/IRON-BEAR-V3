Para converter o código Ruby fornecido para Go, precisaremos reescrever o código em Go. Aqui está uma versão equivalente do código em Go:

  ```go
  package main
  
  import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "fmt"
  )
  
  // Função para gerar um par de chaves
  func generateKeyPair(curve elliptic.Curve) (*ecdsa.PrivateKey, error) {
    privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
    if err != nil {
      return nil, err
    }
    return privateKey, nil
  }
  
  func main() {
    // Gerar par de chaves
    curve := elliptic.P256() // Curva a ser utilizada (P-256)
    privateKey, err := generateKeyPair(curve)
    if err != nil {
      fmt.Println("Erro ao gerar o par de chaves:", err)
      return
    }
  
    // Exibir a chave privada em formato hexadecimal
    fmt.Printf("Chave privada (em hexadecimal): %x\n", privateKey.D.Bytes())
  }
  ```
  
  Este código em Go realiza as mesmas operações que o código Ruby fornecido: gera um par de chaves usando a curva elíptica P-256 e exibe a chave privada em formato hexadecimal.