-- -----------------------------------------------------
-- Schema reddit_clone
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `reddit_clone`;
USE `reddit_clone`;

-- -----------------------------------------------------
-- Table `reddit_clone`.`users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `reddit_clone`.`users`;

CREATE TABLE IF NOT EXISTS `reddit_clone`.`users` (
  `id_user` INT NOT NULL AUTO_INCREMENT,
  `user_name` VARCHAR(100) NOT NULL,
  `email` VARCHAR(320) NOT NULL,
  `hashed_password` VARCHAR(60) NOT NULL,
  PRIMARY KEY (`id_user`),
  UNIQUE INDEX `id_user_UNIQUE` (`id_user` ASC) VISIBLE,
  UNIQUE INDEX `user_name_UNIQUE` (`user_name` ASC) VISIBLE,
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `reddit_clone`.`subreddits`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `reddit_clone`.`subreddits`;

CREATE TABLE IF NOT EXISTS `reddit_clone`.`subreddits` (
  `id_subreddit` INT NOT NULL AUTO_INCREMENT,
  `url` VARCHAR(100) NOT NULL,
  `name` VARCHAR(100) NOT NULL,
  `description` VARCHAR(1000) NULL,
  `id_user` INT NULL,
  PRIMARY KEY (`id_subreddit`),
  UNIQUE INDEX `id_UNIQUE` (`id_subreddit` ASC) VISIBLE,
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE,
  UNIQUE INDEX `url_UNIQUE` (`url` ASC) VISIBLE,
  INDEX `users_idx` (`id_user` ASC) VISIBLE,
  FULLTEXT `subreddits_fulltext` (`url`,`name`,`description`),
  CONSTRAINT `subreddits_owner`
    FOREIGN KEY (`id_user`)
    REFERENCES `reddit_clone`.`users` (`id_user`)
    ON DELETE SET NULL
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `reddit_clone`.`posts`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `reddit_clone`.`posts`;

CREATE TABLE IF NOT EXISTS `reddit_clone`.`posts` (
  `id_post` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(300) NOT NULL,
  `link` VARCHAR(1000) NULL,
  `text` VARCHAR(10000) NULL,
  `id_subreddit` INT NOT NULL,
  `id_user` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_post`),
  UNIQUE INDEX `id_post_UNIQUE` (`id_post` ASC) VISIBLE,
  INDEX `users_idx` (`id_user` ASC) VISIBLE,
  FULLTEXT `posts_fulltext` (`title`,`link`,`text`),
  CONSTRAINT `posts_subreddits`
    FOREIGN KEY (`id_subreddit`)
    REFERENCES `reddit_clone`.`subreddits` (`id_subreddit`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `posts_users`
    FOREIGN KEY (`id_user`)
    REFERENCES `reddit_clone`.`users` (`id_user`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `reddit_clone`.`comments`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `reddit_clone`.`comments`;

CREATE TABLE IF NOT EXISTS `reddit_clone`.`comments` (
  `id_comment` INT NOT NULL AUTO_INCREMENT,
  `comment` VARCHAR(10000) NOT NULL,
  `is_toplevel` BOOLEAN NOT NULL,
  `id_post` INT NOT NULL,
  `parent_id` INT NULL,
  `id_user` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_comment`),
  UNIQUE INDEX `id_comment_UNIQUE` (`id_comment` ASC) VISIBLE,
  INDEX `posts_idx` (`id_post` ASC) VISIBLE,
  INDEX `users_idx` (`id_user` ASC) VISIBLE,
  INDEX `parent_idx` (`parent_id` ASC) VISIBLE,
  CONSTRAINT `comments_posts`
    FOREIGN KEY (`id_post`)
    REFERENCES `reddit_clone`.`posts` (`id_post`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `comments_users`
    FOREIGN KEY (`id_user`)
    REFERENCES `reddit_clone`.`users` (`id_user`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `comments_parent`
    FOREIGN KEY (`parent_id`)
    REFERENCES `reddit_clone`.`comments` (`id_comment`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `reddit_clone`.`scores`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `reddit_clone`.`scores`;

CREATE TABLE IF NOT EXISTS `reddit_clone`.`scores` (
  `id_score` INT NOT NULL AUTO_INCREMENT,
  `score` BOOL NOT NULL,
  `id_post` INT,
  `id_comment` INT,
  `id_user` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_score`),
  UNIQUE INDEX `id_score_UNIQUE` (`id_score` ASC) VISIBLE,
  INDEX `posts_idx` (`id_post` ASC) VISIBLE,
  INDEX `comments_idx` (`id_comment` ASC) VISIBLE,
  INDEX `users_idx` (`id_user` ASC) VISIBLE,
  CONSTRAINT `score_posts`
    FOREIGN KEY (`id_post`)
    REFERENCES `reddit_clone`.`posts` (`id_post`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `score_comment`
    FOREIGN KEY (`id_comment`)
    REFERENCES `reddit_clone`.`comments` (`id_comment`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `score_users`
    FOREIGN KEY (`id_user`)
    REFERENCES `reddit_clone`.`users` (`id_user`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

-- -----------------------------------------------------
-- User `backend`
-- -----------------------------------------------------
DROP USER IF EXISTS backend;

CREATE USER 'backend@127.0.0.1';
-- ALTER USER 'backend' IDENTIFIED BY 'password';

GRANT CREATE, DROP ON reddit_clone.* TO 'backend';
GRANT CREATE, DROP, INSERT, SELECT, UPDATE, DELETE ON TABLE `reddit_clone`.`subreddits` TO 'backend';
GRANT CREATE, DROP, DELETE, INSERT, SELECT, UPDATE ON TABLE `reddit_clone`.`posts` TO 'backend';
GRANT CREATE, DROP, DELETE, INSERT, SELECT, UPDATE ON TABLE `reddit_clone`.`users` TO 'backend';
GRANT CREATE, DROP, DELETE, INSERT, SELECT, UPDATE ON TABLE `reddit_clone`.`comments` TO 'backend';
GRANT CREATE, DROP, DELETE, INSERT, SELECT, UPDATE ON TABLE `reddit_clone`.`scores` TO 'backend';
