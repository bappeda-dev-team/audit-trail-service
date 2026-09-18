package domain

import (
	"encoding/json"
	"time"
)

type AuditAction string

const (
	AuditActionCreate AuditAction = "CREATE"
	AuditActionUpdate AuditAction = "UPDATE"
	AuditActionDelete AuditAction = "DELETE"
)

type AuditEvent struct {
	ID string

	ProjectId string
	EventType string
	Action    AuditAction

	ServiceName string

	EntityType string
	EntityID   string

	ActorID       *string
	ActorUsername *string

	RequestID *string
	SessionID *string

	OccurredAt time.Time

	Before   json.RawMessage
	After    json.RawMessage
	Metadata json.RawMessage

	CreatedAt time.Time
}
