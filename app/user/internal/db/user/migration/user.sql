/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1-3307
 Source Server Type    : MySQL
 Source Server Version : 50741
 Source Host           : localhost:3307
 Source Schema         : bs

 Target Server Type    : MySQL
 Target Server Version : 50741
 File Encoding         : 65001

 Date: 13/03/2023 15:33:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

SET FOREIGN_KEY_CHECKS = 1;
