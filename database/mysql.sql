CREATE DATABASE IF NOT EXISTS gBlog CHARACTER SET utf8;
use gBlog;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) DEFAULT '' COMMENT '用户名',
  `password` varchar(255) DEFAULT '' COMMENT '密码',
  `email` varchar(255) DEFAULT '' COMMENT '邮箱',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `status` int DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户基本表';

DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` int DEFAULT '0' COMMENT '用户ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `category_id` int NOT NULL COMMENT '分类ID',
  `tag` varchar(255) NOT NULL DEFAULT '' COMMENT 'Tag',
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `pv` int DEFAULT '0' COMMENT 'pv',
  `status` int DEFAULT '1' COMMENT '1可用，2禁用，3删除',
  `review` int DEFAULT '0' COMMENT '评论',
  `recommend` int NOT NULL DEFAULT '0' COMMENT '是否顶置，0否；1是，默认否',
  `like` int NOT NULL DEFAULT '0' COMMENT '点赞数量',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) DEFAULT '',
  `pid` int DEFAULT '0' COMMENT '父ID',
  `sort` int DEFAULT '0' COMMENT '排序',
  `status` int DEFAULT '1' COMMENT '状态1正常，2删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `access_log`;
CREATE TABLE `access_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `ip` varchar(15) NOT NULL DEFAULT '' COMMENT 'ip',
  `city` varchar(50) NOT NULL DEFAULT '' COMMENT '城市',
  `uri` varchar(255) NOT NULL DEFAULT '' COMMENT '访问路由',
  `create` datetime DEFAULT NULL COMMENT '访问时间',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

# 20210409
# 增加列
ALTER TABLE `access_log` add COLUMN `country` varchar(50) NOT NULL DEFAULT '' COMMENT '国家' AFTER `ip`;
ALTER TABLE `access_log` add COLUMN `ISP` varchar(50) NOT NULL DEFAULT '' COMMENT '运营商' AFTER `city`;