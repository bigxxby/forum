INSERT INTO posts (user_id, title, content ,  category_id) VALUES (1, 'Заголовок поста 1', 'Текст поста 1' ,1);
INSERT INTO posts (user_id, title, content ,  category_id) VALUES (1, 'Заголовок поста 2', 'Текст поста 2' , 2);
INSERT INTO posts (user_id, title, content, category_id) VALUES (1, 'Заголовок поста 3', 'Текст поста 3',3);
INSERT INTO posts (user_id, title, content, category_id) VALUES (1, 'Заголовок поста 4', 'Текст поста 4',4);

INSERT INTO categories ( name , posts_count) VALUES ('Рандом каотегория',12);
INSERT INTO categories ( name , posts_count) VALUES ('Аниме',11);
INSERT INTO categories ( name , posts_count) VALUES ('Видеоигры',1);
INSERT INTO categories ( name , posts_count) VALUES ('Тачки',123);

INSERT INTO comments (post_id, user_id, content, created_at)
VALUES 
    (1, 1, 'Это комментарий 1 для поста 1', CURRENT_TIMESTAMP),
    (1, 1, 'Это комментарий 2 для поста 1', CURRENT_TIMESTAMP),
    (1, 1, 'Это комментарий 3 для поста 1', CURRENT_TIMESTAMP),
    (2, 1, 'Это комментарий 1 для поста 2', CURRENT_TIMESTAMP),
    (2, 1, 'Это комментарий 2 для поста 2', CURRENT_TIMESTAMP),
    (2, 1, 'Это комментарий 3 для поста 2', CURRENT_TIMESTAMP),
    (3, 1, 'Это комментарий 1 для поста 3', CURRENT_TIMESTAMP),
    (3, 1, 'Это комментарий 2 для поста 3', CURRENT_TIMESTAMP),
    (3, 1, 'Это комментарий 3 для поста 3', CURRENT_TIMESTAMP),
    (4, 1, 'Это комментарий 1 для поста 4', CURRENT_TIMESTAMP),
    (4, 1, 'Это комментарий 2 для поста 4', CURRENT_TIMESTAMP),
    (4, 1, 'Это комментарий 3 для поста 4', CURRENT_TIMESTAMP);
