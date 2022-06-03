DROP TABLE IF EXISTS `follow`;
DROP TABLE IF EXISTS `favourite`;
DROP TABLE IF EXISTS `comment`;
DROP TABLE IF EXISTS `video`;
DROP TABLE IF EXISTS `user`;

CREATE TABLE `user`
(
    `id`             bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`           varchar(255)        NOT NULL DEFAULT '' COMMENT '用户昵称',
    `password`       varchar(255)        NOT NULL DEFAULT '' COMMENT '密码',
    `follow_count`   int                 NOT NULL COMMENT '关注数量',
    `follower_count` bigint(20)          NOT NULL COMMENT '粉丝数量',
    `create_time`    datetime default CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
    `update_time`    datetime default CURRENT_TIMESTAMP NOT NULL COMMENT '修改时间',
    `is_deleted`     tinyint  default 0  NOT NULL COMMENT '逻辑删除',
    PRIMARY KEY (`id`)
)ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

INSERT INTO `user`
VALUES (1, 'Jerry', '123456', 10, 5,  '2022-06-01 10:00:00', '2022-06-01 10:00:00', 0),
       (2, 'Tom', '123456', 10, 5, '2022-06-01 10:00:00', '2022-06-01 10:00:00', 0),
       (3, 'Tony', '123456', 10, 5, '2022-06-01 10:00:00', '2022-06-01 10:00:00', 0);


CREATE TABLE `video`
(
    `id`              bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `uid`             bigint(20) unsigned NOT NULL COMMENT '视频发布者id',
    `play_url`        text                NOT NULL COMMENT '视频文件路径',
    `cover_url`       text                NOT NULL COMMENT '视频封面路径',
    `comment_count`   bigint(20)          NOT NULL COMMENT '评论数量',
    `favourite_count` bigint(20)          NOT NULL COMMENT '点赞数量',
    `title`           varchar(255)        NOT NULL COMMENT '视频标题',
    `create_time`     datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '发布时间',
    `update_time`     datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '修改时间',
    `is_deleted`      tinyint  DEFAULT 0  NOT NULL COMMENT '逻辑删除',
    PRIMARY KEY (`id`),
    constraint video_user_uid_fk
        foreign key (uid) references user (id)
)ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4 COMMENT ='视频表';

INSERT INTO `video`
VALUES (1, 1, 'bear.mp4', 'bear.jpg', 10, 4, 'bear',  '2022-06-01 10:00:00', '2022-06-01 10:00:00', 0),
       (2, 2, 'bear.mp4', 'bear.jpg', 2, 5, 'bear','2022-06-01 10:00:00', '2022-06-01 10:00:00', 0),
       (3, 3, 'bear.mp4', 'bear.jpg', 3, 6, 'bear','2022-06-01 10:00:00', '2022-06-01 10:00:00', 0);


CREATE TABLE `comment`
(
    `id`          bigint(20) unsigned NOT NULL COMMENT '主键id',
    `vid`         bigint(20) unsigned NOT NULL COMMENT '视频id',
    `uid`         bigint(20) unsigned NOT NULL COMMENT '用户id',
    `content`     text                NOT NULL COMMENT '评论内容',
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '发布时间',
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '修改时间',
    `is_deleted`   tinyint  DEFAULT 0  NOT NULL COMMENT '逻辑删除',
    PRIMARY KEY (`id`),
    constraint comment_user_uid_fk
        foreign key (`uid`) references user (`id`),
    constraint comment_video_vid_fk
        foreign key (`vid`) references video (`id`)
)ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4 COMMENT ='评论表';

INSERT INTO `comment`
VALUES (1, 2, 1, '这视频很有意思喔', '2022-06-01 10:00:00', '2022-06-01 10:00:00', 0);




CREATE TABLE `favourite`
(
    `id`            bigint(20) unsigned NOT NULL COMMENT '主键id',
    `uid`           bigint(20) unsigned NOT NULL COMMENT '点赞者id',
    `vid`           bigint(20) unsigned NOT NULL COMMENT '被点赞的视频id',
    `is_favourite`  tinyint  DEFAULT 0  NOT NULL COMMENT '是否点赞',
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '发布时间',
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '修改时间',
    constraint favourite_user_uid_fk
        foreign key (`uid`) references user (`id`),
    constraint favourite_video_vid_fk
        foreign key (`vid`) references video (`id`)
)ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4 COMMENT ='点赞表';

INSERT INTO `favourite`
VALUES (1, 1, 2, 1,'2022-06-01 10:00:00', '2022-06-01 10:00:00'),
       (2, 2, 2, 1,'2022-06-01 10:00:00', '2022-06-01 10:00:00');



CREATE TABLE `follow`
(
    `id`         bigint(20) unsigned NOT NULL COMMENT '关注id',
    `my_uid`     bigint(20) unsigned NOT NULL COMMENT '用户id',
    `his_uid`    bigint(20) unsigned NOT NULL COMMENT '用户查看的其他人的id',
    `is_follow`  tinyint DEFAULT 0 NOT NULL COMMENT '是否关注',
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '发布时间',
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '修改时间',
    constraint follow_my_uid_fk
        foreign key (my_uid) references user (`id`),
    constraint follow_his_uid_fk
            foreign key (his_uid) references user (`id`)
)ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4 COMMENT ='关注表';

INSERT INTO `follow`
VALUES (1, 1, 2, 1,'2022-06-01 10:00:00', '2022-06-01 10:00:00'),
       (2, 2, 2, 1,'2022-06-01 10:00:00', '2022-06-01 10:00:00');