-- RolePermissions table (junction table for many-to-many relationship)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE rolepermission (
    roleid UUID NOT NULL,
    permissionid UUID NOT NULL,
    PRIMARY KEY (roleid, permissionid),
    FOREIGN KEY (roleid) REFERENCES roles(roleid),
    FOREIGN KEY (permissionid) REFERENCES permissions(permissionid)
);