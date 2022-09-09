CREATE TABLE users
(
    id          VARCHAR PRIMARY KEY,
    email       VARCHAR NOT NULL,
    username    VARCHAR,
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
    user_paid_id VARCHAR,
    grand_total INT,
    sub_total INT,
    service_charge INT,
    title       VARCHAR,
    description VARCHAR,
    method      VARCHAR,
    status      VARCHAR,
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
    quantity        INT,
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
    user_from_id    VARCHAR,
    user_to_id      VARCHAR,
    nominal         INT NOT NULL,
    status          VARCHAR,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL
);
