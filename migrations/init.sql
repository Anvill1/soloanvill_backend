create table users
(
    id serial not null
        constraint users_pk
            primary key,
    username text,
    email text,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

alter table users owner to soloanvill;

create unique index users_id_uindex
    on users (id);