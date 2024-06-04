CREATE TABLE users
(
    id            serial       not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    email         varchar(255) not null unique
);

CREATE TABLE subscriptions
(
    id                 serial                                      not null unique,
    subscriber_user_id int references users (id) on delete cascade not null,
    birthday_user_id   int                                         not null,
    birthday_date      timestamp                                   not null,
    notification_date  timestamp                                   not null
);
