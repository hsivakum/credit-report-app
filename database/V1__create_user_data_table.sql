create table Users
(
	ID bigint auto_increment,
	EMAIL varchar(50) not null,
	PASSWORD char(40) not null,
	CREATED_AT timestamp default current_timestamp not null,
	constraint Users_pk
		primary key (ID)
);

create unique index Users_EMAIL_uindex
	on Users (EMAIL);
