-- Users table
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users_table (
    userid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    isactive BOOLEAN NOT NULL DEFAULT TRUE,
    createdon TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modifiedon TIMESTAMP NULL
    );


-- INSERT INTO users_table (username, email, password, groups, isactive)
-- VALUES
--     ('john_doe', 'john@example.com', 'password123', 'admin', 'true'),
--     ('jane_smith', 'jane@example.com', 'securepass', 'user', 'true'),
--     ('alice_green', 'alice@example.com', 'p@ssw0rd', 'user', 'false');