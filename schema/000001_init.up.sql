do
$$
    begin
        execute 'ALTER DATABASE ' || current_database() || ' SET timezone = ''+06''';
    end;
$$;

create table city
(
    id      text
        primary key,
    code    text not null default '',
    name    text not null default '',
    org_bin text not null default ''
);

create table src
(
    id         text
        primary key,
    notify_url text not null
);

create table device
(
    id             text
        primary key,
    created        timestamptz not null default now(),
    token          text        not null default '',
    trade_point_id bigint      not null default 0,
    org_bin        text        not null default ''
);

create table ord
(
    id        text
        primary key,
    created   timestamptz not null default now(),
    modified  timestamptz not null default now(),
    src_id    text
        constraint ord_fk_src_id references src (id) on delete cascade on update cascade,
    device_id text        not null
        constraint ord_fk_device_id references device (id) on delete set null on update cascade,
    city_id   text        not null
        constraint ord_fk_city_id references city (id) on delete cascade on update cascade,
    amount    numeric     not null,
    status    text        not null,
    platform  text        not null
);
create index ord_idx_src_id
    on ord (src_id);
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
    id                bigint
        primary key,
    created           timestamptz not null default now(),
    modified          timestamptz not null default now(),
    ord_id            text        not null
        constraint payment_fk_ord_id references ord (id) on delete cascade on update cascade,
    link              text        not null default '',
    status            text        not null,
    status_changed_at timestamptz not null default now(),
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
