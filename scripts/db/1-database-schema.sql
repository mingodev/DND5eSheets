CREATE SCHEMA IF NOT EXISTS `DND5eSheets` DEFAULT CHARACTER SET latin1;
USE `DND5eSheets`

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

DROP TABLE IF EXISTS `races`;
CREATE TABLE IF NOT EXISTS `races` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(75) NOT NULL,
  `speed` INT,
  `size` CHAR(1),
  PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `characters_levels`;
CREATE TABLE IF NOT EXISTS `characters_levels` (
  `id` INT AUTO_INCREMENT,
  `class_id` INT NOT NULL,
  `hit_points`,
  `level_options` JSON,
  PRIMARY KEY(`id`),
  CONSTRAINT `fk_character_level_class` FOREIGN KEY (`class_id`)
  REFERENCES `classes` (`id`)
  ON UPDATE RESTRICT
  ON DELETE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `classes`;
CREATE TABLE `classes` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(75) NOT NULL,
  `description` TEXT(1000) NOT NULL,
PRIMARY KEY(`id`),
UNIQUE(`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `classes_archetypes`;
CREATE TABLE `classes_archetypes` (
  `id` INT AUTO_INCREMENT,
  `parent_class` INT NOT NULL,
  `name` VARCHAR(75) NOT NULL,
  `description` TEXT(1000) NOT NULL,
PRIMARY KEY(`id`),
UNIQUE(`name`),
CONSTRAINT fk_class_archetype FOREIGN KEY(`parent_class`)
REFERENCES `classes` (`id`)
ON UPDATE CASCADE
ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `attributes`;
CREATE TABLE `attributes` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(30) NOT NULL,  
  PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `characters_attributes`;
CREATE TABLE `characters_attributes` (
  `attribute_id` INT NOT NULL,
  `character_id` INT NOT NULL,
  `proficiency` TINYINT(1) DEFAULT 0,
  `score` INT NOT NULL,
  PRIMARY KEY(`attribute_id`, `character_id`),
  CONSTRAINT `fk_character_attribute` FOREIGN KEY(`character_id`)
  REFERENCES `characters` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE,
  CONSTRAINT `fk_attribute_character` FOREIGN KEY(`attribute_id`)
  REFERENCES `attributes` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `skills`;
CREATE TABLE `skills` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(30) NOT NULL,
  `attribute_id` INT NOT NULL,
  PRIMARY KEY(`id`),
  CONSTRAINT `fk_skill_attribute` FOREIGN KEY(`attribute_id`)
  REFERENCES `attributes` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `characters_skills`;
CREATE TABLE `characters_skills` (
  `skill_id` INT NOT NULL,
  `character_id` INT NOT NULL,
  `trained` TINYINT(1) DEFAULT 0,
  PRIMARY KEY(`attribute_id`, `character_id`),
  CONSTRAINT `fk_character_skill` FOREIGN KEY(`character_id`)
  REFERENCES `characters` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE,
  CONSTRAINT `fk_skill_character` FOREIGN KEY(`skill_id`)
  REFERENCES `skills` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `proficiencies_categories`;
CREATE TABLE `proficiencies_categories` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(25) NOT NULL,
  PRIMARY KEY(`id`),
  UNIQUE(`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `characters_proficiencies`;
CREATE TABLE `characters_proficiencies` (
  `id` INT AUTO_INCREMENT,
  `character_id` INT NOT NULL,
  `category_id` INT NOT NULL,
  `value` VARCHAR(25) NOT NULL,
  PRIMARY KEY(`id`),
  CONSTRAINT `fk_character_proficiency` FOREIGN KEY (`character_id`)
  REFERENCES `characters` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE,
  CONSTRAINT `fk_proficiency_category` FOREIGN KEY (`category_id`)
  REFERENCES `proficiencies_categories` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `features`;
CREATE TABLE `features` (
  `id` INT AUTO_INCREMENT,
  `key` VARCHAR(75) NOT NULL,
  `description` TEXT(1000) NOT NULL,
PRIMARY KEY(`id`),
UNIQUE(`key`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `characters_features`;
CREATE TABLE `characters_features` (
  `feature_id` INT NOT NULL,
  `character_id` INT NOT NULL,
  `feature_choices` JSON,
  PRIMARY KEY(`feature_id`, `character_id`),
  CONSTRAINT `fk_character_feature` FOREIGN KEY(`character_id`)
  REFERENCES `characters` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE,
  CONSTRAINT `fk_feature_character` FOREIGN KEY(`feature_id`)
  REFERENCES `features` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE,
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `magic_schools`;
CREATE TABLE `magic_schools` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(20) NOT NULL,
  PRIMARY KEY(`id`),
  UNIQUE(`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `spells`;
CREATE TABLE `spells` (
  `id` INT AUTO_INCREMENT,
  `key` VARCHAR(75) NOT NULL,
  `level` INT NOT NULL,
  `description` TEXT(500) NOT NULL,
`school_id` INT NOT NULL,
`casting_time` VARCHAR(50) NOT NULL,
`range_area` VARCHAR(50) NOT NULL,
`duration` VARCHAR(50) NOT NULL,
`concentration` TINYINT(1) DEFAULT 0,
`ritual` TINYINT(1) DEFAULT 0,
`somatic_component` TINYINT(1) DEFAULT 0,
`verbal_component` TINYINT(1) DEFAULT 0,
`material_component` TINYINT(1) DEFAULT 0,
`material_component_description` VARCHAR(100),
PRIMARY KEY(`id`),
UNIQUE(`key`),
CONSTRAINT `fk_spell_school` FOREIGN KEY(`school_id`)
REFERENCES `magic_schools` (`id`)
ON UPDATE CASCADE
ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `characters_spells`;
CREATE TABLE `characters_spells` (
  `spell_id` INT NOT NULL,
  `character_id` INT NOT NULL,
  `learned` TINYINT(1) NOT NULL,
  `prepared` TINYINT(1) NOT NULL,
  PRIMARY KEY(`spell_id`, `character_id`),
  CONSTRAINT `fk_character_spell` FOREIGN KEY(`character_id`)
  REFERENCES `characters` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE,
  CONSTRAINT `fk_spell_character` FOREIGN KEY (`spell_id`)
  REFERENCES `spells` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
