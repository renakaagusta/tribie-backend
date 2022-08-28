CREATE TABLE users
(
    id          VARCHAR PRIMARY KEY,
    username    VARCHAR NOT NULL,
    name        VARCHAR,
    password    VARCHAR,
    apple_id    VARCHAR,
    device_id   VARCHAR,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
);
CREATE TABLE trip
(
    id          VARCHAR PRIMARY KEY,
    title       VARCHAR NOT NULL,
    description VARCHAR,
    place       VARCHAR,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
);
CREATE TABLE trip_member
(
    id         VARCHAR PRIMARY KEY,
    trip_id    VARCHAR,
    user_id    VARCHAR,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
CREATE TABLE transaction
(
    id          VARCHAR PRIMARY KEY,
    trip_id     VARCHAR,
    title       VARCHAR,
    description VARCHAR,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL
);
CREATE TABLE transaction_item
(
    id              VARCHAR PRIMARY KEY,
    trip_id         VARCHAR,
    transaction_id  VARCHAR,
    title           VARCHAR NOT NULL,
    description     VARCHAR,
    price           INT NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL
);
CREATE TABLE transaction_expenses
(
    id              VARCHAR PRIMARY KEY,
    trip_id         VARCHAR,
    trip_member_id  VARCHAR,
    transaction_id  VARCHAR,
    item_id         VARCHAR,
    quantity        INT NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL
);
CREATE TABLE transaction_payment
(
    id              VARCHAR PRIMARY KEY,
    trip_id         VARCHAR,
    trip_member_id  VARCHAR,
    transaction_id  VARCHAR,
    nominal         INT NOT NULL,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL
);
