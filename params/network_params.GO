Aqui está uma estrutura básica de código para criar uma rede local de blockchain e conectar uma DApp em Go:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // Iniciar a rede local de blockchain
    if err := startBlockchain(); err != nil {
        log.Fatalf("Erro ao iniciar a rede blockchain: %v", err)
    }
    
    // Iniciar a DApp em Go
    go startDApp()
    
    // Aguardar sinal de interrupção para encerrar a execução
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
    <-interrupt
    
    // Encerrar a rede blockchain
    if err := stopBlockchain(); err != nil {
        log.Printf("Erro ao encerrar a rede blockchain: %v", err)
    }
    
    fmt.Println("Encerrando a execução...")
}

func startBlockchain() error {
    // Iniciar o nó da blockchain com o Geth (Go Ethereum)
    cmd := exec.Command("geth", "--datadir", "~/blockchain", "init", "genesis.json")
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("falha ao inicializar o nó da blockchain: %v", err)
    }
    
    // Iniciar o nó da blockchain
    cmd = exec.Command("geth", "--datadir", "~/blockchain", "--rpc", "--rpcapi", "eth,net,web3,personal", "console")
    if err := cmd.Start(); err != nil {
        return fmt.Errorf("falha ao iniciar o nó da blockchain: %v", err)
    }
    
    // Aguardar a conclusão do nó da blockchain
    go func() {
        if err := cmd.Wait(); err != nil {
            log.Printf("Erro ao executar o nó da blockchain: %v", err)
        }
    }()
    
    return nil
}

func stopBlockchain() error {
    // Encerrar o nó da blockchain
    cmd := exec.Command("pkill", "geth")
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("falha ao encerrar o nó da blockchain: %v", err)
    }
    
    return nil
}

func startDApp() {
    // Iniciar a DApp em Go
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Bem-vindo à DApp em Go!")
    })
    
    fmt.Println("DApp está sendo executada em: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Este código em Go estrutura uma aplicação que inicia uma rede local de blockchain com Geth (Go Ethereum) e uma DApp em Go simultaneamente. A rede blockchain é iniciada como um processo separado e a DApp é servida através do servidor HTTP interno do Go. O programa aguarda um sinal de interrupção para encerrar a execução e, em seguida, encerra a rede blockchain antes de finalizar. Certifique-se de ajustar os comandos Geth conforme necessário para iniciar sua blockchain local.