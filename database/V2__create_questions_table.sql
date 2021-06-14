create table Questions
(
    ID bigint auto_increment,
    Question text not null,
    QuestionHash char(40) not null,
    constraint Questions_pk
        primary key (ID)
);

create unique index Questions_QuestionHash_uindex
    on Questions (QuestionHash);

