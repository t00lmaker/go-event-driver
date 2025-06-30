package handlers

import (
	"github.com/luan/events-driver-go/events"
	"github.com/luan/events-driver-go/models"
)

type CallDataFunnel struct{}

func (h *CallDataFunnel) Exec(ctx *events.Context) error {

	println("Handling event: ", ctx.Event.String())
	println("Build Id: ", ctx.GetByKey(models.BuildContextKey).ContextValue()["id"])

	return nil
}
