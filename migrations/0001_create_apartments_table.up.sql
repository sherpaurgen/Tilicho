CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS apartments (
    apartment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    apartment_name varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    city varchar(255) NOT NULL,
    state varchar(255) NOT NULL,
    zip_code varchar(255) NOT NULL
);





