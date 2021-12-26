DROP TABLE IF EXISTS HELPYOU;

CREATE TABLE HELPYOU (
    helpYouId           CHAR(12) NOT NULL,
    helpMeId            CHAR(12) NOT NULL,
    helpYouUserId       CHAR(12) NOT NULL,
    helpYouDescription  TEXT,
    helpYouCreateTime   DATETIME,
    helpYouChangeTime   DATETIME,
    helpYouState        INT,
    PRIMARY KEY (`helpYouId`)
);

INSERT INTO HELPYOU
(helpYouId,      helpMeId,       helpYouUserId,  helpYouDescription, helpYouCreateTime, helpYouChangeTime, helpYouState)
VALUES
("testhelpyou1", "testhelpme01", "012345678977", "",                 "2021-12-19",      "2021-12-19",      1);
