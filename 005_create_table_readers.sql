-- Write your migrate up statements here
create table readers(
    id serial PRIMARY KEY,
    name varchar(255) not null,
    email varchar(255) not null,
    membership_date date not null
);
---- create above / drop below ----
drop table readers;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
