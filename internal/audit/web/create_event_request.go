package web

import "time"

type CreateEventRequest struct {
	EventType string `json:"event_type"`
	Action    string `json:"action"`

	ServiceName string `json:"service_name"`

	EntityType string `json:"entity_type"`
	EntityID   string `json:"entity_id"`

	ActorID       *string `json:"actor_id,omitempty"`
	ActorUsername *string `json:"actor_username,omitempty"`

	RequestID *string `json:"request_id,omitempty"`
	SessionID *string `json:"session_id,omitempty"`

	OccurredAt *time.Time `json:"occurred_at,omitempty"`

	Before   any            `json:"before,omitempty"`
	After    any            `json:"after,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}
