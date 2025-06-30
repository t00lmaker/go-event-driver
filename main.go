package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/luan/events-driver-go/events"
	"github.com/luan/events-driver-go/events/handlers"
	"github.com/luan/events-driver-go/models"
)

func main() {
	fmt.Println("🚀 Bem-vindo ao seu primeiro projeto Go!")
	fmt.Println("Este é o projeto events-driver")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Pressione Enter para contin0uar...")
		input, _ := reader.ReadString('\n')

		switch input {
		case "0\n":

			fmt.Println("Você escolheu sair do programa")
			fmt.Println("=== Fim do meu programa ===")
			return

		case "1\n":

			fmt.Println("Você escolheu a opção 1")

			build := &models.Build{
				ID:        "build-123",
				Status:    "error",
				ProjectID: "project-456",
			}

			handlers.GetHandleManager().TriggerEvent(events.BuildError.WithContext(build))

		case "2\n":

			fmt.Println("Você escolheu a opção 2")

			inspection := &models.Inspection{
				ID:        "inspection-789",
				Status:    "error",
				ProjectID: "project-456",
			}

			handlers.GetHandleManager().TriggerEvent(events.InspectionError.WithContext(inspection))

		case "3\n":
			fmt.Println("Você escolheu a opção 3")

			inspection := &models.Inspection{
				ID:        "inspection-789",
				Status:    "success",
				ProjectID: "project-456",
			}

			handlers.GetHandleManager().TriggerEvent(events.InspectionSuccess.WithContext(inspection))

		case "4\n":
			fmt.Println("Você escolheu a opção 4")

			sonnar := &models.Sonnar{
				ID:        "sonnar-101112",
				Status:    "failed",
				ProjectID: "project-456",
			}

			handlers.GetHandleManager().TriggerEvent(events.SonnarFailed.WithContext(sonnar))

		default:
			fmt.Println("Opção inválida, tente novamente")
		}
	}
}
