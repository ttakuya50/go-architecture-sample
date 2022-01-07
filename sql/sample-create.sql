CREATE DATABASE IF NOT EXISTS mydb CHARACTER SET utf8mb4;

use mydb;

CREATE TABLE IF NOT EXISTS `mydb`.`user` (
                                             `id` BIGINT NOT NULL,
                                             `name` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`))
    ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `mydb`.`list` (
    `id` BIGINT NOT NULL,
    `user_id` BIGINT NOT NULL COMMENT 'ユーザーID',
    `title` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `user_id_idx` (`user_id` ASC) VISIBLE,
    CONSTRAINT `fk_user_id_list`
    FOREIGN KEY (`user_id`)
    REFERENCES `mydb`.`user` (`id`)
        ON DELETE NO ACTION
        ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `mydb`.`task` (
                                             `id` BIGINT NOT NULL,
                                             `list_id` BIGINT NOT NULL,
                                             `user_id` BIGINT NOT NULL,
                                             `title` VARCHAR(255) NOT NULL,
    `memo` VARCHAR(255) NULL,
    `is_done` TINYINT(1) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `list_id_idx` (`list_id` ASC) VISIBLE,
    INDEX `user_id_idx` (`user_id` ASC) VISIBLE,
    CONSTRAINT `fk_list_id_task`
    FOREIGN KEY (`list_id`)
    REFERENCES `mydb`.`list` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
    CONSTRAINT `fk_user_id_task`
    FOREIGN KEY (`user_id`)
    REFERENCES `mydb`.`user` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
    ENGINE = InnoDB;
