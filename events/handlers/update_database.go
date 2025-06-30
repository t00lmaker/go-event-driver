package handlers

import (
	"github.com/luan/events-driver-go/events"
	"github.com/luan/events-driver-go/models"
)

type UpdateDatabase struct{}

func (h *UpdateDatabase) Exec(ctx *events.Context) error {
	println("Handling event:", ctx.Event.String())

	// Simulate updating the database
	switch ctx.Event {
	case events.InspectionSuccess:
		println("Updating database with inspection success data...")
		println("Inspection details:", ctx.GetByKey(models.InspectionContextKey))
	case events.InspectionError:
		println("Updating database with inspection error data...")
		println("Inspection details:", ctx.GetByKey(models.InspectionContextKey))
	default:
		println("No database update for this event.")
	}

	return nil
}
