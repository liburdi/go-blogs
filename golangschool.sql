/*
 Navicat Premium Data Transfer

 Source Server         : 39.108.104.253
 Source Server Type    : MySQL
 Source Server Version : 50611
 Source Host           : 39.108.104.253:3306
 Source Schema         : 5iweb

 Target Server Type    : MySQL
 Target Server Version : 50611
 File Encoding         : 65001

 Date: 25/01/2021 19:01:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for php41_auth
-- ----------------------------
DROP TABLE IF EXISTS `php41_auth`;
CREATE TABLE `php41_auth` (
  `auth_id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `auth_name` varchar(20) NOT NULL COMMENT '名称',
  `auth_pid` smallint(6) unsigned NOT NULL COMMENT '父id',
  `auth_c` varchar(32) NOT NULL DEFAULT '' COMMENT '模块',
  `auth_a` varchar(32) NOT NULL DEFAULT '' COMMENT '操作方法',
  `auth_path` varchar(32) NOT NULL DEFAULT '' COMMENT '全路径',
  `auth_level` tinyint(4) NOT NULL DEFAULT '0' COMMENT '基别',
  PRIMARY KEY (`auth_id`)
) ENGINE=InnoDB AUTO_INCREMENT=122 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for php41_duanzi
-- ----------------------------
DROP TABLE IF EXISTS `php41_duanzi`;
CREATE TABLE `php41_duanzi` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0',
  `comment` text,
  `oo` int(4) unsigned NOT NULL DEFAULT '0',
  `xx` int(4) unsigned NOT NULL DEFAULT '0',
  `author` varchar(25) NOT NULL DEFAULT '',
  `is_show` enum('展示','不展示') DEFAULT '展示' COMMENT '是否显示',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=85 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for php41_goods
-- ----------------------------
DROP TABLE IF EXISTS `php41_goods`;
CREATE TABLE `php41_goods` (
  `goods_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `goods_name` varchar(100) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '名称',
  `goods_number` smallint(6) NOT NULL DEFAULT '1' COMMENT '浏览数',
  `goods_small_logo` char(100) CHARACTER SET utf8 NOT NULL DEFAULT './Home/Public/img/gopher_default.jpg' COMMENT '封面图',
  `is_hot` enum('热销','不热销') CHARACTER SET utf8 NOT NULL DEFAULT '不热销' COMMENT '热销与否',
  `is_new` enum('新品','不新品') CHARACTER SET utf8 NOT NULL DEFAULT '不新品' COMMENT '新品与否',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `upd_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `is_del` enum('删除','不删除') CHARACTER SET utf8 NOT NULL DEFAULT '不删除' COMMENT '是否删除',
  `author` varchar(32) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '作者',
  `keywords` varchar(300) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '关键字',
  `description` varchar(500) DEFAULT NULL,
  `cate` char(32) CHARACTER SET utf8 NOT NULL DEFAULT 'Go语言' COMMENT '分类 默认Go语言',
  `cat_id` int(2) unsigned NOT NULL DEFAULT '1' COMMENT '展示形式 默认1',
  `origin_author` varchar(32) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '原始作者',
  `origin_url` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '原始url',
  `user_id` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '用户id',
  `city` varchar(50) CHARACTER SET utf8 NOT NULL DEFAULT '未知' COMMENT '城市',
  PRIMARY KEY (`goods_id`),
  KEY `add_time` (`add_time`),
  KEY `uid_gnumber` (`user_id`,`goods_id`),
  KEY `uid` (`user_id`),
  KEY `gid_cate_gnum` (`goods_id`,`goods_number`,`cate`),
  KEY `cate_gid` (`cate`,`goods_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4124 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for php41_goods_introduce
-- ----------------------------
DROP TABLE IF EXISTS `php41_goods_introduce`;
CREATE TABLE `php41_goods_introduce` (
  `goods_id` int(10) DEFAULT NULL,
  `goods_introduce` mediumtext,
  KEY `goods_id` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for php41_goods_pics
-- ----------------------------
DROP TABLE IF EXISTS `php41_goods_pics`;
CREATE TABLE `php41_goods_pics` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `goods_id` int(10) unsigned NOT NULL COMMENT '商品',
  `pics_big` char(100) NOT NULL DEFAULT '' COMMENT '相册原图',
  `pics_small` char(100) NOT NULL DEFAULT '' COMMENT '相册缩略图',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='商品相册表';

-- ----------------------------
-- Table structure for php41_manager
-- ----------------------------
DROP TABLE IF EXISTS `php41_manager`;
CREATE TABLE `php41_manager` (
  `mg_id` int(11) NOT NULL AUTO_INCREMENT,
  `mg_name` varchar(32) NOT NULL,
  `mg_pwd` varchar(32) NOT NULL,
  `mg_time` int(10) unsigned NOT NULL COMMENT '时间',
  `mg_role_id` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  PRIMARY KEY (`mg_id`),
  UNIQUE KEY `mg_name` (`mg_name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for php41_ooxx
-- ----------------------------
DROP TABLE IF EXISTS `php41_ooxx`;
CREATE TABLE `php41_ooxx` (
  `id` mediumint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `target_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父id',
  `target_type` tinyint(4) DEFAULT '1' COMMENT '1:note',
  `msg` mediumtext,
  `like_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '赞',
  `unlike_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '踩',
  `from_user` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '来自谁',
  `to_user` int(11) NOT NULL DEFAULT '0' COMMENT '发送给谁',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级id，子评论才有',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=342 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for php41_phonecode
-- ----------------------------
DROP TABLE IF EXISTS `php41_phone_code`;
CREATE TABLE `php41_phone_code` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `phone_num` varchar(11) NOT NULL DEFAULT '',
  `code` int(4) unsigned NOT NULL DEFAULT '0',
  `result` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=207 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for php41_pic
-- ----------------------------
DROP TABLE IF EXISTS `php41_pic`;
CREATE TABLE `php41_pic` (
  `id` mediumint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `author` varchar(225) NOT NULL DEFAULT 'Dee' COMMENT '用户名',
  `url` longtext COMMENT 'url',
  `oo` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '赞',
  `xx` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '踩',
  `add_time` int(10) unsigned NOT NULL DEFAULT '1510000000' COMMENT '添加时间',
  `type` enum('jpg','gif') DEFAULT 'jpg' COMMENT '图片类型',
  `small_pic` text,
  `is_show` enum('不展示','展示') DEFAULT '不展示',
  `tag` char(10) NOT NULL DEFAULT '无聊图',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=163 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for php41_project
-- ----------------------------
DROP TABLE IF EXISTS `php41_project`;
CREATE TABLE `php41_project` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `is_del` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
  `creator_id` int(10) NOT NULL DEFAULT '0' COMMENT '作者',
  `desc` varchar(300) NOT NULL DEFAULT '' COMMENT '描述',
  `origin_url` varchar(255) NOT NULL DEFAULT '' COMMENT '原始url',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for php41_reply
-- ----------------------------
DROP TABLE IF EXISTS `php41_reply`;
CREATE TABLE `php41_reply` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `ooxx_id` int(10) unsigned NOT NULL DEFAULT '0',
  `uid` int(10) unsigned NOT NULL DEFAULT '0',
  `comment` text,
  `com_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for php41_role
-- ----------------------------
DROP TABLE IF EXISTS `php41_role`;
CREATE TABLE `php41_role` (
  `role_id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `role_name` varchar(20) NOT NULL COMMENT '角色名称',
  `role_auth_ids` varchar(128) NOT NULL DEFAULT '' COMMENT '权限ids,1,2,5',
  `role_auth_ac` text COMMENT '模块-操作',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for php41_tabs
-- ----------------------------
DROP TABLE IF EXISTS `php41_tabs`;
CREATE TABLE `php41_tabs` (
  `id` mediumint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(2) unsigned NOT NULL DEFAULT '0',
  `cate` varchar(32) NOT NULL COMMENT '用户名',
  `lv` tinyint(1) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for php41_type
-- ----------------------------
DROP TABLE IF EXISTS `php41_type`;
CREATE TABLE `php41_type` (
  `type_id` smallint(5) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `type_name` varchar(32) NOT NULL COMMENT '类型名称',
  PRIMARY KEY (`type_id`)
) ENGINE=MyISAM AUTO_INCREMENT=63 DEFAULT CHARSET=utf8 COMMENT='商品类型表';

-- ----------------------------
-- Table structure for php41_url
-- ----------------------------
DROP TABLE IF EXISTS `php41_url`;
CREATE TABLE `php41_url` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `url` varchar(100) NOT NULL DEFAULT '' COMMENT 'url',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '标题',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '简介',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `add_time` int(11) NOT NULL DEFAULT '0' COMMENT '时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for php41_users
-- ----------------------------
DROP TABLE IF EXISTS `php41_users`;
CREATE TABLE `php41_users` (
  `user_id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '昵称',
  `tel` varchar(20) NOT NULL DEFAULT '',
  `password` varchar(100) NOT NULL,
  `tag_list` text,
  `head_img` varchar(255) NOT NULL DEFAULT './Common/Uploads/2018-08-07/5b690555b0a58.png',
  `position` varchar(25) NOT NULL DEFAULT '程序猿',
  `company` varchar(25) NOT NULL DEFAULT '',
  `self_intro` varchar(255) NOT NULL DEFAULT '爱编程，爱Golang!',
  `homepage` varchar(255) NOT NULL DEFAULT '',
  `t_head_img` varchar(255) NOT NULL DEFAULT './Common/Uploads/2018-08-07/t_5b690555b0a58.png',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `unique_phone` (`tel`)
) ENGINE=MyISAM AUTO_INCREMENT=69 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
