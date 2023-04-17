CREATE TABLE organization
(
    bin  text
        primary key,
    name text not null default ''
);

CREATE TABLE city
(
    code             text
        primary key,
    name             text not null default '',
    organization_bin text not null
        constraint city_fk_organization_bin references organization (bin) on delete cascade on update cascade
);

CREATE TABLE device
(
    device_id        text
        primary key,
    token            text not null default '',
    organization_bin text not null
        constraint device_fk_organization_bin references organization (bin) on delete cascade on update cascade
);

CREATE TABLE orders
(
    order_number     text
        primary key,
    created          timestamptz not null default now(),
    modified         timestamptz not null default now(),
    organization_bin text        not null
        constraint orders_fk_organization_bin references organization (bin) on delete cascade on update cascade,
    status           text        not null
);

CREATE TABLE payment
(
    payment_id                   text
        primary key,
    created                      timestamptz not null default now(),
    modified                     timestamptz not null default now(),
    order_number                 text
        constraint payment_fk_order_number references orders (order_number) on delete cascade on update cascade,
    status                       text        not null,
    payment_method               text        not null,
    amount                       numeric not null default 0
);


INSERT INTO organization (name, bin)
VALUES ('Test', '160640004075');
INSERT INTO city (name, organization_bin, code)
VALUES ('Test', '160640004075', 'test');



