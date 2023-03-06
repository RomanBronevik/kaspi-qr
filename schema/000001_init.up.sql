CREATE TABLE organization (
                        name varchar(255) not null unique,
                        bin varchar(255) not null unique primary key
);

CREATE TABLE city (
                      code varchar(255) not null unique primary key,
                      name varchar(255) not null unique,
                      organization_bin varchar(255) references organization(bin) on delete cascade not null
);

CREATE TABLE device (
                        device_id varchar(255) not null unique primary key,
                        token varchar(255) not null unique,
                        organization_bin varchar(255) references organization (bin) on delete cascade

);
CREATE TABLE orders (
                        created timestamp default null,
                        modified timestamp default null,
                        order_number varchar(255) not null unique primary key,
                        organization_bin varchar(255) references organization (bin) on delete cascade not null,
                        status varchar(255) not null
);

CREATE TABLE payment (
                         created timestamp default null,
                         modified timestamp default null,
                         status varchar(255) not null,
                         order_number varchar(255) references orders (order_number) on delete cascade not null,
                         payment_id varchar(255) not null unique primary key,
                         payment_method varchar(255) not null,
                         wait_timeout timestamp default null,
                         polling_interval int default 5,
                         payment_confirmation_timeout int default 65,
                         amount float
);


INSERT INTO organization (name, bin) VALUES ('Test', '160640004075');
INSERT INTO city (name, organization_bin, code) VALUES ('Test', '160640004075', 'test');



