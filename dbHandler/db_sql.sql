-- Database: motivatio

-- DROP DATABASE IF EXISTS motivatio;

CREATE DATABASE motivatio
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United Kingdom.1250'
    LC_CTYPE = 'English_United Kingdom.1250'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

COMMENT ON DATABASE motivatio
    IS 'motivatio local db';
	
CREATE TABLE goals (
	id serial PRIMARY KEY,
	title VARCHAR(50),
	description VARCHAR(300),
	start_date Date,
	relevancy REAL
)

CREATE TABLE tags (
	id serial PRIMARY KEY,
	tag_name VARCHAR(50),
	color INTEGER
)

CREATE TABLE notifications (
	id serial PRIMARY KEY,
	notification_type VARCHAR(50)
) 





