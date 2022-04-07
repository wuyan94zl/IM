CREATE TABLE users
(
    id          INT auto_increment NOT NULL,
    user_name   varchar(100) NOT NULL DEFAULT '',
    nick_name   varchar(100) NOT NULL DEFAULT '',
    password    varchar(100) NOT NULL DEFAULT '',
    mobile      char(11)     NOT NULL DEFAULT '',
    create_time TIMESTAMP NULL,
    update_time TIMESTAMP NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';
-- goctl model mysql ddl -src="./user.sql" -dir="./user" -c