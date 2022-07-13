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

CREATE TABLE `warehouses` (
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
  `minimum_capacity` INT NOT NULL,
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
  `card_number_id` VARCHAR(255) NOT NULL UNIQUE,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL
);

CREATE TABLE `localities` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `locality_name` VARCHAR(255) NOT NULL UNIQUE,
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
    `warehouse_id` INT NOT NULL
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
    `order_date` DATETIME(6) NOT NULL,
    `order_number` VARCHAR(255) NOT NULL UNIQUE,
    `employee_id` INT NOT NULL,
    `product_batch_id` INT NOT NULL,
    `warehouse_id` INT NOT NULL
);

CREATE TABLE `product_batches` (
	`id` INT AUTO_INCREMENT PRIMARY KEY,
    `batch_number` INT NOT NULL UNIQUE,
    `current_quantity` INT,
    `current_temperature` DECIMAL(19,2),
    `due_date` DATETIME(6),
    `initial_quantity` INT NOT NULL,
    `manufacturing_date` DATETIME(6),
    `manufacturing_hour` INT,
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
    `last_update_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `purchase_price` DECIMAL(19,2) NOT NULL,
    `sale_price` DECIMAL(19,2) NOT NULL,
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

ALTER TABLE `sections` ADD FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`);

ALTER TABLE `sections` ADD FOREIGN KEY (`product_type_id`) REFERENCES `products_types` (`id`);

ALTER TABLE `employees` ADD FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`);

ALTER TABLE `warehouses` ADD FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`);

ALTER TABLE `sellers` ADD FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`);

ALTER TABLE `localities` ADD FOREIGN KEY (`province_id`) REFERENCES `provinces` (`id`);

ALTER TABLE `provinces` ADD FOREIGN KEY (`id_country_fk`) REFERENCES `countries` (`id`);

ALTER TABLE `users_rol` ADD FOREIGN KEY (`usuario_id`) REFERENCES `users` (`id`);

ALTER TABLE `users_rol` ADD FOREIGN KEY (`rol_id`) REFERENCES `rol` (`id`);

ALTER TABLE `purchase_orders` ADD FOREIGN KEY (`buyer_id`) REFERENCES `buyers` (`id`);

ALTER TABLE `purchase_orders` ADD FOREIGN KEY (`carrier_id`) REFERENCES `carriers` (`id`);

ALTER TABLE `purchase_orders` ADD FOREIGN KEY (`order_status_id`) REFERENCES `order_status` (`id`);

ALTER TABLE `purchase_orders` ADD FOREIGN KEY (`wareHouse_id`) REFERENCES `warehouses` (`id`);

ALTER TABLE `carriers` ADD FOREIGN KEY (`locality_id`) REFERENCES `localities` (`id`);

ALTER TABLE `inbound_orders` ADD FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`);

ALTER TABLE `inbound_orders` ADD FOREIGN KEY (`product_batch_id`) REFERENCES `product_batches` (`id`);

ALTER TABLE `inbound_orders` ADD FOREIGN KEY (`warehouse_id`) REFERENCES `warehouses` (`id`);
