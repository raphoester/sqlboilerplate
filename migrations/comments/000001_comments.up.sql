CREATE TABLE comments
(
    id         varchar(36) PRIMARY KEY     NOT NULL,
    created_at timestamp without time zone NOT NULL,
    author_id  varchar(36)                 NOT NULL,
    content    TEXT                        NOT NULL
);

CREATE TABLE replies (
    id         varchar(36) PRIMARY KEY     NOT NULL,
    created_at timestamp without time zone NOT NULL,
    author_id  varchar(36)                 NOT NULL,
    content    TEXT                        NOT NULL,
    comment_id varchar(36)                 NOT NULL
);