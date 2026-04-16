CREATE DATABASE IF NOT EXISTS `shorturl`;

USE `shorturl`;

CREATE TABLE IF NOT EXISTS `map` (
    `id`         bigint UNSIGNED  NOT NULL AUTO_INCREMENT,
    `created_at` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` varchar(64)      NOT NULL DEFAULT '',
    `is_del`     tinyint UNSIGNED NOT NULL DEFAULT 0,

    `lurl` varchar(2048) DEFAULT NULL,
    `md5`  char(32)      DEFAULT NULL,
    `surl` varchar(11)   DEFAULT NULL,

    PRIMARY KEY (`id`),
    INDEX       (`is_del`),
    UNIQUE      (`md5`),
    UNIQUE      (`surl`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;