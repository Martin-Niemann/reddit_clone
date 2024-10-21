-- MySQL Script generated by MySQL Workbench
-- 2024-10-14T16:17:54 CEST
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema reddit_clone
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema reddit_clone
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `reddit_clone` ;
USE `reddit_clone` ;

-- -----------------------------------------------------
-- Table `reddit_clone`.`user`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `reddit_clone`.`user` ;

CREATE TABLE IF NOT EXISTS `reddit_clone`.`user` (
  `id_user` INT NOT NULL AUTO_INCREMENT,
  `user_name` VARCHAR(100) NOT NULL,
  `display_name` VARCHAR(200) NULL,
  `email` VARCHAR(320) NOT NULL,
  `hashed_password` VARCHAR(256) NOT NULL,
  PRIMARY KEY (`id_user`),
  UNIQUE INDEX `id_user_UNIQUE` (`id_user` ASC) VISIBLE,
  UNIQUE INDEX `user_name_UNIQUE` (`user_name` ASC) VISIBLE,
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `reddit_clone`.`subreddit`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `reddit_clone`.`subreddit` ;

CREATE TABLE IF NOT EXISTS `reddit_clone`.`subreddit` (
  `id_subreddit` INT NOT NULL AUTO_INCREMENT,
  `url` VARCHAR(100) NOT NULL,
  `name` VARCHAR(100) NOT NULL,
  `description` VARCHAR(1000) NULL,
  `id_user` INT NULL,
  PRIMARY KEY (`id_subreddit`),
  UNIQUE INDEX `id_UNIQUE` (`id_subreddit` ASC) VISIBLE,
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE,
  UNIQUE INDEX `url_UNIQUE` (`url` ASC) VISIBLE,
  INDEX `owner_idx` (`id_user` ASC) VISIBLE,
  CONSTRAINT `subreddit_owner`
    FOREIGN KEY (`id_user`)
    REFERENCES `reddit_clone`.`user` (`id_user`)
    ON DELETE SET NULL
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `reddit_clone`.`post`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `reddit_clone`.`post` ;

CREATE TABLE IF NOT EXISTS `reddit_clone`.`post` (
  `id_post` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(300) NOT NULL,
  `link` VARCHAR(1000) NULL,
  `text` VARCHAR(10000) NULL,
  `created_date` DATETIME NOT NULL,
  `edited_date` DATETIME NULL,
  `id_subreddit` INT NOT NULL,
  `id_user` INT NOT NULL,
  PRIMARY KEY (`id_post`),
  UNIQUE INDEX `idpost_UNIQUE` (`id_post` ASC) VISIBLE,
  UNIQUE INDEX `id_subreddit_UNIQUE` (`id_subreddit` ASC) VISIBLE,
  INDEX `user_idx` (`id_user` ASC) VISIBLE,
  CONSTRAINT `post_subreddit`
    FOREIGN KEY (`id_subreddit`)
    REFERENCES `reddit_clone`.`subreddit` (`id_subreddit`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `post_user`
    FOREIGN KEY (`id_user`)
    REFERENCES `reddit_clone`.`user` (`id_user`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `reddit_clone`.`comment`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `reddit_clone`.`comment` ;

CREATE TABLE IF NOT EXISTS `reddit_clone`.`comment` (
  `id_comment` INT NOT NULL AUTO_INCREMENT,
  `created_date` DATETIME NOT NULL,
  `edited_date` DATETIME NULL,
  `id_post` INT NOT NULL,
  `parent_id` INT NULL,
  `id_user` INT NOT NULL,
  PRIMARY KEY (`id_comment`),
  UNIQUE INDEX `id_comment_UNIQUE` (`id_comment` ASC) VISIBLE,
  INDEX `post_idx` (`id_post` ASC) VISIBLE,
  INDEX `user_idx` (`id_user` ASC) VISIBLE,
  UNIQUE INDEX `parent_id_UNIQUE` (`parent_id` ASC) VISIBLE,
  CONSTRAINT `comment_post`
    FOREIGN KEY (`id_post`)
    REFERENCES `reddit_clone`.`post` (`id_post`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `comment_parent`
    FOREIGN KEY (`id_comment`)
    REFERENCES `reddit_clone`.`comment` (`parent_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `comment_user`
    FOREIGN KEY (`id_user`)
    REFERENCES `reddit_clone`.`user` (`id_user`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

USE `reddit_clone`;

DELIMITER $$

USE `reddit_clone`$$
DROP TRIGGER IF EXISTS `reddit_clone`.`post_BEFORE_INSERT` $$
USE `reddit_clone`$$
CREATE DEFINER = CURRENT_USER TRIGGER `reddit_clone`.`post_BEFORE_INSERT` BEFORE INSERT ON `post` FOR EACH ROW
BEGIN
	SET NEW.created_date = NOW();
END$$


USE `reddit_clone`$$
DROP TRIGGER IF EXISTS `reddit_clone`.`post_BEFORE_UPDATE` $$
USE `reddit_clone`$$
CREATE DEFINER = CURRENT_USER TRIGGER `reddit_clone`.`post_BEFORE_UPDATE` BEFORE UPDATE ON `post` FOR EACH ROW
BEGIN
	SET NEW.edited_date = NOW();
END$$


USE `reddit_clone`$$
DROP TRIGGER IF EXISTS `reddit_clone`.`comment_BEFORE_INSERT` $$
USE `reddit_clone`$$
CREATE DEFINER = CURRENT_USER TRIGGER `reddit_clone`.`comment_BEFORE_INSERT` BEFORE INSERT ON `comment` FOR EACH ROW
BEGIN
	SET NEW.created_date = NOW();
END$$


USE `reddit_clone`$$
DROP TRIGGER IF EXISTS `reddit_clone`.`comment_BEFORE_UPDATE` $$
USE `reddit_clone`$$
CREATE DEFINER = CURRENT_USER TRIGGER `reddit_clone`.`comment_BEFORE_UPDATE` BEFORE UPDATE ON `comment` FOR EACH ROW
BEGIN
	SET NEW.edited_date = NOW();
END$$


DELIMITER ;
SET SQL_MODE = '';
DROP USER IF EXISTS backend;
SET SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
CREATE USER 'backend';

GRANT CREATE, DROP ON reddit_clone.* TO 'backend';
GRANT CREATE, DROP, INSERT, SELECT, UPDATE, DELETE ON TABLE `reddit_clone`.`subreddit` TO 'backend';
GRANT CREATE, DROP, DELETE, INSERT, SELECT, UPDATE ON TABLE `reddit_clone`.`post` TO 'backend';
GRANT CREATE, DROP, DELETE, INSERT, SELECT, UPDATE ON TABLE `reddit_clone`.`user` TO 'backend';
GRANT CREATE, DROP, DELETE, INSERT, SELECT, UPDATE ON TABLE `reddit_clone`.`comment` TO 'backend';

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;