BEGIN;

CREATE TABLE cka.clients (
    id UUUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    age INT,
    birthday DATE
);

COMMIT;