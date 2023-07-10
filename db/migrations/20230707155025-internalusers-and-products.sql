
-- +migrate Up
CREATE TABLE IF NOT EXISTS `internal_users`
(
	`id`	INT AUTO_INCREMENT PRIMARY KEY,
	`name`	VARCHAR(255) NOT NULL,
	`role_type`	SMALLINT NOT NULL COMMENT "1: normal; 2: manager; 3: admin",
	`created_at`	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at`	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at`	TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS `products` 
(
	`id`	INT AUTO_INCREMENT PRIMARY KEY,
	`name`	VARCHAR(255) NOT NULL,
	`price`	DOUBLE DEFAULT 0,
	`description`	VARCHAR(1000) NOT NULL,
	`created_by`	INT NOT NULL,
	`updated_by`	INT NOT NULL,

	`created_at`	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at`	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_at`	TIMESTAMP NULL,

	CONSTRAINT `fk_products_created_by` FOREIGN KEY (`created_by`) REFERENCES `internal_users`(`id`),
	CONSTRAINT `fk_products_updated_by` FOREIGN KEY (`updated_by`) REFERENCES `internal_users`(`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS internal_users;
