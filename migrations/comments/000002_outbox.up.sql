CREATE TABLE outbox_events
(
    id         varchar(40) PRIMARY KEY     NOT NULL,
    created_at timestamp without time zone NOT NULL,
    topic      TEXT                        NOT NULL,
    payload    JSONB                       NOT NULL
);