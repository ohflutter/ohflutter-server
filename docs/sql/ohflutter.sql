/*
 Navicat Premium Data Transfer

 Source Server         : 本地主机
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : ohflutter

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 27/05/2019 23:10:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ohflutter_articles
-- ----------------------------
DROP TABLE IF EXISTS `ohflutter_articles`;
CREATE TABLE `ohflutter_articles`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created` int(11) NULL DEFAULT NULL,
  `modified` int(11) NULL DEFAULT NULL,
  `deleted` int(11) NULL DEFAULT NULL,
  `user_id` int(11) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `ohflutter_articles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `ohflutter_users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for ohflutter_users
-- ----------------------------
DROP TABLE IF EXISTS `ohflutter_users`;
CREATE TABLE `ohflutter_users`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created` int(11) NULL DEFAULT NULL,
  `modified` int(11) NULL DEFAULT NULL,
  `deleted` int(11) NULL DEFAULT NULL,
  `mobile` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `age` bigint(20) NULL DEFAULT NULL,
  `birthday` timestamp(0) NULL DEFAULT NULL,
  `city` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
