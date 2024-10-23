CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

CREATE TABLE IF NOT EXISTS usuarios (
    id int auto_increment primary key,
    nome varchar(100) not null,
    nick varchar(100) not null unique,
    email varchar(100) not null unique,
    senha varchar(100) not null,
    criado_em timestamp default current_timestamp()
 ) ENGINE=INNODB;
