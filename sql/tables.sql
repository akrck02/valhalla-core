CREATE TABLE user(
	id INTEGER PRIMARY KEY,
	email TEXT NOT NULL,
	profile_pic TEXT,
	insert_date INTEGER NOT NULL,
);

CREATE INDEX idx_user_email ON id(email);

CREATE TABLE userAuth(
	user_id INTEGER NOT NULL,
	password TEXT NOT NULL,
	PRIMARY KEY (user_id, passwords)
);

CREATE TABLE Device(
	id INTEGER PRIMARY KEY,
	user_id INTEGER NOT NULL,
	address TEXT NOT NULL,
	userAgent TEXT NOT NULL,
	token TEXT NOT NULL,
);

CREATE TABLE Project(
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT,
	owner_user INTEGER NOT NULL,
	insert_date INTEGER NOT NULL,
	update_date INTEGER NOT NULL,
);
