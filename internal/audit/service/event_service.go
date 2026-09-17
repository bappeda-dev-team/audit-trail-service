package service

import (
	"context"

	"github.com/bappeda-dev-team/audit-trail-service/internal/audit/domain"
)

type EventService interface {
	CreateEvent(
		ctx context.Context,
		event domain.AuditEvent,
	) (*domain.AuditEvent, error)
	FindAllEvent(
		ctx context.Context,
		filter domain.AuditEventFilter,
	) ([]domain.AuditEvent, error)
}
