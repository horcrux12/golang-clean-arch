-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE IF NOT EXISTS category_pkey_seq;
CREATE TABLE IF NOT EXISTS category
(
    id      bigint          NOT NULL DEFAULT nextval('category_pkey_seq'::regclass),
    name    VARCHAR(200)    NOT NULL,
    CONSTRAINT pk_category_id PRIMARY KEY (id)
);

CREATE SEQUENCE IF NOT EXISTS user_pkey_seq;
CREATE TABLE IF NOT EXISTS "user"
(
    id              bigint              NOT NULL DEFAULT nextval('user_pkey_seq'::regclass),
    username		VARCHAR(256)        NOT NULL,
    password		VARCHAR(256),
    user_secret     VARCHAR(256),
    first_name      VARCHAR(256),
    last_name       VARCHAR(256),
    CONSTRAINT pk_user_id PRIMARY KEY (id),
    CONSTRAINT username_unique UNIQUE (username)
);

-- +migrate StatementEnd