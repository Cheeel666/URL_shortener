create database test;

create table links(
id_url serial primary key,
url varchar not null,
shorten_url varchar
);
