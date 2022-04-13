CREATE TABLE `todos` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `is_done` tinyint(1) unsigned NOT NULL DEFAULT '0'
);
