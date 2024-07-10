CREATE TABLE users
(
    id             serial      NOT NULL UNIQUE,
    name           varchar(255) NOT NULL,
    username       varchar(255) NOT NULL UNIQUE,
    password  varchar(255) NOT NULL
);

CREATE TABLE sales_lists
(
    id       serial       NOT NULL UNIQUE,
    title    varchar(255) NOT NULL,
    price    numeric      NOT NULL,
    amount   numeric      NOT NULL,
    total    numeric        NOT NULL
);

CREATE TABLE users_lists (
    list_id    integer PRIMARY KEY,
    user_id  integer
);
