DROP TABLE IF EXISTS USERINFO;

CREATE TABLE USERINFO (
    userId           CHAR(12) NOT NULL,
    userName         VARCHAR(30) NOT NULL,
    userPwd          VARCHAR(30) NOT NULL,
    userType         INT NOT NULL,
    userXingming     VARCHAR(30),
    certiType        INT NOT NULL,
    certiNumber      VARCHAR(30) NOT NULL,
    userPhone        CHAR(11) NOT NULL,
    userLevel        INT NOT NULL,
    userIntroduction TEXT,
    userCity         VARCHAR(30) NOT NULL,
    userCommunity    INT NOT NULL,
    registerTime     DATETIME NOT NULL,
    changeTime       DATETIME NOT NULL,
    PRIMARY KEY (userName, userPwd)
);

INSERT INTO USERINFO
(userId,         userName, userPwd,   userType, userXingming,      certiType, certiNumber,  userPhone,     userLevel, userIntroduction, userCity,  userCommunity, registerTime, changeTime)
VALUES
('adminadmin77', 'admin',  'admin77', 0,        "admin",           0,         "2018211619", "19801332707", 1,         "",               "Beijing", 0,             "2021-01-01", "2021-01-01");

INSERT INTO USERINFO
(userId,         userName,   userPwd,  userType, userXingming,   certiType, certiNumber,   userPhone,    userLevel, userIntroduction, userCity,  userCommunity, registerTime, changeTime)
VALUES
('012345678977', 'test77',   'test77', 1,        "77",           0,         "2018111111", "13311111111", 0,         "",               "Beijing", 0,             "2021-12-01", "2021-12-01");

INSERT INTO USERINFO
(userId,         userName,   userPwd,  userType, userXingming,   certiType, certiNumber,   userPhone,    userLevel, userIntroduction, userCity,  userCommunity, registerTime, changeTime)
VALUES
('012345678988', 'test88',   'test88', 1,        "88",           0,         "2018222222", "13322222222", 0,         "",               "Beijing", 0,             "2021-12-01", "2021-12-01");
