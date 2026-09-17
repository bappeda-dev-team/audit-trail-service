package repository

import (
	"context"

	"github.com/bappeda-dev-team/audit-trail-service/internal/audit/domain"
)

type EventRepository interface {
	Create(
		ctx context.Context,
		event *domain.AuditEvent,
	) error
	FindAll(
		ctx context.Context,
		filter domain.AuditEventFilter,
	) ([]domain.AuditEvent, error)
}
