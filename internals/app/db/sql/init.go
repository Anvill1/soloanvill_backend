package sql

var InitSQL = `
create table users
(
    id serial not null
        constraint users_pk
            primary key,
    username text,
    email text,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ip VARCHAR(20)
);

alter table users owner to soloanvill;

create unique index users_id_uindex
    on users (id);
`
