CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SELECT uuid_generate_v1();


CREATE TABLE IF NOT EXISTS users (
    id           uuid DEFAULT uuid_generate_v4 (),
    first_name   varchar(250) NOT NULL,
    last_name    varchar(250) NOT NULL,
    email        varchar(250) NOT NULL,
    password     varchar NOT NULL,
    is_active    boolean DEFAULT TRUE,               
    created_at   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id)
);