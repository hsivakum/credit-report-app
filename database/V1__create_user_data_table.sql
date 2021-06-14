create table Users
(
    ID         bigint  auto_increment,
    EMAIL      varchar(50) not null,
    PASSWORD   varchar(50) not null,
    CREATED_AT TIMESTAMP   not null,
    constraint Users_pk
        primary key (ID)
);

create unique index Users_ID_uindex
    on Users (ID);

