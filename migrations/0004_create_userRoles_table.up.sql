-- UserRoles table (junction table for many-to-many relationship)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE userroles (
    userid INT NOT NULL FOREIGN KEY REFERENCES Users(UserID),
    roleid INT NOT NULL FOREIGN KEY REFERENCES Roles(RoleID),
    PRIMARY KEY (UserID, RoleID)
);
