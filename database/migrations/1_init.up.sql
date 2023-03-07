/*
 Navicat MySQL Data Transfer

 Source Server         : root
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : herman

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 06/03/2023 20:10:52
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`
(
    `id`           int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
    `user`         varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '管理员用户名',
    `password`     char(60) CHARACTER SET utf8 COLLATE utf8_general_ci    NOT NULL COMMENT '管理员密码',
    `photo`        varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
    `name`         varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
    `card`         varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '身份证号码',
    `sex`          tinyint(4) NOT NULL DEFAULT 3 COMMENT '性别(1为女,2为男,3为保密)',
    `age`          tinyint(4) NOT NULL DEFAULT 0 COMMENT '年龄',
    `region`       varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '住址',
    `phone`        varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号码',
    `email`        varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
    `introduction` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '简介',
    `state`        tinyint(4) NOT NULL DEFAULT 2 COMMENT '状态(1已停用,2已启用)',
    `login_out_at` datetime NULL DEFAULT NULL COMMENT '上一次登录时间',
    `login_out_ip` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '上一次登录IP地址',
    `login_total`  int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录总数',
    `sort`         int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
    `created_at`   datetime                                               NOT NULL COMMENT '创建时间',
    `updated_at`   datetime                                               NOT NULL COMMENT '更新时间',
    `deleted_at`   datetime NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `idx_user`(`user`) USING BTREE COMMENT '管理员用户名索引'
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin`
VALUES (1, 'admin', '$2a$10$e3Jv5Fa7WU0e5c9QTYjEQ.G1E6ex30Q404DzmBwUsZoNBASvPsZty',
        'https://avatars.githubusercontent.com/u/61196516?v=4', '超级管理员', '650000197000000000', 3, 100, '中国北京',
        '15288888888', 'wetalk.vip@foxmail.com', '这是一个超级管理员', 2, NULL, NULL, 4, 1, '2023-01-15 23:08:51',
        '2023-03-04 21:48:21', NULL);

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`
(
    `id`         int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `admin_id`   int(11) UNSIGNED NULL DEFAULT NULL COMMENT '管理员ID',
    `role_key`   varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色KEY',
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` datetime NOT NULL COMMENT '更新时间',
    `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX        `idx_admin_id`(`admin_id`) USING BTREE COMMENT '管理员角色索引',
    INDEX        `idx_role_key`(`role_key`) USING BTREE COMMENT '角色索引',
    CONSTRAINT `管理员外键` FOREIGN KEY (`admin_id`) REFERENCES `admin` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT `角色外键` FOREIGN KEY (`role_key`) REFERENCES `roles` (`role`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员角色中间表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_role
-- ----------------------------

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`
(
    `id`           int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
    `name`         varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名称',
    `role`         varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色英文KEY',
    `state`        tinyint(4) NOT NULL DEFAULT 2 COMMENT '状态(1已停用,2已启用)',
    `sort`         int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
    `introduction` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '简介',
    `created_at`   datetime                                               NOT NULL COMMENT '创建时间',
    `updated_at`   datetime                                               NOT NULL COMMENT '更新时间',
    `deleted_at`   datetime NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `role_index`(`role`) USING BTREE COMMENT '角色名索引',
    INDEX          `id`(`id`, `role`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of roles
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
    `id`           int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `user`         varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
    `password`     char(60) CHARACTER SET utf8 COLLATE utf8_general_ci    NOT NULL COMMENT '用户密码',
    `photo`        varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户头像',
    `nickname`     varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '昵称',
    `name`         varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
    `card`         char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '身份证号码',
    `sex`          tinyint(4) NOT NULL DEFAULT 3 COMMENT '性别(1为女，2为男，3为保密)',
    `age`          tinyint(4) NOT NULL DEFAULT 0 COMMENT '年龄',
    `region`       varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地区',
    `phone`        varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号码',
    `email`        varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
    `introduction` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '简介',
    `state`        tinyint(4) NOT NULL DEFAULT 2 COMMENT '状态(1已停用,2已启用)',
    `login_out_ip` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '最后登录IP地址',
    `login_total`  int(11) NOT NULL DEFAULT 0 COMMENT '登录总数',
    `login_out_at` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
    `created_at`   datetime                                               NOT NULL COMMENT '创建时间',
    `updated_at`   datetime                                               NOT NULL COMMENT '更新时间',
    `deleted_at`   datetime NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `user`(`user`) USING BTREE COMMENT '用户索引'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------


-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`
(
    `id`    int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `ptype` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `v0`    varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `v1`    varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `v2`    varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `v3`    varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `v4`    varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `v5`    varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'Casbin表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------

-- ----------------------------
-- Table structure for dictionary
-- ----------------------------
CREATE TABLE `dictionary`
(
    `id`         int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '数据字典ID',
    `name`       varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
    `code`       varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '唯一KEY',
    `remark`     varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
    `state`      tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1表示禁用，2表示启用',
    `created_at` datetime                                               NOT NULL COMMENT '创建时间',
    `updated_at` datetime                                               NOT NULL COMMENT '更新时间',
    `deleted_at` datetime                                                DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_code` (`code`) USING BTREE COMMENT '唯一标识码索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='数据字典表';

-- ----------------------------
-- Records of data_dictionary
-- ----------------------------

-- ----------------------------
-- Table structure for dictionary_detail
-- ----------------------------
CREATE TABLE `dictionary_detail`
(
    `id`            int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '数据字典明细ID',
    `dictionary_id` int(11) unsigned NOT NULL COMMENT '数据字典ID',
    `name`          varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '明细名称',
    `code`          varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '唯一明细KEY',
    `value`         varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '明细值',
    `remark`        varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
    `sort`          int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
    `state`         tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1表示禁用，2表示启用',
    `created_at`    datetime                                               NOT NULL COMMENT '创建时间',
    `updated_at`    datetime                                               NOT NULL COMMENT '更新时间',
    `deleted_at`    datetime                                                DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY             `idx_dictionary_id` (`dictionary_id`,`code`) USING BTREE COMMENT '数据字典索引',
    CONSTRAINT `数据字典外键` FOREIGN KEY (`dictionary_id`) REFERENCES `dictionary` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='数据字典明细表';

-- ----------------------------
-- Records of dictionary_detail
-- ----------------------------

-- ----------------------------
-- Table structure for file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `file_chunks`;
CREATE TABLE `file_chunks`
(
    `id`           int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分片信息ID',
    `file_id`      int(11) UNSIGNED NOT NULL COMMENT '文件ID',
    `chunk_number` int(11) UNSIGNED NOT NULL COMMENT '分片编号',
    `chunk_size`   bigint(20) UNSIGNED NOT NULL COMMENT '分片大小',
    `chunk_path`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '分片路径',
    `hash`         varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '分片hash值',
    `state`        tinyint(4) NOT NULL DEFAULT 1 COMMENT '上传状态，1表示未上传，2表示已上传',
    `progress`     bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上传进度',
    `created_at`   datetime                                                NOT NULL COMMENT '创建时间',
    `updated_at`   datetime                                                NOT NULL COMMENT '更新时间',
    `deleted_at`   datetime NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX          `idx_file_chunks`(`file_id`, `chunk_number`, `hash`) USING BTREE COMMENT '用户索引',
    CONSTRAINT `文件外键` FOREIGN KEY (`file_id`) REFERENCES `files` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文件分片信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of file_chunks
-- ----------------------------

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files`
(
    `id`         int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '文件ID',
    `user_id`    int(11) UNSIGNED NOT NULL COMMENT '用户ID',
    `file_name`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '文件名',
    `file_type`  varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL COMMENT '文件类型',
    `file_path`  varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '文件路径',
    `hash`       varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '文件hash值',
    `file_size`  bigint(20) UNSIGNED NOT NULL COMMENT '文件大小',
    `created_at` datetime                                                NOT NULL COMMENT '创建时间',
    `updated_at` datetime                                                NOT NULL COMMENT '更新时间',
    `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX        `idx_files`(`user_id`, `hash`) USING BTREE COMMENT '用户索引'
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文件信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of files
-- ----------------------------

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`
(
    `id`         int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `pid`        int(11) UNSIGNED NULL DEFAULT NULL COMMENT '菜单父ID',
    `name`       varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL COMMENT '菜单名称',
    `path`       varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '路由PATH',
    `method`     varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci  NOT NULL COMMENT 'PATH的请求方法',
    `sort`       int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
    `created_at` datetime                                                NOT NULL COMMENT '创建时间',
    `updated_at` datetime                                                NOT NULL COMMENT '更新时间',
    `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of menus
-- ----------------------------

-- ----------------------------
-- Table structure for schema_migrations
-- ----------------------------
DROP TABLE IF EXISTS `schema_migrations`;
CREATE TABLE `schema_migrations`
(
    `version` bigint(20) NOT NULL,
    `dirty`   tinyint(1) NOT NULL,
    PRIMARY KEY (`version`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '数据库迁移表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of schema_migrations
-- ----------------------------

SET
FOREIGN_KEY_CHECKS = 1;
