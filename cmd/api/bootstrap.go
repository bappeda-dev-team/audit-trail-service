package main

import (
	"database/sql"
	"log/slog"

	"github.com/bappeda-dev-team/audit-trail-service/internal/api"
	"github.com/bappeda-dev-team/audit-trail-service/internal/audit/repository"
	"github.com/bappeda-dev-team/audit-trail-service/internal/audit/service"
)

// buildApplication adalah fungsi untuk injeksi repo dan service
// tambahkan service dan repository baru kesini
func buildApplication(
	cfg api.Config,
	logger *slog.Logger,
	db *sql.DB,
) *api.Application {

	// repository
	eventRepository := repository.NewEventRepository(db)

	// service
	eventService := service.NewEventService(
		eventRepository,
		logger,
	)

	// application
	app := &api.Application{
		Config: cfg,
		Logger: logger,

		EventService: eventService,
	}

	return app
}
