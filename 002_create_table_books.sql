-- Write your migrate up statements here
create table books (
    id serial primary key,
    title varchar(255) NOT NULL, 
    author_id varchar(255) NOT NULL,
    published_date DATE NOT NULL
    isbn varchar(13) NOT NULL,
    available boolean DEFAULT TRUE,
    FOREIGN KEY (author_id) REFERENCES authors(id)
);
---- create above / drop below ----
drop table book;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
