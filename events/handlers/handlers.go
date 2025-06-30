package handlers

import (
	"fmt"
	"sync"

	"github.com/luan/events-driver-go/events"
)

type Handler interface {
	Exec(ctx *events.Context) error
}

type HandleManager struct {
	handlersByEvents map[events.Event][]Handler
}

var (
	handlerManager *HandleManager
	handlerOnce    sync.Once
)

func GetHandleManager() *HandleManager {
	handlerOnce.Do(func() {
		fmt.Println("🔧 Criando instância da Database (apenas uma vez)")
		handlerManager = &HandleManager{
			handlersByEvents: map[events.Event][]Handler{
				events.BuildError: {
					&CallDataFunnel{},
					&SendNotification{},
				},
				events.InspectionError: {
					&UpdateDatabase{},
					&SendNotification{},
				},
				events.InspectionSuccess: {
					&UpdateDatabase{},
					&SendNotification{},
				},
				events.SonnarFailed: {
					&CallDataFunnel{},
					&SendNotification{},
				},
			},
		}
	})

	return handlerManager
}

func (hm *HandleManager) TriggerEvent(ctx *events.Context) {
	if handlers, exists := hm.handlersByEvents[ctx.Event]; exists {
		fmt.Printf("Disparando evento: %s\n", ctx.Event)
		for _, handler := range handlers {
			handler.Exec(ctx)
		}
	} else {
		fmt.Printf("Nenhum handler registrado para o evento: %s\n", ctx.Event)
	}
}
