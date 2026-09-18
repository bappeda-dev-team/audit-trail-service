package domain

import "time"

type AuditEventFilter struct {
	ServiceName   string
	ProjectId     string
	EventType     string
	Action        *AuditAction
	EntityType    string
	EntityID      string
	ActorID       string
	ActorUsername string
	RequestID     string
	SessionID     string

	OccurredFrom *time.Time
	OccurredTo   *time.Time

	Page  int
	Limit int
}
