-- File name   : book-store.sql --
-- Author      : NRaghuraj      --
-- Date Created: 2023 MAR 23    --
-- Description : Database with book table consisting of entry fields: {id(Primary Key), title, isbn and author} --
-- Host: localhost, Port: 8080  Database: book-store (MySQL) --
--
-- Table structure for table `book_table`--
--
-- Check if table already exists --
DROP TABLE IF EXISTS `book_table`;      

CREATE TABLE book_table (
    id VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    isbn VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

-- Insert entries into the table to test API calls -- 
INSERT INTO book_table (id, title, isbn, author)
VALUES ('1', 'To Kill a Mockingbird', '9780061120084', 'Harper Lee'),
       ('2', '1984', '9780451524935', 'George Orwell'),
       ('3', 'Pride and Prejudice', '9780486284736', 'Jane Austen'),
       ('4', 'The Catcher in the Rye', '9780316769488', 'J.D. Salinger'),
       ('5', 'To Kill a Mockingbird', '9780446310789', 'Harper Lee'),
       ('6', '1984', '9780451524935', 'George Orwell'),
       ('7', 'Pride and Prejudice', '9780141439518', 'Jane Austen'),
       ('8', 'The Great Gatsby', '9780743273565', 'F. Scott Fitzgerald'),
       ('9', 'The Hobbit', '9780547928227', 'J.R.R. Tolkien'),
       ('10', 'The Lord of the Rings', '9780618640157', 'J.R.R. Tolkien');

-- Dump completed on 2022-03-25 17:15:35 --