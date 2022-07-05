-- Active: 1657024936523@@127.0.0.1@3306@mercado_fresco
DROP SCHEMA IF EXISTS mercado_fresco;
CREATE SCHEMA mercado_fresco;
USE mercado_fresco;

CREATE TABLE `sellers` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `cid` VARCHAR(255) NOT NULL UNIQUE,
  `company_name` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `telephone` VARCHAR(255) NOT NULL,
  `locality_id` INT NOT NULL
);

CREATE TABLE `products` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `description` varchar(255) NOT NULL,
  `expiration_rate` int NOT NULL,
  `freezing_rate` int NOT NULL,
  `height` decimal(19, 2) NOT NULL,
  `length` decimal(19, 2) NOT NULL,
  `net_weight` decimal(19, 2) NOT NULL,
  `product_code` varchar(255) NOT NULL UNIQUE,
  `recommended_freezing_temperature` decimal(19, 2) NOT NULL,
  `width` decimal(19, 2) NOT NULL,
  `product_type_id` int NOT NULL,
  `seller_id` int NOT NULL
);

CREATE TABLE `warehouse` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `address` varchar(255) NOT NULL,
  `telephone` varchar(255) NOT NULL,
  `warehouse_code` varchar(255) NOT NULL UNIQUE,
  `minimum_capacity` int NOT NULL,
  `minimum_temperature` DECIMAL(19,2) NOT NULL,
  `locality_id` INT NOT NULL
);

CREATE TABLE `sections` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `section_number` INT NOT NULL UNIQUE,
  `current_temperature` DECIMAL(19,2) NOT NULL,
  `minimum_temperature` DECIMAL(19,2) NOT NULL,
  `current_capacity` INT NOT NULL,
  `maximum_capacity` int NOT NULL,
  `warehouse_id` int NOT NULL,
  `product_type_id` int NOT NULL
);

CREATE TABLE `employees` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `card_number_id` VARCHAR(255) NOT NULL UNIQUE,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  `warehouse_id` int NOT NULL
);

CREATE TABLE `buyers` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `card_number_id` int NOT NULL UNIQUE,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL
);

CREATE TABLE `localities` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `locality_name` VARCHAR(255) NOT NULL,
    `province_id` INT NOT NULL
);

CREATE TABLE `provinces` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `province_name` VARCHAR(255) NOT NULL,
    `id_country_fk` INT NOT NULL
);

CREATE TABLE `countries` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `country_name` VARCHAR(255) NOT NULL
);

CREATE TABLE `users` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `password` VARCHAR(255) NOT NULL,
    `username` VARCHAR(255) NOT NULL
);

CREATE TABLE `users_rol` (
	`usuario_id` INT NOT NULL,
    `rol_id` INT NOT NULL,
    PRIMARY KEY (`usuario_id`, `rol_id`)
);

CREATE TABLE `rol` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `description` VARCHAR(255) NOT NULL,
    `rol_name` VARCHAR(255) NOT NULL
);

CREATE TABLE `purchase_orders` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `order_number` VARCHAR(255) NOT NULL,
    `order_date` DATETIME(6) NOT NULL,
    `tracking_code` VARCHAR(255) NOT NULL,
    `buyer_id` INT NOT NULL,
    `carrier_id` INT NOT NULL,
    `order_status_id` INT NOT NULL,
    `wareHouse_id` INT NOT NULL
);

CREATE TABLE `order_status` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `description` VARCHAR(255) NOT NULL
);

CREATE TABLE `carriers` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `cid` VARCHAR(255) NOT NULL UNIQUE,
    `company_name` VARCHAR(255) NOT NULL,
    `address` VARCHAR(255) NOT NULL,
    `telephone` VARCHAR(255) NOT NULL,
    `locality_id` INT NOT NULL
);

CREATE TABLE `inbound_orders` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `order_date` DATETIME(6),
    `order_number` VARCHAR(255) NOT NULL,
    `employe_id` INT NOT NULL,
    `product_batch_id` INT NOT NULL,
    `wareHouse_id` INT NOT NULL
);

CREATE TABLE `product_batches` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `batch_number` VARCHAR(255) NOT NULL UNIQUE,
    `current_quantity` INT,
    `current_temperature` DECIMAL(19,2),
    `due_date` DATETIME(6),
    `initial_quantity` INT NOT NULL,
    `manufacturing_date` DATETIME(6),
    `manufacturing_hour` DATETIME(6),
    `minimum_temperature` DECIMAL(19,2),
    `product_id` INT NOT NULL,
    `section_id` INT NOT NULL,
    FOREIGN KEY (`product_id`) REFERENCES `products`(`id`),
    FOREIGN KEY (`section_id`) REFERENCES `sections`(`id`)
);

CREATE TABLE `products_types` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `description` VARCHAR(255) NOT NULL
);

CREATE TABLE `product_records` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `last_update_date` DATETIME(6),
    `purchase_price` DECIMAL(19,2),
    `sale_price` DECIMAL(19,2),
    `product_id` INT NOT NULL,
    FOREIGN KEY (`product_id`) REFERENCES `products`(`id`)
);

CREATE TABLE `order_details` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `clean_liness_status` VARCHAR(255) NOT NULL,
    `quantity` INT NOT NULL,
    `temperature` DECIMAL(19,2),
    `product_record_id` INT NOT NULL,
    `purchase_order_id` INT NOT NULL,
    FOREIGN KEY (`product_record_id`) REFERENCES `product_records`(`id`),
    FOREIGN KEY (`purchase_order_id`) REFERENCES `purchase_orders`(`id`)
);

ALTER TABLE `products` ADD FOREIGN KEY (`seller_id`) REFERENCES `sellers` (`id`);

ALTER TABLE `products` ADD FOREIGN KEY (`product_type_id`) REFERENCES `products_types` (`id`);

ALTER TABLE `sections` ADD FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`);

ALTER TABLE `sections` ADD FOREIGN KEY (`product_type_id`) REFERENCES `products_types` (`id`);

ALTER TABLE `employees` ADD FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`);

ALTER TABLE `warehouse` ADD FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`);

ALTER TABLE `sellers` ADD FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`);

ALTER TABLE `localities` ADD FOREIGN KEY (`province_id`) REFERENCES `provinces` (`id`);

ALTER TABLE `provinces` ADD FOREIGN KEY (`id_country_fk`) REFERENCES `countries` (`id`);

ALTER TABLE `users_rol` ADD FOREIGN KEY (`usuario_id`) REFERENCES `users` (`id`);

ALTER TABLE `users_rol` ADD FOREIGN KEY (`rol_id`) REFERENCES `rol` (`id`);

ALTER TABLE `purchase_orders` ADD FOREIGN KEY (`buyer_id`) REFERENCES `buyers` (`id`);

ALTER TABLE `purchase_orders` ADD FOREIGN KEY (`carrier_id`) REFERENCES `carriers` (`id`);

ALTER TABLE `purchase_orders` ADD FOREIGN KEY (`order_status_id`) REFERENCES `order_status` (`id`);

ALTER TABLE `purchase_orders` ADD FOREIGN KEY (`wareHouse_id`) REFERENCES `warehouse` (`id`);

ALTER TABLE `carriers` ADD FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`);

ALTER TABLE `inbound_orders` ADD FOREIGN KEY (`employe_id`) REFERENCES `employees` (`id`);

ALTER TABLE `inbound_orders` ADD FOREIGN KEY (`product_batch_id`) REFERENCES `product_batches` (`id`);

ALTER TABLE `inbound_orders` ADD FOREIGN KEY (`wareHouse_id`) REFERENCES `warehouse` (`id`);

INSERT INTO `mercado_fresco`.`countries` (`id`, `country_name`) VALUES (1, 'a');

INSERT INTO `mercado_fresco`.`provinces` (`id`, `province_name`, `id_country_fk`) VALUES (1, 'a', 1);
INSERT INTO `mercado_fresco`.`localities` (`id`, `locality_name`, `province_id`) VALUES (1, 'a', 1);
INSERT INTO `mercado_fresco`.`sellers` (`id`, `cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (1, '1', 'a', 'a', 'a', 1);
INSERT INTO `mercado_fresco`.`products` (`id`, `description`, `expiration_rate`, `freezing_rate`, `height`, `length`, `net_weight`, `product_code`, `recommended_freezing_temperature`, `width`, `product_type_id`, `seller_id`) VALUES (1, 'a', 1, 1, 1, 1, 1, 'a', 1, 1, 1, 1);
INSERT INTO `mercado_fresco`.`products_types` (`id`, `description`) VALUES (1, 'a');
INSERT INTO `mercado_fresco`.`warehouse` (`id`, `address`, `telephone`, `warehouse_code`, `minimun_capacity`, `minimun_temperature`, `locality_id`) VALUES (1, 'a', 'a', 'a', 1, 1, 1);
INSERT INTO `mercado_fresco`.`sections` (`id`, `section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (1, 1, 1, 1, 1, 1, 1, 1);
INSERT INTO `mercado_fresco`.`product_batches` (`id`, `batch_number`, `current_quantity`, `current_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `minimum_temperature`, `product_id`, `section_id`) VALUES (1, '1', 1, 1, '2004-05-23T14:25:10', 1, '2004-05-23T14:25:10', '2004-05-23T14:25:10', 1, 1, 1);
INSERT INTO `mercado_fresco`.`product_records` (`id`, `last_update_date`, `purchase_price`, `sale_price`, `product_id`) VALUES (1, '2004-05-23T14:25:10', 1, 1, 1);
INSERT INTO `mercado_fresco`.`employees` (`id`, `card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES (1, 'a', 'a', 'a', 1);
INSERT INTO `mercado_fresco`.`carriers` (`id`, `cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (1, 'a', 'a', 'a', 'a', 1);
INSERT INTO `mercado_fresco`.`order_status` (`id`, `description`) VALUES (1, 'a');
INSERT INTO `mercado_fresco`.`buyers` (`id`, `card_number_id`, `first_name`, `last_name`) VALUES (1, 1, 'a', 'a');
INSERT INTO `mercado_fresco`.`purchase_orders` (`id`, `order_number`, `order_date`, `tracking_code`, `buyer_id`, `carrier_id`, `order_status_id`, `wareHouse_id`) VALUES (1, '1', '2004-05-23T14:25:10', 'a', 1, 1, 1, 1);
INSERT INTO `mercado_fresco`.`inbound_orders` (`id`, `order_date`, `order_number`, `employe_id`, `product_batch_id`, `wareHouse_id`) VALUES (1, '2004-05-23T14:25:10', '1', 1, 1, 1);
INSERT INTO `mercado_fresco`.`order_details` (`id`, `clean_liness_status`, `quantity`, `temperature`, `product_record_id`, `purchase_order_id`) VALUES (1, 'a', 1, 1, 1, 1);
INSERT INTO `mercado_fresco`.`users` (`id`, `password`, `username`) VALUES (1, 'a', 'a');
INSERT INTO `mercado_fresco`.`rol` (`id`, `description`, `rol_name`) VALUES (1, 'a', 'a');
INSERT INTO `mercado_fresco`.`users_rol` (`usuario_id`, `rol_id`) VALUES (1, 1);