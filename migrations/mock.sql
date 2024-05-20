-- Добавляем категории
INSERT INTO categories (name, posts_count) VALUES
('Technology', 0),
('Science', 0),
('Art', 0);

-- Добавляем посты
INSERT INTO posts (user_id, title, content, category_id, created_at, likes, dislikes) VALUES
(1, 'First Post', 'Content of the first post', 1, CURRENT_TIMESTAMP, 5, 2),
(1, 'Second Post', 'Content of the second post', 2, CURRENT_TIMESTAMP, 10, 1),
(1, 'Third Post', 'Content of the third post', 3, CURRENT_TIMESTAMP, 3, 0);

-- Обновляем счетчики постов в категориях
UPDATE categories SET posts_count = (SELECT COUNT(*) FROM posts WHERE category_id = categories.id);

-- Добавляем комментарии
INSERT INTO comments (post_id, user_id, content, created_at, edited, likes, dislikes) VALUES
(1, 1, 'First comment on first post', CURRENT_TIMESTAMP, FALSE, 2, 0),
(1, 1, 'Second comment on first post', CURRENT_TIMESTAMP, FALSE, 1, 0),
(2, 1, 'First comment on second post', CURRENT_TIMESTAMP, FALSE, 0, 1);

-- Добавляем лайки и дизлайки
INSERT INTO likes_dislikes (user_id, post_id, comment_id, value) VALUES
(1, 1, NULL, TRUE),  -- Лайк на первый пост
(1, 2, NULL, TRUE),  -- Лайк на второй пост
(1, 3, NULL, FALSE), -- Дизлайк на третий пост
(1, NULL, 1, TRUE),  -- Лайк на первый комментарий первого поста
(1, NULL, 2, TRUE),  -- Лайк на второй комментарий первого поста
(1, NULL, 3, FALSE); -- Дизлайк на первый комментарий второго поста
