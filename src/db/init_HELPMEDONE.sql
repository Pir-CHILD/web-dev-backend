DROP TABLE IF EXISTS HELPMEDONE;

CREATE TABLE HELPMEDONE (
    helpMeId           CHAR(12) NOT NULL,
    helpMeUserId       CHAR(12) NOT NULL,
    helpYouUserId      CHAR(12) NOT NULL,
    helpYouFinishTime  DATETIME,
    helpMeUserSpend    INT,
    helpYouUserSpend   INT,
    PRIMARY KEY (`helpMeId`)
);

INSERT INTO HELPMEDONE
(helpMeId, helpMeUserId, helpYouUserId, helpYouFinishTime, helpMeUserSpend, helpYouUserSpend)
VALUES
("testhelpme01", "012345678988", "012345678977", "2021-12-19", 3, 1);
