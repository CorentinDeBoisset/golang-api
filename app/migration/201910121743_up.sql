CREATE TABLE user_account (
	id serial,
	username VARCHAR (50) UNIQUE NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	password_hash VARCHAR(255),
	first_name VARCHAR(255),
	last_name VARCHAR (255),
	created_on TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
	last_login TIMESTAMP(0),
	PRIMARY KEY(id)
)
