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
    id      text
        primary key,
    token   text not null default '',
    org_bin text not null default ''
);

create table ord
(
    id       text
        primary key,
    created  timestamptz not null default now(),
    modified timestamptz not null default now(),
    org_bin  text        not null default '',
    status   text        not null
);

create table payment
(
    id                          text
        primary key,
    created                     timestamptz not null default now(),
    modified                    timestamptz not null default now(),
    ord_id                      text        not null
        constraint payment_fk_ord_id references ord (id) on delete cascade on update cascade,
    status                      text        not null,
    payment_method              text        not null,
    expire_dt                   timestamptz          default null,
    qr_payment_behavior_options jsonb       not null default '{}',
    amount                      numeric     not null default 0
);



