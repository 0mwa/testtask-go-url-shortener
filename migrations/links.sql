create table links
(
    id           serial
        constraint links_pk
            primary key,
    original_url varchar(255) not null,
    short_url    char(6)      not null
        constraint links_pk_2
            unique
);

alter table links
    owner to postgres;