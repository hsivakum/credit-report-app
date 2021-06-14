create table CreditOrders
(
    ID bigint auto_increment,
    UserID bigint not null,
    ProductCode varchar(10) not null,
    ReportKey varchar(36) not null,
    constraint CreditOrders_pk
        primary key (ID),
    constraint CreditOrders_Users_ID_fk
        foreign key (UserID) references Users (ID)
);

