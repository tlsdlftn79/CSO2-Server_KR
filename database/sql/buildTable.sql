Create Table If Not Exists UserData(
    uid INT(10) NOT NULL AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL DEFAULT 'unkown_user',
    gamename VARCHAR(64) NOT NULL DEFAULT 'unkown_user',
    password VARCHAR(64) NOT NULL DEFAULT 'unkown_passwd',
    mail VARCHAR(64) NULL DEFAULT 'unkown',
    data text ,
    PRIMARY KEY(uid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;