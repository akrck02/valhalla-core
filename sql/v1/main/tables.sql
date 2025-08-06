CREATE TABLE database_metadata(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	version INTEGER NOT NULL
);

CREATE TABLE user(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL,
	profile_pic TEXT,
	password TEXT NOT NULL,
	database TEXT NOT NULL,
	insert_date INTEGER NOT NULL
);

CREATE INDEX idx_user_email ON user(email);

CREATE TABLE device(
	user_id INTEGER NOT NULL,
	address TEXT NOT NULL,
	user_agent TEXT NOT NULL,
	token TEXT NOT NULL,
	insert_date INTEGER NOT NULL,
	update_date INTEGER NOT NULL,
	PRIMARY KEY (user_id, address, user_agent),
 	FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE shared_project(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT,
	insert_date INTEGER NOT NULL,
	update_date INTEGER NOT NULL
);

CREATE TABLE shared_task(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	project_id INTEGER NOT NULL,
	name TEXT NOT NULL,
	description TEXT,
	insert_date INTEGER NOT NULL,
	update_date INTEGER NOT NULL,
 	FOREIGN KEY (project_id) REFERENCES shared_project(id)
);

CREATE TABLE project_member(
	project_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	PRIMARY KEY (project_id, user_id),
 	FOREIGN KEY (project_id) REFERENCES shared_project(id),
   	FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE project_member_permissions(
	project_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	permission INTEGER NOT NULL,
	PRIMARY KEY (project_id, user_id, permission),
 	FOREIGN KEY (project_id) REFERENCES shared_project(id),
   	FOREIGN KEY (user_id) REFERENCES user(id)
);
