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
    locale          VARCHAR(15),
    CONSTRAINT pk_user_id PRIMARY KEY (id),
    CONSTRAINT username_unique UNIQUE (username)
);

CREATE SEQUENCE IF NOT EXISTS auth_refresh_token_pkey_seq;
CREATE TABLE IF NOT EXISTS refresh_token
(
    id              bigint              NOT NULL DEFAULT nextval('auth_refresh_token_pkey_seq'::regclass),
    user_id		    bigint              NOT NULL,
    refresh_token	VARCHAR(256)        NOT NULL,
    CONSTRAINT pk_refresh_token_id PRIMARY KEY (id),
    CONSTRAINT fk_refreshtoken_user FOREIGN KEY (user_id) REFERENCES "user"(id)
);

-- +migrate StatementEnd