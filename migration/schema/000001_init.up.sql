CREATE SCHEMA IF NOT EXISTS amazon;

CREATE TABLE IF NOT EXISTS amazon.categories
(
    id      uuid      default gen_random_uuid() primary key,
    name    varchar(255) not null,
    created timestamp default now(),
    updated timestamp default now(),
    deleted timestamp
);

CREATE TABLE IF NOT EXISTS amazon.products
(
    id          uuid      default gen_random_uuid() primary key,
    name        varchar(255)                           not null,
    price       numeric(15, 2)                         not null,
    sku         varchar(255)                           not null,
    category_id uuid references amazon.categories (id) not null,
    created     timestamp default now(),
    updated     timestamp default now(),
    deleted     timestamp
);

CREATE TABLE IF NOT EXISTS amazon.users
(
    id        uuid      default gen_random_uuid() primary key,
    firstname varchar(127) not null,
    lastname  varchar(127) not null,
    created   timestamp default now(),
    updated   timestamp default now(),
    deleted   timestamp
);

CREATE TABLE IF NOT EXISTS amazon.orders
(
    id          uuid      default gen_random_uuid() primary key,
    user_id     uuid references amazon.users (id) not null,
    total_price numeric(15, 2)                    not null,
    created     timestamp default now(),
    updated     timestamp default now(),
    deleted     timestamp
);

CREATE TABLE IF NOT EXISTS amazon.items
(
    id         uuid      default gen_random_uuid() primary key,
    quantity   int                                  not null,
    product_id uuid references amazon.products (id) not null,
    order_id   uuid references amazon.orders (id)   not null,
    created    timestamp default now(),
    updated    timestamp default now(),
    deleted    timestamp
);

