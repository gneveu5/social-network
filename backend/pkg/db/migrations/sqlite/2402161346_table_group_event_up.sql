-- +migrate Up
CREATE TABLE IF NOT EXISTS group_event (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(30) NOT NULL,
    event_description VARCHAR(200) NOT NULL,
    event_time TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER NOT NULL,
    users_group_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(users_group_id) REFERENCES users_group(id) ON DELETE CASCADE
);