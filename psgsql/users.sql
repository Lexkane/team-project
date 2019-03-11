CREATE TABLE users (
                     id uuid DEFAULT uuid_generate_v1(),
                     name VARCHAR(34) NOT NULL,
                     login VARCHAR(34) NOT NULL,
                     password chkpass,
                     role VARCHAR (16),
                     PRIMARY KEY (id)
);