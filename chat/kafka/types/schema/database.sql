CREATE DATABASE IF NOT EXISTS chatting;

-- room 관련 테이블
CREATE TABLE room (
    `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL UNIQUE,
    `createdAt` timestamp DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

describe room;

-- chat 관련 테이블
CREATE TABLE chat (
    `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `room` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `message` varchar(255) NOT NULL,
    `when` timestamp DEFAULT CURRENT_TIMESTAMP
);

-- server 관리 테이블
CREATE TABLE serverInfo (
    `id` varchar(255) PRIMARY KEY NOT NULL,
    `available` bool NOT NULL
);