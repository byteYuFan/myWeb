# 数据库介绍

## 1. 权限控制

为了保证数据库的安全性，现在分配给`pogf`用户`xaut`这个数据库的所有权限。

```mysql
CREATE USER 'pogf'@'%' IDENTIFIED BY '123456';
GRANT ALL PRIVILEGES ON xaut.* TO 'pogf'@'%';
FLUSH PRIVILEGES;
```

之后在程序中用该用户访问数据库。

## 2. 数据库介绍

在此我创建了一个名为`xaut`的数据库，其字符集以及相关默认数据如下所示。

```mysql
 CREATE DATABASE `xaut` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */
```

## 3. 相关表介绍

### 3.1. 用户注册信息表

每当有一个新用户注册时，我们会将该用户的信息存入到这张表中去，目前暂定的表字段有，`id`,`username`,`email`,`password`,`create_at`,`flag`这六个字段，分别代表，`用户id`,`用户名`,`用户邮箱`,`用户密码`,`创建时间`，`该用户是否被删除`,考虑到后续会经常使用这张表的`username`,`email`,`flag`这三个字段，于是我们决定将给这三个字段设置索引，flag的类型为bool默认为true，表示用户在这张表中。

```mysql
CREATE TABLE `user_register_info` (
                                      `id` bigint NOT NULL AUTO_INCREMENT,
                                      `username` varchar(50) NOT NULL,
                                      `email` varchar(50) NOT NULL,
                                      `password` varchar(255) NOT NULL,
                                      `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                      `flag` tinyint(1) DEFAULT '1',
                                      PRIMARY KEY (`id`),
                                      UNIQUE KEY `username` (`username`),
                                      UNIQUE KEY `email` (`email`),
                                      KEY `idx_username` (`username`),
                                      KEY `idx_email` (`email`),
                                      KEY `idx_flag` (`flag`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```

