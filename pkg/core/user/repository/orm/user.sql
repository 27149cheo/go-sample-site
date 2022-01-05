CREATE TABLE IF NOT EXISTS `user`(
    `id` varchar(64) NOT NULL COMMENT '用户ID',
    `name` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
    `alias` varchar(100) NOT NULL DEFAULT '' COMMENT '用户别名',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 59 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户信息表' ROW_FORMAT = Compact;
