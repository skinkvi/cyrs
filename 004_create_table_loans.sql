-- Write your migrate up statements here
create table loans(
    id serial primary key,
    book_id int not null,
    reader_id int not null,
    loan_date date not null,
    return_date date,
    FOREIGN KEY(book_id) REFERENCES books(id),
    FOREIGN KEY(reader_id) REFERENCES readers(id)
);
---- create above / drop below ----
drop table loans;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
