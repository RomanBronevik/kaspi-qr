create table organisation
(
    id   text not null
        primary key,
    bin  text not null default '',
    name text not null default ''
);

create table city
(
    id              text not null
        primary key,
    code            text,
    name            text not null default '',
    organisation_id text not null
        constraint city_fk_organisation_id references organisation (id) on delete cascade on update cascade
);

create table device
(
    id              text
        primary key,
    token           text not null default '',
    organisation_id text not null
        constraint device_fk_organisation_id references organisation (id) on delete cascade on update cascade
);

create table orders
(
    id              text
        primary key,
    created         timestamptz not null default now(),
    modified        timestamptz not null default now(),
    organisation_id text        not null
        constraint orders_fk_organisation_id references organisation (id) on delete cascade on update cascade,
    status          text        not null
);

create table payment
(
    id             text
        primary key,
    created        timestamptz not null default now(),
    modified       timestamptz not null default now(),
    order_id       text        not null
        constraint payment_fk_order_id references orders (id) on delete cascade on update cascade,
    status         text        not null,
    payment_method text        not null,
    amount         numeric     not null default 0
);


insert into organisation (id, name, bin)
values ('test_org', 'test', '160640004075');
insert into city (id, name, organisation_id, code)
values ('test_city', 'test', 'test_org', 'test_city');



