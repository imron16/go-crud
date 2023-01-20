CREATE DATABASE IF NOT EXISTS dbcoba;

use dbcoba;

create table IF NOT EXISTS term_and_condition
(
  id bigint NOT NULL AUTO_INCREMENT,
  title varchar(255) DEFAULT NULL,
  title_en varchar(255) DEFAULT NULL,
  content varchar(255) DEFAULT NULL,
  content_en varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
 ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
);