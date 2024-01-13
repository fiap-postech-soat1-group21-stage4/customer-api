CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP DATABASE IF EXISTS restaurant; 
CREATE DATABASE restaurant;
USE restaurant;

-- customer
DROP TABLE IF EXISTS customer;
CREATE TABLE customer (
	ID  uuid DEFAULT uuid_generate_v4 (),
	NAME VARCHAR(250) NOT NULL,
	CPF VARCHAR(11) NULL unique,
	EMAIL VARCHAR(250) NULL unique,
	CREATED_AT timestamptz NOT NULL DEFAULT now(),
	UPDATED_AT timestamptz NOT NULL DEFAULT now()
)