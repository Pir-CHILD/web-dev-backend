DROP TABLE IF EXISTS HELPME;

CREATE TABLE HELPME (
    helpMeId          CHAR(12) NOT NULL,
    helpMeUserId      CHAR(12) NOT NULL,
    helpMeType        INT,
    helpMeTheme       VARCHAR(30),
    helpMeDescription TEXT,
    helpMePeople      INT,
    helpMeFinishTime  DATETIME,
    helpMeIntroPhoto  VARCHAR(30),
    helpMeCreateTime  DATETIME,
    helpMeChangeTime  DATETIME,
    helpMeState       INT,
    PRIMARY KEY (`helpMeId`)
);

INSERT INTO HELPME
(helpMeId,       helpMeUserId,   helpMeType, helpMeTheme,  helpMeDescription,           helpMePeople, helpMeFinishTime, helpMeIntroPhoto, helpMeCreateTime, helpMeChangeTime, helpMeState)
VALUES
("testhelpme00", "012345678977", 1,          "banzhongwu", "ban zhong wu help me test", 1,            "2021-12-22",     "",               "2021-12-20",     "2021-12-20",     1);

INSERT INTO HELPME
(helpMeId,       helpMeUserId,   helpMeType, helpMeTheme,        helpMeDescription,                   helpMePeople, helpMeFinishTime, helpMeIntroPhoto, helpMeCreateTime, helpMeChangeTime, helpMeState)
VALUES
("testhelpme01", "012345678988", 2,          "shangxiabandache", "shang xia ban da che help me test", 1,            "2021-12-19",     "",               "2021-12-19",     "2021-12-19",    0);
