package handlers

import (
	"github.com/luan/events-driver-go/events"
	"github.com/luan/events-driver-go/models"
)

type SendNotification struct{}

func (h *SendNotification) Exec(ctx *events.Context) error {
	println("Handling event:", ctx.Event.String())

	// Simulate sending a notification
	switch ctx.Event {
	case events.InspectionSuccess:
		println("Sending notification for inspection success...")
		println("Inspection details:", ctx.GetByKey(models.InspectionContextKey))
	case events.InspectionError:
		println("Sending notification for inspection error...")
		println("Inspection details:", ctx.GetByKey(models.InspectionContextKey))
	case events.BuildError:
		println("Sending notification for build error...")
		println("Inspection details:", ctx.GetByKey(models.BuildContextKey))
	case events.SonnarFailed:
		println("Sending notification for Sonnar failure...")
		println("Sonnar details:", ctx.GetByKey(models.SonnarContextKey))
	default:
		println("No notification sent for this event.")
	}

	return nil
}
