/*
 Navicat Premium Data Transfer

 Source Server         : 124.220.164.90
 Source Server Type    : MySQL
 Source Server Version : 80405
 Source Host           : 124.220.164.90:3306
 Source Schema         : ginchat

 Target Server Type    : MySQL
 Target Server Version : 80405
 File Encoding         : 65001

 Date: 25/09/2025 17:13:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for communities
-- ----------------------------
DROP TABLE IF EXISTS `communities`;
CREATE TABLE `communities`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `owner_id` bigint UNSIGNED NULL DEFAULT NULL,
  `img` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_communities_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of communities
-- ----------------------------

-- ----------------------------
-- Table structure for contact
-- ----------------------------
DROP TABLE IF EXISTS `contact`;
CREATE TABLE `contact`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `owner_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '谁的信息',
  `target_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '关联谁',
  `type` bigint NULL DEFAULT NULL COMMENT '对应的类型 ',
  `desc` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_contact_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 185 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '人员关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of contact
-- ----------------------------

-- ----------------------------
-- Table structure for group_basic
-- ----------------------------
DROP TABLE IF EXISTS `group_basic`;
CREATE TABLE `group_basic`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '群名称',
  `owner_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '群主id',
  `icon` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '图片',
  `type` bigint NULL DEFAULT NULL COMMENT '类型',
  `desc` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_group_basic_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '群' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_basic
-- ----------------------------

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `form_id` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '发送者',
  `target_id` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '接收者',
  `type` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '发送类型  ',
  `media` bigint NULL DEFAULT NULL COMMENT '消息类型 文字  图片  音频',
  `content` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '消息内容',
  `pic` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '图片',
  `url` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '地址',
  `desc` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '描述',
  `amount` bigint NULL DEFAULT NULL COMMENT '其它数字相关的',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_message_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '消息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------

-- ----------------------------
-- Table structure for user_basic
-- ----------------------------
DROP TABLE IF EXISTS `user_basic`;
CREATE TABLE `user_basic`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `pass_word` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `phone` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `email` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `identity` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `client_ip` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `client_port` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `login_time` datetime(3) NULL DEFAULT NULL,
  `heartbeat_time` datetime(3) NULL DEFAULT NULL,
  `login_out_time` datetime(3) NULL DEFAULT NULL,
  `is_logout` tinyint(1) NULL DEFAULT NULL,
  `device_info` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `salt` longtext CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_basic_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_basic
-- ----------------------------
INSERT INTO `user_basic` VALUES (26, '2025-09-25 16:53:38.897', '2025-09-25 16:53:41.716', NULL, '18174918548', '11df8ef6ec146c33b4dcf2189abe5be5', '', '', '9702414BCDC5CAC3AFE8B21449B62869', '', '', '2025-09-25 16:53:38.871', '2025-09-25 16:53:38.871', '2025-09-25 16:53:38.871', 0, '', '1758006517', '');

SET FOREIGN_KEY_CHECKS = 1;
