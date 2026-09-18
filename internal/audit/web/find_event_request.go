package web

import (
	"net/url"
	"strconv"
)

type FindEventRequest struct {
	ServiceName   string
	ProjectId     string
	EventType     string
	Action        string
	EntityType    string
	EntityID      string
	ActorID       string
	ActorUsername string
	RequestID     string
	SessionID     string

	OccurredFrom string
	OccurredTo   string

	Page  int
	Limit int
}

func NewFindEventRequest(query url.Values) FindEventRequest {
	return FindEventRequest{
		ServiceName:   query.Get("service_name"),
		ProjectId:     query.Get("project_id"),
		EventType:     query.Get("event_type"),
		Action:        query.Get("action"),
		EntityType:    query.Get("entity_type"),
		EntityID:      query.Get("entity_id"),
		ActorID:       query.Get("actor_id"),
		ActorUsername: query.Get("actor_username"),
		RequestID:     query.Get("request_id"),
		SessionID:     query.Get("session_id"),
		OccurredFrom:  query.Get("occurred_from"),
		OccurredTo:    query.Get("occurred_to"),
		Page:          getPage(query),
		Limit:         getLimit(query),
	}
}

func getPage(query url.Values) int {
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil || page < 1 {
		return 1
	}

	return page
}

func getLimit(query url.Values) int {
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || limit < 1 {
		return 20
	}

	if limit > 100 {
		return 100
	}

	return limit
}
