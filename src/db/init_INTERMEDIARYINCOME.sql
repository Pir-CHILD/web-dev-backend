DROP TABLE IF EXISTS INTERMEDIARYINCOME;

CREATE TABLE INTERMEDIARYINCOME (
    month CHAR(7) NOT NULL,
    provinceCity VARCHAR(30),
    community INT,
    helpMeType INT,
    finishNumber INT,
    interIncome INT,
    PRIMARY KEY (`month`, `community`, `helpMeType`)
);

INSERT INTO INTERMEDIARYINCOME
(month, provinceCity, community, helpMeType, finishNumber, interIncome)
VALUES
("2021-11", "Beijing", 0, 0, 10, 40);