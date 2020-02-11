-- create the databases
CREATE DATABASE IF NOT EXISTS library_db;
USE library_db;

-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS books
(
    id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    title          VARCHAR(255) NOT NULL,
    author         VARCHAR(255) NOT NULL,
    published_date DATE         NOT NULL,
    image_url      VARCHAR(255) NULL,
    description    TEXT         NULL,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (id)
);

INSERT INTO books (title, author, published_date, image_url, description, created_at)
VALUES ('I Too Had a Love Story', 'Ravinder Singh', '2008-01-01', 'https://images-na.ssl-images-amazon.com/images/I/81phwRtlzCL.jpg', 'I Too Had a Love Story is an English autobiographical novel written by Ravinder Singh', NOW());

COMMIT;