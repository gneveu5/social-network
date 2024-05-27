-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	nickname VARCHAR(15),
	email VARCHAR(100) NOT NULL,
	user_password VARCHAR(100) NOT NULL,
	first_name VARCHAR(30) NOT NULL,
  	last_name VARCHAR(30) NOT NULL,
	date_of_birth TEXT NOT NULL,
	about_me VARCHAR(200),
	public_private INTEGER CHECK (public_private BETWEEN 0 AND 1) NOT NULL, --can't register bool directly, this is the best approche
	avatar VARCHAR(255),
  	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	has_validated INTEGER CHECK (has_validated BETWEEN 0 AND 1) NOT NULL
);
