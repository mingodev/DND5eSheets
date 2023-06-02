CREATE SCHEMA IF NOT EXISTS `DND5eSheets` DEFAULT CHARACTER SET latin1;
USE `DND5eSheets`

-- TODO : Remove once done testing
DROP TABLE IF EXISTS `characters`;

CREATE TABLE IF NOT EXISTS `characters` (
    `id` INT AUTO_INCREMENT,
    `name` VARCHAR(75) NOT NULL,
    `race_id` INT NOT NULL,
    `hit_points` INT,
    `hit_points_max` INT,
    `speed` INT,
    PRIMARY KEY(`id`),
    CONSTRAINT `fk_race_id_characters` FOREIGN KEY (`race_id`) 
        REFERENCES `races` (`id`)
        ON UPDATE RESTRICT
        ON DELETE RESTRICT

) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- TODO : Remove once done testing
DROP TABLE IF EXISTS `races`;

CREATE TABLE IF NOT EXISTS `races` (
    `id` INT AUTO_INCREMENT,
    `name` VARCHAR(75) NOT NULL,
    `speed` INT,
    `size` CHAR(1)
    PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;