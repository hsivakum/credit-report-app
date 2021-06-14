create table Users
(
	ID bigint auto_increment,
	APP_KEY varchar (36) not null,
	USER_ID varchar(50) not null,
	FIRSTNAME char(40) not null,
	LASTNAME char(40) not null,
	DOB timestamp null,
	SSN varchar (9) null,
	STATE varchar (50) not null,
	STREET varchar (50) not null,
	CITY varchar (50) not null,
	ZIP varchar (50) not null,
	AUTH_TOKEN varchar (36) not null,
	constraint Users_pk
		primary key (ID)
);

create unique index Users_USER_ID_uindex
	on Users (USER_ID);
