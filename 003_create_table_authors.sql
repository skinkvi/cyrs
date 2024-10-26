-- Write your migrate up statements here
create table authors(
    id serial primary key,
    name varchar(255) not null,
    birth_date DATE
);
---- create above / drop below ----
drop table authors;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
