create table Answers
(
    ID bigint auto_increment,
    `Option` varchar(250) not null,
    IS_CORRECT_ANSWER bool default false not null,
    QuestionID bigint null,
    constraint Answers_pk
        primary key (ID),
    constraint Answers_Questions_ID_fk
        foreign key (QuestionID) references Questions (ID)
);

