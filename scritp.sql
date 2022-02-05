drop database db_logistics;
CREATE DATABASE db_logistics CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
use db_logistics;

create table clients(
    id int unsigned auto_increment primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    phone varchar(100) not null,
    `created_at` datetime default now()
);

insert into clients (name, email, phone) values ('Higor Diego', 'higordiegoti@gmail.com', '88997613741');

create table destinations (
    id int unsigned auto_increment primary key,
    street varchar(255) not null,
    district varchar(255) not null,
    city varchar(255) not null,
    country varchar(255) not null,
    state varchar(255) not null,
    number varchar(255) not null,
    lat varchar(255),
    lng varchar(255),
    id_client int unsigned not null,
	FOREIGN KEY (id_client) REFERENCES clients(id),
    `created_at` datetime default now()
);

insert into destinations (street, district, city, country, state, number, lat, lng, id_client) values ('Padre josé alves', 'Salesianos', 'Juazeiro do Norte', 'Brasil', 'Ceará', '790', '-7.205440', '-39.324280', 1);

create table products (
    id int unsigned auto_increment primary key,
    name varchar(255) not null,
    price decimal(10, 2) not null,
    id_client int unsigned not null,
	FOREIGN KEY (id_client) REFERENCES clients(id),
    `created_at` datetime default now()
);

insert into products (`name`, price, id_client) values ('Iphone 12', 4900.00, 1);


create table carryings (
    id int unsigned auto_increment primary key,
    name varchar(255) not null,
    lat varchar (255) not null,
    lng varchar (255) not null,
    `created_at` datetime default now()
);

insert into carryings (name, lat, lng) values ('Transportadora example', '-9.752860', '-36.665240');

create table trucks (
    id int unsigned auto_increment primary key,
    brand varchar(255) not null,
    model varchar(255) not null,
    year varchar(255) not null,
    plate varchar(255) not null,
    `created_at` datetime default now()
);

insert into trucks (brand, model, year, plate) values ('Mercedes-Benz', 'Accelo', '2021-2022', 'PNO-7672');

create table drivers (
    id int unsigned auto_increment primary key,
    name varchar(255) not null,
    image varchar(255) not null,
    id_carring int unsigned not null,
	FOREIGN KEY (id_carring) REFERENCES carryings(id),
    id_truck int unsigned not null,
	FOREIGN KEY (id_truck) REFERENCES trucks(id),
    `created_at` datetime default now()
);

insert into drivers (name, image, id_carring, id_truck) values ('Motorista teste', 'https://cdn-icons-png.flaticon.com/512/147/147144.png', 1, 1);

create table couriers (
    id int unsigned auto_increment primary key,
    id_driver int unsigned not null,
	FOREIGN KEY (id_driver) REFERENCES drivers(id),
    id_product int unsigned not null,
	FOREIGN KEY (id_product) REFERENCES products(id),
    `created_at` datetime default now()
);