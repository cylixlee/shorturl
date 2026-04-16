CREATE DATABASE IF NOT EXISTS `shorturl`;

USE `shorturl`;

CREATE TABLE IF NOT EXISTS `sequence` (
    `id`        bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `stub`      varchar(1)          NOT NULL,
    `timestamp` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY                 (`id`),
    UNIQUE  KEY `idx_uniq_stub` (`stub`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;