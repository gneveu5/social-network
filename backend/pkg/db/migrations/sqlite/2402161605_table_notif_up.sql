-- +migrate Up
CREATE TABLE IF NOT EXISTS notif (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    target_user INTEGER,
    id_one INTEGER,
    id_two INTEGER,
    seen INTEGER CHECK (seen BETWEEN 0 AND 1) DEFAULT 0,
    notif_type INTEGER NOT NULL,
    FOREIGN KEY(target_user) REFERENCES users(id) ON DELETE CASCADE
);