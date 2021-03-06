DROP DATABASE IF EXISTS Trill;

CREATE DATABASE Trill CHARACTER SET utf8;

USE Trill;

CREATE TABLE User(
    id INT AUTO_INCREMENT PRIMARY KEY,
    str_firstName VARCHAR(100) NOT NULL,
    str_lastName VARCHAR(100) NOT NULL,
    str_email VARCHAR(60) NOT NULL UNIQUE,
    str_username VARCHAR(16) NOT NULL UNIQUE,
    str_password VARCHAR(100) NOT NULL
);

CREATE TABLE Board(
    id INT AUTO_INCREMENT PRIMARY KEY,
    str_name VARCHAR(100) NOT NULL,
    str_backgroundColor VARCHAR(30) NOT NULL DEFAULT "gray" 
);

CREATE TABLE Team(
    id INT AUTO_INCREMENT PRIMARY KEY,
    str_name VARCHAR(100) NOT NULL,
    str_code VARCHAR(6) NOT NULL
);

CREATE TABLE List(
    id INT AUTO_INCREMENT PRIMARY KEY,
    str_name VARCHAR(100) NOT NULL,
    enu_status ENUM("Archived", "Active") DEFAULT "Active"
);

CREATE TABLE Tag(
    id INT AUTO_INCREMENT PRIMARY KEY,
    str_name VARCHAR(100) NOT NULL,
    str_color VARCHAR(30) NOT NULL DEFAULT "gray"
);

CREATE TABLE Card(
    id int AUTO_INCREMENT PRIMARY KEY,
    str_name VARCHAR(100) NOT NULL,
    str_description VARCHAR(350) NOT NULL,
    dat_startDate DATE NOT NULL DEFAULT NOW(),
    dat_endDate DATE NOT NULL
);

CREATE TABLE UserBoard(
    id_user INT NOT NULL,
    id_board INT NOT NULL,
    FOREIGN KEY (id_user) REFERENCES User(id) ON DELETE CASCADE,
    FOREIGN KEY (id_board) REFERENCES Board(id) ON DELETE CASCADE
);

CREATE TABLE BoardMember(
    id_board INT NOT NULL,
    id_member INT NOT NULL,
    FOREIGN KEY (id_board) REFERENCES Board(id) ON DELETE CASCADE,
    FOREIGN KEY (id_member) REFERENCES User(id) ON DELETE CASCADE
);

CREATE TABLE TeamBoard(
    id_team INT NOT NULL,
    id_board INT NOT NULL,
    FOREIGN KEY (id_team) REFERENCES Team(id) ON DELETE CASCADE,
    FOREIGN KEY (id_board) REFERENCES Board(id) ON DELETE CASCADE
);

CREATE TABLE UserTeam(
    id_team INT NOT NULL,
    id_user INT NOT NULL,
    FOREIGN KEY (id_team) REFERENCES Team(id) ON DELETE CASCADE,
    FOREIGN KEY (id_user) REFERENCES User(id) ON DELETE CASCADE
);

CREATE TABLE BoardList(
    id_list INT NOT NULL,
    id_board INT NOT NULL,
    FOREIGN KEY (id_list) REFERENCES List(id) ON DELETE CASCADE,
    FOREIGN KEY (id_board) REFERENCES Board(id) ON DELETE CASCADE
);

CREATE TABLE ListCard(
    id_list INT NOT NULL,
    id_card INT NOT NULL,
    FOREIGN KEY (id_list) REFERENCES List(id) ON DELETE CASCADE,
    FOREIGN KEY (id_card) REFERENCES Card(id) ON DELETE CASCADE
);

CREATE TABLE CardMember(
    id_card INT NOT NULL,
    id_member INT NOT NULL,
    FOREIGN KEY (id_card) REFERENCES Card(id) ON DELETE CASCADE,
    FOREIGN KEY (id_member) REFERENCES User(id) ON DELETE CASCADE
);

CREATE TABLE CardTag(
    id_card INT NOT NULL,
    id_tag INT NOT NULL,
    FOREIGN KEY (id_card) REFERENCES Card(id) ON DELETE CASCADE,
    FOREIGN KEY (id_tag) REFERENCES Tag(id) ON DELETE CASCADE
);

