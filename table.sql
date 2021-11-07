USE db_posts; -- change db_posts dengan nama database yang akan digunakan
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
     `id` bigint NOT NULL AUTO_INCREMENT,
     `title` varchar(200),
     `content` longtext,
     `category` varchar(200),
     `status` varchar(100),
     `created_at` datetime(3) DEFAULT CURRENT_TIMESTAMP ,
     `updated_at` datetime(3) DEFAULT CURRENT_TIMESTAMP ,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;