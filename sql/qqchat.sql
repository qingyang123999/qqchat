/*
 Navicat Premium Data Transfer

 Source Server         : 124.220.164.90
 Source Server Type    : MySQL
 Source Server Version : 80405
 Source Host           : 124.220.164.90:3306
 Source Schema         : qqchat

 Target Server Type    : MySQL
 Target Server Version : 80405
 File Encoding         : 65001

 Date: 28/08/2025 19:47:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_basic
-- ----------------------------
DROP TABLE IF EXISTS `user_basic`;
CREATE TABLE `user_basic`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `email` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `identity` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `client_ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `client_port` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `login_time` datetime NULL DEFAULT NULL,
  `heartbeat_time` datetime NULL DEFAULT NULL,
  `logout_time` datetime NULL DEFAULT NULL,
  `is_logout` int UNSIGNED NULL DEFAULT NULL,
  `device_info` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_basic
-- ----------------------------
INSERT INTO `user_basic` VALUES (32, 'HHHHF', 'm#$%UYIY', '18146495249', 'i.ffgytlydh@qq.com', '92', '131.125.153.142', '80', '2007-10-09 08:10:37', '1971-06-04 06:10:49', '1970-11-14 09:46:07', 1, 'incididunt ea cillum sint', '2025-08-28 18:57:49', '2025-08-28 18:57:49', NULL);
INSERT INTO `user_basic` VALUES (33, 'HHHHF', 'm#$%UYIY', '18146495249', 'i.ffgytl1ydh@qq.com', '92', '131.125.153.142', '800', '2007-10-09 08:10:37', '1971-06-04 06:10:49', '1970-11-14 09:46:07', 1, 'incididu111cillum sint', '2025-08-28 18:58:03', '2025-08-28 18:58:03', NULL);
INSERT INTO `user_basic` VALUES (34, 'JJJJJ', 'in @#$%^cons', '18626291539', 'i.ntndqkkueo@qq.com', '95', '234.74.184.102', '8245254', '1976-07-24 09:42:16', '2007-11-14 21:31:53', '1981-08-29 00:46:36', 1, 'consequat sed labore', '2025-08-28 18:58:38', '2025-08-28 18:58:38', NULL);
INSERT INTO `user_basic` VALUES (35, 'JJJJJ', 'in @#$%^cons', '18626291539', 'i.ntndqkkueo@qq.com', '95', '234.74.184.102', '8245254', '1976-07-24 09:42:16', '2007-11-14 21:31:53', '1981-08-29 00:46:36', 1, 'consequat sed labore', '2025-08-28 18:58:40', '2025-08-28 18:58:40', '2025-08-28 19:22:51');
INSERT INTO `user_basic` VALUES (36, 'HGFGH', 'incid!@#$%illum', '18159972164', 's.yoeb@qq.com', '53', '225.202.88.88', '7777', '2013-08-20 12:14:05', '2023-08-24 03:44:03', '1996-02-29 10:40:27', 0, 'Ut ad mollit', '2025-08-28 18:58:41', '2025-08-28 19:18:32', '2025-08-28 19:43:36');

SET FOREIGN_KEY_CHECKS = 1;
