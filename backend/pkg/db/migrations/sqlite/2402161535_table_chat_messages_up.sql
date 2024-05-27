-- +migrate Up
CREATE TABLE IF NOT EXISTS chat_messages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_message TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    author_id INTEGER NOT NULL,
    target_id INTEGER NOT NULL,
    message_type INTEGER CHECK (message_type BETWEEN 0 AND 1) NOT NULL, -- 0 for private, 1 for group
    FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE
);