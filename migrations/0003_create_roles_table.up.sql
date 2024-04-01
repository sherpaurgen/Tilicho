-- Roles table with UUID as primary key
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE roles (
    roleid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    rolename VARCHAR(50) NOT NULL UNIQUE,
    description VARCHAR(200) NULL,
    createdon TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modifiedon TIMESTAMP NULL
);