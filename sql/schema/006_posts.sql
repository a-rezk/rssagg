-- +goose Up

  CREATE TABLE posts(
    id CHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    published_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    url VARCHAR(2048) UNIQUE NOT NULL,
    feed_id CHAR(36) NOT NULL,
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE

  );

-- +goose Down

  DROP TABLE posts;