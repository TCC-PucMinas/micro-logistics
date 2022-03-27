drop database if exists db_logistics;

CREATE DATABASE db_logistics CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
use db_logistics;

create table carryings (
   id int unsigned auto_increment primary key,
   name varchar(255) not null,
   street varchar(255) not null,
   district varchar(255) not null,
   city varchar(255) not null,
   country varchar(255) not null,
   state varchar(255) not null,
   number varchar(255) not null,
   zipCode varchar(255) not null,
   lat varchar (255) not null,
   lng varchar (255) not null,
   `created_at` datetime default now()
);

insert into carryings (name, street, district, city, country, state, number, zipCode, lat, lng) values ('Transportadora example', 'Padre josé alves', 'Salesianos', 'Juazeiro do Norte', 'Brasil', 'Ceará', '790', '63050222','-7.205440', '-39.324280');

create table deposits(
     id int unsigned auto_increment primary key,
     name varchar(255) not null,
     street varchar(255) not null,
     district varchar(255) not null,
     city varchar(255) not null,
     country varchar(255) not null,
     state varchar(255) not null,
     number varchar(255) not null,
     zipCode varchar(255) not null,
     lat varchar (255) not null,
     lng varchar (255) not null,
     id_carry int unsigned not null,
     FOREIGN KEY (id_carry) REFERENCES carryings(id)
);

insert into deposits (name,street, district, city, country, state, number, lat, lng, zipCode, id_carry) values ('deposito exemplo', 'Padre josé alves', 'Salesianos', 'Juazeiro do Norte', 'Brasil', 'Ceará', '790', '-7.205440', '-39.324280', '63050222',1);


create table trucks (
    id int unsigned auto_increment primary key,
    brand varchar(255) not null,
    model varchar(255) not null,
    year varchar(255) not null,
    plate varchar(255) not null,
    id_carry int unsigned not null,
    FOREIGN KEY (id_carry) REFERENCES carryings(id),
    `created_at` datetime default now()
);

insert into trucks (brand, model, year, plate, id_carry) values ('Mercedes-Benz', 'Accelo', '2021-2022', 'PNO-7672', 1);

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
    id_deposit int unsigned not null,
    id_client int unsigned not null,
    id_product int unsigned not null,
    delivered boolean default false,
    doc JSON,
    FOREIGN KEY (id_driver) REFERENCES drivers(id),
    FOREIGN KEY (id_deposit) REFERENCES deposits(id),
   `created_at` datetime default now()
);

create table courier_routes(
    id int unsigned auto_increment primary key,
    id_courier int unsigned not null,
    `order` int default 0,
    `latInit` JSON not null,
    `latFinish` JSON not null,
    FOREIGN KEY (id_courier) REFERENCES couriers(id),
   `created_at` datetime default now()
);
