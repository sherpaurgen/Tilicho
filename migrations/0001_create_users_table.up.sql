CREATE TABLE IF NOT EXISTS users_table (
    userid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    groups varchar(255) NOT NULL,
    isactive varchar(255) NOT NULL
);