CREATE TABLE IF NOT EXISTS users (
    id  INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT,
    login TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    title TEXT,
    content TEXT,
    category_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    likes INTEGER DEFAULT 0,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name  TEXT UNIQUE NOT NULL,
    posts_count INTEGER  
);
CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id  INTEGER,
    user_id TEXT, 
    content TEXT, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    edited BOOLEAN DEFAULT FALSE,
    likes INTEGER DEFAULT 0,
    parent_id INTEGER NULL,
    FOREIGN KEY  (post_id) REFERENCES posts(id),
    FOREIGN KEY  (user_id) REFERENCES users(id)
    FOREIGN KEY  (parent_id) REFERENCES comments(id)
);


CREATE TABLE IF NOT EXISTS likes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    post_id INTEGER,
    comment_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (comment_id) REFERENCES comments(id)
);
