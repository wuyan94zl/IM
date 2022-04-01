CREATE TABLE users
(
    id         INT auto_increment NOT NULL,
    user_name  varchar(100) NOT NULL DEFAULT '',
    nick_name  varchar(100) NOT NULL  DEFAULT '',
    password   varchar(100) NOT NULL  DEFAULT '',
    mobile     char(11)     NOT NULL  DEFAULT '',
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci
COMMENT='用户表';
