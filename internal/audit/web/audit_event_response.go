package web

import (
	"time"

	"github.com/bappeda-dev-team/audit-trail-service/internal/audit/domain"
)

type AuditEventResponse struct {
	ID string `json:"id" example:"8f4a2c1e-7b9d-4e62-a3f1-91d8c5b2e704"`

	EventType string `json:"event_type" example:"DATA_CHANGED"`
	Action    string `json:"action" example:"UPDATE"`

	ServiceName string `json:"service_name" example:"perencanaan-service"`

	EntityType string `json:"entity_type" example:"rencana_kinerja"`
	EntityID   string `json:"entity_id" example:"rk-2026-0001"`

	ActorID       *string `json:"actor_id,omitempty" example:"user-002"`
	ActorUsername *string `json:"actor_username,omitempty" example:"operator.opd"`

	RequestID *string `json:"request_id,omitempty" example:"req-004"`
	SessionID *string `json:"session_id,omitempty" example:"sess-004"`

	OccurredAt time.Time `json:"occurred_at" example:"2026-09-17T08:00:00+07:00"`

	Before any `json:"before,omitempty" swaggertype:"object"`
	After  any `json:"after,omitempty" swaggertype:"object"`

	Metadata any `json:"metadata,omitempty" swaggertype:"object"`

	CreatedAt time.Time `json:"created_at" example:"2026-09-17T08:00:01+07:00"`
}

func NewAuditEventResponse(
	event *domain.AuditEvent,
) AuditEventResponse {
	return AuditEventResponse{
		ID:            event.ID,
		EventType:     event.EventType,
		Action:        string(event.Action),
		ServiceName:   event.ServiceName,
		EntityType:    event.EntityType,
		EntityID:      event.EntityID,
		ActorID:       event.ActorID,
		ActorUsername: event.ActorUsername,
		RequestID:     event.RequestID,
		SessionID:     event.SessionID,
		OccurredAt:    event.OccurredAt,
		Before:        event.Before,
		After:         event.After,
		Metadata:      event.Metadata,
		CreatedAt:     event.CreatedAt,
	}
}

func NewAuditEventResponses(
	events []domain.AuditEvent,
) []AuditEventResponse {
	responses := make([]AuditEventResponse, 0, len(events))

	for i := range events {
		responses = append(
			responses,
			NewAuditEventResponse(&events[i]),
		)
	}

	return responses
}
