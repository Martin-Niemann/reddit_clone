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
  `parent_id` INT NOT NULL,
  `id_user` INT NOT NULL,
  PRIMARY KEY (`id_comment`),
  UNIQUE INDEX `id_comment_UNIQUE` (`id_comment` ASC) VISIBLE,
  INDEX `post_idx` (`id_post` ASC) VISIBLE,
  INDEX `user_idx` (`id_user` ASC) VISIBLE,
  INDEX `comment_parent_idx` (`parent_id` ASC) VISIBLE,
  CONSTRAINT `comment_post`
    FOREIGN KEY (`id_post`)
    REFERENCES `reddit_clone`.`post` (`id_post`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `comment_user`
    FOREIGN KEY (`id_user`)
    REFERENCES `reddit_clone`.`user` (`id_user`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `comment_parent`
    FOREIGN KEY (`parent_id`)
    REFERENCES `reddit_clone`.`comment` (`id_comment`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
