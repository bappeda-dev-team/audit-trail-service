CREATE TABLE audit_event (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    project_id VARCHAR(100) NOT NULL,

    event_type VARCHAR(50) NOT NULL,
    action VARCHAR(30) NOT NULL,

    service_name VARCHAR(100) NOT NULL,

    entity_type VARCHAR(100) NOT NULL,
    entity_id VARCHAR(255) NOT NULL,

    actor_id VARCHAR(255),
    actor_username VARCHAR(255),

    request_id VARCHAR(255),
    session_id VARCHAR(255),

    occurred_at TIMESTAMPTZ NOT NULL,

    before_data JSONB,
    after_data JSONB,

    metadata JSONB,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_audit_event_project_entity
    ON audit_event(project_id, entity_type, entity_id);

CREATE INDEX idx_audit_event_project_actor
    ON audit_event(project_id, actor_id);

CREATE INDEX idx_audit_event_project_service
    ON audit_event(project_id, service_name);

CREATE INDEX idx_audit_event_project_occurred
    ON audit_event(project_id, occurred_at DESC);

CREATE INDEX idx_audit_event_request
    ON audit_event(request_id);

CREATE INDEX idx_audit_event_session
    ON audit_event(session_id);
