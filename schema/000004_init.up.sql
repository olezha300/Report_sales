CREATE TABLE users
(
    id             serial      NOT NULL UNIQUE,
    name           varchar(255) NOT NULL,
    username       varchar(255) NOT NULL UNIQUE,
    password_hash  varchar(255) NOT NULL
);

CREATE TABLE sales_lists
(
    id       serial       NOT NULL UNIQUE,
    title    varchar(255) NOT NULL,
    price    numeric      NOT NULL,
    amount   numeric      NOT NULL
);

CREATE TABLE users_lists
(
    id      serial                                       NOT NULL UNIQUE,
    user_id int REFERENCES users (id) ON DELETE CASCADE  NOT NULL,
    list_id int REFERENCES sales_lists (id) ON DELETE CASCADE NOT NULL
);