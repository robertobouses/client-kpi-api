BEGIN;

CREATE TABLE cka.clients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    age INT,
    birthday DATE,
    telephone_number VARCHAR(15) NOT NULL UNIQUE,

);

COMMIT;