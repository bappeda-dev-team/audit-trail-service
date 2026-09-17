CREATE TABLE audit_event (
    id UUID PRIMARY KEY,

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

CREATE INDEX idx_audit_event_entity
ON audit_event(entity_type, entity_id);

CREATE INDEX idx_audit_event_actor
ON audit_event(actor_id);

CREATE INDEX idx_audit_event_service
ON audit_event(service_name);

CREATE INDEX idx_audit_event_occurred_at
ON audit_event(occurred_at);

CREATE INDEX idx_audit_event_request
ON audit_event(request_id);
