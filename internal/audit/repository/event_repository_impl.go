package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/bappeda-dev-team/audit-trail-service/internal/audit/domain"
)

type EventRepositoryImpl struct {
	DB *sql.DB
}

func NewEventRepository(
	db *sql.DB,
) *EventRepositoryImpl {
	return &EventRepositoryImpl{
		DB: db,
	}
}

func (repo *EventRepositoryImpl) Create(
	ctx context.Context,
	event *domain.AuditEvent,
) error {
	const op = "audit.repository.Create"

	query := `
		INSERT INTO audit_event (
			id,
			event_type,
			action,
			service_name,
			entity_type,
			entity_id,
			actor_id,
			actor_username,
			request_id,
			session_id,
			occurred_at,
			before_data,
			after_data,
			metadata
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12,
			$13,
			$14
		)
	`

	_, err := repo.DB.ExecContext(
		ctx,
		query,
		event.ID,
		event.EventType,
		event.Action,
		event.ServiceName,
		event.EntityType,
		event.EntityID,
		event.ActorID,
		event.ActorUsername,
		event.RequestID,
		event.SessionID,
		event.OccurredAt,
		event.Before,
		event.After,
		event.Metadata,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *EventRepositoryImpl) FindAll(
	ctx context.Context,
	filter domain.AuditEventFilter,
) ([]domain.AuditEvent, error) {
	const op = "audit.repository.FindAll"

	query := `
		SELECT
			id,
			event_type,
			action,
			service_name,
			entity_type,
			entity_id,
			actor_id,
			actor_username,
			request_id,
			session_id,
			occurred_at,
			before_data,
			after_data,
			metadata,
			created_at
		FROM audit_event
	`

	var (
		conditions []string
		args       []any
		argIndex   = 1
	)

	addCondition := func(condition string, value any) {
		conditions = append(
			conditions,
			fmt.Sprintf(condition, argIndex),
		)
		args = append(args, value)
		argIndex++
	}

	if filter.ServiceName != "" {
		addCondition(
			"service_name = $%d",
			filter.ServiceName,
		)
	}

	if filter.EventType != "" {
		addCondition(
			"event_type = $%d",
			filter.EventType,
		)
	}

	if filter.Action != nil {
		addCondition(
			"action = $%d",
			*filter.Action,
		)
	}

	if filter.EntityType != "" {
		addCondition(
			"entity_type = $%d",
			filter.EntityType,
		)
	}

	if filter.EntityID != "" {
		addCondition(
			"entity_id = $%d",
			filter.EntityID,
		)
	}

	if filter.ActorID != "" {
		addCondition(
			"actor_id = $%d",
			filter.ActorID,
		)
	}

	if filter.ActorUsername != "" {
		addCondition(
			"actor_username = $%d",
			filter.ActorUsername,
		)
	}

	if filter.RequestID != "" {
		addCondition(
			"request_id = $%d",
			filter.RequestID,
		)
	}

	if filter.SessionID != "" {
		addCondition(
			"session_id = $%d",
			filter.SessionID,
		)
	}

	if filter.OccurredFrom != nil {
		addCondition(
			"occurred_at >= $%d",
			*filter.OccurredFrom,
		)
	}

	if filter.OccurredTo != nil {
		addCondition(
			"occurred_at <= $%d",
			*filter.OccurredTo,
		)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += `
		ORDER BY occurred_at DESC, id DESC
	`

	limit := filter.Limit
	if limit <= 0 {
		limit = 20
	}

	page := filter.Page
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	query += fmt.Sprintf(
		" LIMIT $%d OFFSET $%d",
		argIndex,
		argIndex+1,
	)

	args = append(args, limit, offset)

	rows, err := repo.DB.QueryContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	events := make([]domain.AuditEvent, 0)

	for rows.Next() {
		var event domain.AuditEvent

		err := rows.Scan(
			&event.ID,
			&event.EventType,
			&event.Action,
			&event.ServiceName,
			&event.EntityType,
			&event.EntityID,
			&event.ActorID,
			&event.ActorUsername,
			&event.RequestID,
			&event.SessionID,
			&event.OccurredAt,
			&event.Before,
			&event.After,
			&event.Metadata,
			&event.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return events, nil
}
