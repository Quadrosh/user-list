CREATE DATABASE userlist ENCODING 'UTF8';
\connect userlist postgres;
CREATE TABLE users
	 (
		id VARCHAR(255) NOT NULL CONSTRAINT unique_id UNIQUE,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		age INTEGER DEFAULT 0,
		recording_date BIGINT NOT NULL
	
	);