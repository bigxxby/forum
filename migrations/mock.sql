-- Добавляем категории
INSERT INTO categories (name, posts_count) VALUES
('Technology', 0),
('Science', 0),
('Art', 0);

-- Добавляем посты
INSERT INTO posts (user_id, title, content, created_at, likes, dislikes) VALUES
(1, 'First post', 'Content of the first post', CURRENT_TIMESTAMP, 0, 0),
(1, 'Second post', 'Content of the second post', CURRENT_TIMESTAMP, 0, 0),
(1, 'Third post', 'Content of the third post', CURRENT_TIMESTAMP, 0, 0);

-- Добавляем категории для каждого поста
INSERT INTO posts_categories (post_id, category_id) VALUES

(1, 1), -- Post 1 belongs to category 1
(1, 2), -- Post 1 belongs to category 2
(2, 2), -- Post 2 belongs to category 2
(3, 3); -- Post 3 belongs to category 3

-- Добавляем лайки и дизлайки
INSERT INTO likes_dislikes (user_id, post_id, comment_id, value) VALUES
(1, 1, NULL, TRUE),  -- Лайк на первый пост
(1, 2, NULL, TRUE),  -- Лайк на второй пост
(1, 3, NULL, FALSE), -- Дизлайк на третий пост
(1, NULL, 1, TRUE),  -- Лайк на первый комментарий первого поста
(1, NULL, 2, TRUE),  -- Лайк на второй комментарий первого поста
(1, NULL, 3, FALSE); -- Дизлайк на первый комментарий второго поста
