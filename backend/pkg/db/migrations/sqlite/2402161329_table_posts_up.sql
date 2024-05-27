-- +migrate Up
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(30) NOT NULL,
    body TEXT NOT NULL,
    img VARCHAR(255),
    author_id INTEGER,
    group_id INTEGER, -- 0 if post global
    view_status INTEGER CHECK (view_status BETWEEN 0 AND 2) NOT NULL, --0 public, 1 semi public, 2 private
  	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(group_id) REFERENCES users_group(id) ON DELETE CASCADE
);