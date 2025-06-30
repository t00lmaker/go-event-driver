package examples

import (
	"fmt"

	"github.com/luan/events-driver-go/events"
)

// Go way: Sempre retornar erros
func (hm *HandleManager) TriggerEvent(ctx *events.EventContext) error {
	handlers, exists := hm.handlersByEvents[ctx.Event]
	if !exists {
		return fmt.Errorf("nenhum handler registrado para evento %s", ctx.Event)
	}

	// Go way: Coletar erros de múltiplos handlers
	var errors []error
	for _, handler := range handlers {
		if err := handler.Exec(ctx); err != nil {
			errors = append(errors, fmt.Errorf("handler falhou: %w", err))
		}
	}

	// Go way: Retornar erro combinado se houver falhas
	if len(errors) > 0 {
		return fmt.Errorf("alguns handlers falharam: %v", errors)
	}

	return nil
}

// Go way: Validação de entrada
func (c *EventContext) Validate() error {
	if c.Event == 0 {
		return fmt.Errorf("evento não pode ser vazio")
	}
	if len(c.Params) == 0 {
		return fmt.Errorf("contexto deve ter pelo menos um parâmetro")
	}
	return nil
}
