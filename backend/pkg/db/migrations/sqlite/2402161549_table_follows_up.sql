-- +migrate Up
CREATE TABLE IF NOT EXISTS follows (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    follower_user_id INTEGER NOT NULL,
    following_user_id INTEGER NOT NULL,
    follow_status INTEGER CHECK (follow_status BETWEEN 0 AND 1) NOT NULL, --0 has asked, --1 is following, no entry in that table otherwise
    FOREIGN KEY(follower_user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(following_user_id) REFERENCES users(id) ON DELETE CASCADE
);