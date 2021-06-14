create table UserSurveys
(
    ID bigint auto_increment,
    QuestionID bigint not null,
    AnswerID bigint not null,
    UserID bigint not null,
    constraint UserSurveys_pk
        primary key (ID),
    constraint UserSurveys_Answers_ID_fk
        foreign key (AnswerID) references Answers (ID),
    constraint UserSurveys_Questions_ID_fk
        foreign key (QuestionID) references Questions (ID),
    constraint UserSurveys_Users_ID_fk
        foreign key (UserID) references Users (ID)
);

