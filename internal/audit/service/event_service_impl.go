package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/bappeda-dev-team/audit-trail-service/internal/audit/domain"
	"github.com/bappeda-dev-team/audit-trail-service/internal/audit/repository"
)

type EventServiceImpl struct {
	EventRepository repository.EventRepository
	Logger          *slog.Logger
}

func NewEventService(
	eventRepository repository.EventRepository,
	logger *slog.Logger,
) *EventServiceImpl {
	return &EventServiceImpl{
		EventRepository: eventRepository,
		Logger:          logger,
	}
}

func (service *EventServiceImpl) CreateEvent(
	ctx context.Context,
	event domain.AuditEvent,
) (*domain.AuditEvent, error) {
	const op = "audit.service.CreateEvent"

	if strings.TrimSpace(event.ProjectId) == "" {
		return nil, errors.New("project_id is required")
	}

	if strings.TrimSpace(event.EventType) == "" {
		return nil, errors.New("event_type is required")
	}

	if strings.TrimSpace(event.ServiceName) == "" {
		return nil, errors.New("service_name is required")
	}

	if strings.TrimSpace(event.EntityType) == "" {
		return nil, errors.New("entity_type is required")
	}

	if strings.TrimSpace(event.EntityID) == "" {
		return nil, errors.New("entity_id is required")
	}

	if event.OccurredAt.IsZero() {
		event.OccurredAt = time.Now()
	}

	if err := service.EventRepository.Create(
		ctx,
		&event,
	); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &event, nil
}

func (service *EventServiceImpl) FindAllEvent(
	ctx context.Context,
	filter domain.AuditEventFilter,
) ([]domain.AuditEvent, error) {
	const op = "audit.service.FindAllEvent"

	if filter.Page < 1 {
		filter.Page = 1
	}

	if filter.Limit < 1 {
		filter.Limit = 20
	}

	if filter.Limit > 100 {
		filter.Limit = 100
	}

	result, err := service.EventRepository.FindAll(
		ctx,
		filter,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}
