CREATE TABLE IF NOT EXISTS users_table (
    userid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    groups varchar(255) NOT NULL,
    isactive varchar(255) NOT NULL
);

-- INSERT INTO users_table (username, email, password, groups, isactive)
-- VALUES
--     ('john_doe', 'john@example.com', 'password123', 'admin', 'true'),
--     ('jane_smith', 'jane@example.com', 'securepass', 'user', 'true'),
--     ('alice_green', 'alice@example.com', 'p@ssw0rd', 'user', 'false');