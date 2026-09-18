package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/bappeda-dev-team/audit-trail-service/internal/audit/domain"
	auditWeb "github.com/bappeda-dev-team/audit-trail-service/internal/audit/web"
	"github.com/bappeda-dev-team/audit-trail-service/internal/web"
)

// CreateEventHandler godoc
//
// @Summary     Create Audit Event
// @Description Menyimpan event audit trail
// @Tags        Audit
// @Accept      json
// @Produce     json
//
// @Param       request body auditWeb.CreateEventRequest true "Audit Event"
//
// @Success     201 {object} web.Response[auditWeb.AuditEventResponse] "Berhasil membuat audit event"
// @Failure     400 {object} web.ValidationErrorResponse "Bad Request"
// @Failure     500 {object} web.ErrorResponse "Internal Server Error"
//
// @Router      /events [post]
func (app *Application) CreateEventHandler(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	var req auditWeb.CreateEventRequest

	if err := app.ReadJSON(w, r, &req); err != nil {
		app.BadRequestResponse(w, r, err)
		return
	}

	event := domain.AuditEvent{
		ProjectId:     req.ProjectId,
		EventType:     req.EventType,
		Action:        domain.AuditAction(req.Action),
		ServiceName:   req.ServiceName,
		EntityType:    req.EntityType,
		EntityID:      req.EntityID,
		ActorID:       req.ActorID,
		ActorUsername: req.ActorUsername,
		RequestID:     req.RequestID,
		SessionID:     req.SessionID,
	}

	if req.OccurredAt != nil {
		event.OccurredAt = *req.OccurredAt
	}

	before, err := json.Marshal(req.Before)
	if err != nil {
		app.BadRequestResponse(w, r, err)
		return
	}

	after, err := json.Marshal(req.After)
	if err != nil {
		app.BadRequestResponse(w, r, err)
		return
	}

	metadata, err := json.Marshal(req.Metadata)
	if err != nil {
		app.BadRequestResponse(w, r, err)
		return
	}

	event.Before = before
	event.After = after
	event.Metadata = metadata

	result, err := app.EventService.CreateEvent(
		r.Context(),
		event,
	)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	response := web.Response[auditWeb.AuditEventResponse]{
		Data: auditWeb.NewAuditEventResponse(result),
	}

	if err := app.WriteJSON(
		w,
		http.StatusCreated,
		response,
		nil,
	); err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}
}

// FindEventHandler godoc
//
// @Summary     Get Audit Events
// @Description Mengambil audit event berdasarkan filter
// @Tags        Audit
// @Accept      json
// @Produce     json
//
// @Param       service_name   query string false "Nama service"
// @Param       project_id     query string false "Id Project"
// @Param       event_type     query string false "Tipe event"
// @Param       action         query string false "Action: CREATE, UPDATE, DELETE"
// @Param       entity_type    query string false "Tipe entity"
// @Param       entity_id      query string false "ID entity"
// @Param       actor_id       query string false "ID actor"
// @Param       actor_username query string false "Username actor"
// @Param       request_id     query string false "Request ID"
// @Param       session_id     query string false "Session ID"
// @Param       occurred_from  query string false "Waktu mulai" Format(date-time)
// @Param       occurred_to    query string false "Waktu akhir" Format(date-time)
// @Param       page           query int    false "Nomor halaman" default(1)
// @Param       limit          query int    false "Jumlah data" default(20)
//
// @Success     200 {object} web.Response[[]web.AuditEventResponse] "Berhasil mengambil audit events"
// @Failure     400 {object} web.ValidationErrorResponse "Bad Request"
// @Failure     500 {object} web.ErrorResponse "Internal Server Error"
//
// @Router      /events [get]
func (app *Application) FindEventHandler(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	req := auditWeb.NewFindEventRequest(
		r.URL.Query(),
	)

	filter := domain.AuditEventFilter{
		ServiceName:   req.ServiceName,
		ProjectId:     req.ProjectId,
		EventType:     req.EventType,
		EntityType:    req.EntityType,
		EntityID:      req.EntityID,
		ActorID:       req.ActorID,
		ActorUsername: req.ActorUsername,
		RequestID:     req.RequestID,
		SessionID:     req.SessionID,
		Page:          req.Page,
		Limit:         req.Limit,
	}

	if req.Action != "" {
		action := domain.AuditAction(req.Action)
		filter.Action = &action
	}

	result, err := app.EventService.FindAllEvent(
		r.Context(),
		filter,
	)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	response := web.Response[[]auditWeb.AuditEventResponse]{
		Data: auditWeb.NewAuditEventResponses(result),
	}

	if err := app.WriteJSON(
		w,
		http.StatusOK,
		response,
		nil,
	); err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}
}
