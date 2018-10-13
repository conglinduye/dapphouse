/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : tron_dapphouse

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 13/10/2018 18:43:38
*/

CREATE DATABASE `tron_dapphouse` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dapp
-- ----------------------------
DROP TABLE IF EXISTS `dapp`;
CREATE TABLE `dapp` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `account_id` int(10) NOT NULL COMMENT '账户id',
  `dapp_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'dapp名字',
  `developer` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '开发者',
  `catagory` tinyint(4) NOT NULL DEFAULT '0' COMMENT '分类 0:other, 1:games, 2:exchanges, 3:collectibles, 4:marketplaces, 5:gambling',
  `tagline` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'tagline',
  `version` varchar(24) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '版本',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '描述',
  `website_url` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'url',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0:提交未审核通过, 1:已审核通过, 2:审核拒绝, 3:下架',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `onsale_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上架时间',
  `offsale_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下架时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='dapp表';

-- ----------------------------
-- Table structure for dapp_contract
-- ----------------------------
DROP TABLE IF EXISTS `dapp_contract`;
CREATE TABLE `dapp_contract` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dapp_id` int(10) NOT NULL COMMENT 'dappid',
  `mainnet` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'mainnet',
  `testnet` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'testnet',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='dapp_contract表';

-- ----------------------------
-- Table structure for dapp_ext
-- ----------------------------
DROP TABLE IF EXISTS `dapp_ext`;
CREATE TABLE `dapp_ext` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `dapp_id` int(10) NOT NULL COMMENT 'dappid',
  `github` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'github',
  `twitter` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'twitter',
  `facebook` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'facebook',
  `telegram` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'telegram',
  `reddit` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'reddit',
  `medium` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'medium',
  `webchat` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'webchat',
  `weibo` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'weibo',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='dapp_ext表';

-- ----------------------------
-- Table structure for dapp_preview
-- ----------------------------
DROP TABLE IF EXISTS `dapp_preview`;
CREATE TABLE `dapp_preview` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `dapp_id` int(10) NOT NULL COMMENT 'dappid',
  `logo_image_url` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'logo_image_url',
  `screenshot_image_url` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'screenshot_image_url',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 0:删除, 1:正常',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='dapp_preview表';

-- ----------------------------
-- Table structure for demo
-- ----------------------------
DROP TABLE IF EXISTS `demo`;
CREATE TABLE `demo` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户',
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_account
-- ----------------------------
DROP TABLE IF EXISTS `user_account`;
CREATE TABLE `user_account` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `email` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `salt` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'salt',
  `pwd_crypt` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码加密',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 -1:注销, 0:注册未激活, 1:激活',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户账户表';

SET FOREIGN_KEY_CHECKS = 1;
