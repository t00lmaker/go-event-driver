// Exemplo de main.go mais idiomático
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/luan/events-driver-go/events"
	"github.com/luan/events-driver-go/events/handlers"
	"github.com/luan/events-driver-go/models"
)

func main() {
	// Go way: Context para cancelamento gracioso
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Go way: Capturar sinais do sistema
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("Recebido sinal de interrupção, finalizando...")
		cancel()
	}()

	// Go way: Criar dependências explicitamente
	manager := handlers.NewHandleManager()

	// Go way: Configurar handlers
	setupHandlers(manager)

	fmt.Println("🚀 Bem-vindo ao Events Driver Go!")
	fmt.Println("Pressione Ctrl+C para sair")

	// Go way: Loop com context
	if err := runEventLoop(ctx, manager); err != nil {
		log.Fatalf("Erro na execução: %v", err)
	}
}

func setupHandlers(manager *handlers.HandleManager) {
	// Go way: Configuração separada e explícita
	manager.RegisterHandler(events.BuildError, &handlers.CallDataFunnel{})
	manager.RegisterHandler(events.BuildError, &handlers.SendNotification{})
	// ... outros handlers
}

func runEventLoop(ctx context.Context, manager *handlers.HandleManager) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Go way: Tratamento de erro explícito
			if err := handleUserInput(reader, manager); err != nil {
				return fmt.Errorf("erro ao processar entrada: %w", err)
			}
		}
	}
}

func handleUserInput(reader *bufio.Reader, manager *handlers.HandleManager) error {
	fmt.Print("Digite uma opção (1-4, 0 para sair): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	switch input {
	case "0\n":
		return fmt.Errorf("saída solicitada pelo usuário")
	case "1\n":
		return triggerBuildError(manager)
	case "2\n":
		return triggerInspectionError(manager)
	// ... outros casos
	default:
		fmt.Println("Opção inválida")
		return nil
	}
}

func triggerBuildError(manager *handlers.HandleManager) error {
	build := &models.Build{
		ID:        "build-123",
		Status:    "error",
		ProjectID: "project-456",
	}

	return manager.TriggerEvent(events.BuildError.WithContext(build))
}
