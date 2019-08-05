## LearnOpDB
Learn how to operation gorm by go

## database
test

## Config
the common function : connect and close databse

## table
```
CREATE TABLE `users` (
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `age` int(11) DEFAULT '18',
  `birthday` timestamp NULL DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `member_number` varchar(255) DEFAULT NULL,
  `num` int(11) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_users_name` (`name`),
  UNIQUE KEY `member_number` (`member_number`),
  KEY `idx_users_deleted_at` (`deleted_at`),
  KEY `addr` (`address`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `animals` (
  `animal_id` varchar(255) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `master_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`animal_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `teachers` (
  `id` int(11) NOT NULL,
  `name` varchar(45) NOT NULL,
  `stu_id` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

## deploy
```
cd /{ProjectName}
go get github.com/kataras/rizla
go mod tiny
```

## run
go run main
