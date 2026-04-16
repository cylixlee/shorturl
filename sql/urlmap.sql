CREATE DATABASE IF NOT EXISTS `shorturl`;

USE `shorturl`;

CREATE TABLE IF NOT EXISTS `urlmap` (
    `id`         bigint       UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` varchar(64)  NOT NULL DEFAULT '',
    `is_del`     tinyint      UNSIGNED NOT NULL DEFAULT 0,
    `lurl`       varchar(160) DEFAULT NULL,
    `surl`       varchar(11)  DEFAULT NULL,

    PRIMARY KEY (`id`),
    INDEX       (`is_del`),
    UNIQUE      (`lurl`),
    UNIQUE      (`surl`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;