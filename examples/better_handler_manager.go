package examples

import "github.com/luan/events-driver-go/events"

// Go way: Constructor function ao invés de singleton
func NewHandleManager() *HandleManager {
	return &HandleManager{
		handlersByEvents: make(map[events.Event][]Handler),
	}
}

// Go way: Métodos que recebem dependências explicitamente
func (hm *HandleManager) RegisterHandler(event events.Event, handler Handler) {
	hm.handlersByEvents[event] = append(hm.handlersByEvents[event], handler)
}

// Exemplo de uso mais idiomático
func ExampleUsage() {
	// Criar instância explicitamente
	manager := NewHandleManager()

	// Registrar handlers de forma explícita
	manager.RegisterHandler(events.BuildError, &CallDataFunnel{})
	manager.RegisterHandler(events.BuildError, &SendNotification{})

	// Usar a instância
	manager.TriggerEvent(context)
}
