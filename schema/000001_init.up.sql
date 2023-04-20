do
$$
    begin
        execute 'ALTER DATABASE ' || current_database() || ' SET timezone = ''+06''';
    end;
$$;

create table city
(
    id      text not null
        primary key,
    code    text not null default '',
    name    text not null default '',
    org_bin text not null default ''
);

create table device
(
    id             text
        primary key,
    token          text   not null default '',
    trade_point_id bigint not null default 0,
    org_bin        text   not null default ''
);

create table ord
(
    id        text
        primary key,
    created   timestamptz not null default now(),
    modified  timestamptz not null default now(),
    src       text        not null,
    device_id text        not null
        constraint ord_fk_device_id references device (id) on delete set null on update cascade,
    city_id   text        not null
        constraint ord_fk_city_id references city (id) on delete cascade on update cascade,
    amount    numeric     not null,
    status    text        not null,
    platform  text        not null
);
create index ord_idx_src
    on ord (src);
create index ord_idx_device_id
    on ord (device_id);
create index ord_idx_city_id
    on ord (city_id);
create index ord_idx_status
    on ord (status);
create index ord_idx_platform
    on ord (platform);

create table payment
(
    id                text
        primary key,
    created           timestamptz not null default now(),
    modified          timestamptz not null default now(),
    ord_id            text        not null
        constraint payment_fk_ord_id references ord (id) on delete cascade on update cascade,
    status            text        not null,
    status_changed_at timestamptz not null default now(),
    payment_method    text        not null,
    amount            numeric     not null default 0,
    expire_dt         timestamptz          default null,
    pbo               jsonb       not null default '{}'
);
create index payment_idx_ord_id
    on payment (ord_id);
create index payment_idx_status
    on payment (status);
create index payment_idx_status_changed_at
    on payment (status_changed_at);
