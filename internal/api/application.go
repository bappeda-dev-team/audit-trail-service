package api

import (
	"log/slog"

	"github.com/bappeda-dev-team/audit-trail-service/internal/audit/service"
)

const Version = "1.0.0"

// port: which port of server should be listening to
// env: which env should used, prod, staging, and dev
type Config struct {
	Port int
	Env  string

	DB struct {
		Dsn string
	}
}

// hold application dependencies for
// http handler, middleware and helpers
type Application struct {
	Config Config
	Logger *slog.Logger

	// inject service
	EventService service.EventService
}
