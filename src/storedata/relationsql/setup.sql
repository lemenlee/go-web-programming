DROP TABLE posts CASCADE;
DROP TABLE comments;

CREATE TABLE posts (
    id          serial primary key,
    content     text,
    author      varchar(255)
);

CREATE TABLE comments (
    id          serial primary key,
    content     text,
    author      varchar(255),
    post_id integer references posts(id)
);

