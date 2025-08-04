
CREATE TABLE database_metadata(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	version INTEGER NOT NULL
);

CREATE TABLE user(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL,
	profile_pic TEXT,
	insert_date INTEGER NOT NULL
);

CREATE INDEX idx_user_email ON user(email);

CREATE TABLE user_auth(
	user_id INTEGER NOT NULL,
	password TEXT NOT NULL,
	PRIMARY KEY (user_id, password)
);

CREATE TABLE device(
	user_id INTEGER NOT NULL,
	address TEXT NOT NULL,
	user_agent TEXT NOT NULL,
	token TEXT NOT NULL,
	PRIMARY KEY (user_id, address, user_agent)
);

CREATE TABLE project(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT,
	owner_user INTEGER NOT NULL,
	insert_date INTEGER NOT NULL,
	update_date INTEGER NOT NULL
);
