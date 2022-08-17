/*
 Navicat Premium Data Transfer

 Source Server         : LocalMysql
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : localhost:3306
 Source Schema         : blogserver

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 17/08/2022 22:07:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `post_id` bigint(20) NOT NULL,
  `author_id` bigint(20) NOT NULL,
  `parent_id` bigint(20) NULL DEFAULT NULL,
  `status` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_author_Id`(`author_id`) USING BTREE,
  INDEX `idx_comment_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, 'test', 0, 0, 0, 200, '2022-08-08 20:46:25.522', '2022-08-08 20:46:25.522', NULL);
INSERT INTO `comment` VALUES (2, 'test', 0, 0, 0, 200, '2022-08-08 20:47:06.468', '2022-08-08 20:47:06.468', NULL);

-- ----------------------------
-- Table structure for community
-- ----------------------------
DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `community_id` int(10) UNSIGNED NOT NULL,
  `community_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `introduction` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_community_id`(`community_id`) USING BTREE,
  UNIQUE INDEX `idx_community_name`(`community_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of community
-- ----------------------------
INSERT INTO `community` VALUES (1, 1, 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10');
INSERT INTO `community` VALUES (2, 2, 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00');

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `author_id` bigint(20) NULL DEFAULT NULL,
  `community_id` bigint(20) NULL DEFAULT NULL,
  `status` tinyint(4) NULL DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_author_id`(`author_id`) USING BTREE,
  INDEX `idx_community_id`(`community_id`) USING BTREE,
  INDEX `idx_post_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of post
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `gender` tinyint(4) NOT NULL DEFAULT 0,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 50 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (47, 'abcd', '$2a$10$Gc8g44IM/V/EDDn6kd5En.25l96dEatECv2.CCIvSmpmq/9Hz.xmi', '', 0, '2022-08-15 22:07:09.460', '2022-08-15 22:07:09.460', NULL, 138807352517197825);
INSERT INTO `user` VALUES (48, 'ddff', '$2a$10$F09P5EZpEADzsmf/djWODevj5hwhw/QChuHIhPOKdUVN88XaD5UgO', '', 0, '2022-08-16 16:42:55.323', '2022-08-16 16:42:55.323', NULL, 138919669049327617);
INSERT INTO `user` VALUES (49, '123', '$2a$10$imYj9LjGpakv28RdA6Xmx.KbbBwkrTLFpfUblO7K/qD6FmxLIdjG6', '', 0, '2022-08-16 17:43:22.623', '2022-08-16 17:43:22.623', NULL, 138925754648887297);

SET FOREIGN_KEY_CHECKS = 1;
