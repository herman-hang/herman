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

 Date: 03/03/2023 00:15:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
                          `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
                          `user` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '管理员用户名',
                          `password` char(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '管理员密码',
                          `photo` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
                          `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
                          `card` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '身份证号码',
                          `sex` tinyint(4) NOT NULL DEFAULT 3 COMMENT '性别(1为女,2为男,3为保密)',
                          `age` tinyint(4) NOT NULL DEFAULT 0 COMMENT '年龄',
                          `region` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '住址',
                          `phone` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号码',
                          `email` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
                          `introduction` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '简介',
                          `state` tinyint(4) NOT NULL DEFAULT 2 COMMENT '状态(1已停用,2已启用)',
                          `login_out_at` datetime NULL DEFAULT NULL COMMENT '上一次登录时间',
                          `login_out_ip` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '上一次登录IP地址',
                          `login_total` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录总数',
                          `sort` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
                          `created_at` datetime NOT NULL COMMENT '创建时间',
                          `updated_at` datetime NOT NULL COMMENT '更新时间',
                          `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `user_index`(`user`) USING BTREE COMMENT '管理员用户名索引'
) ENGINE = InnoDB AUTO_INCREMENT = 68 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, 'admin', '$2a$10$e3Jv5Fa7WU0e5c9QTYjEQ.G1E6ex30Q404DzmBwUsZoNBASvPsZty', 'http://njgpuwilx.qa/brxpexftc', '蔡艳', '650000197006212198', 2, 104, '西南', '13425523344', 'j.mjifbm@mmt.be', '素新山织感报民院实连术已出元。', 2, NULL, NULL, 30, 26, '2023-01-15 23:08:51', '2023-03-02 22:15:27', NULL);
INSERT INTO `admin` VALUES (2, 'ehtzwczhre', '$2a$10$Svq3BUjDxLi0KsiGSEcoIORY1PwVUCayHOK6RQPdZ.ll.d8N.A7Mm', 'http://yostpo.ee/pwogm', '易秀英', '310000198410028495', 2, 65, '东北', '18189593589', 'q.mvvxjexkl@tkgizrofhx.tn', '太定较也低对年三拉县收个影离任。', 1, '1970-01-01 20:01:22', '', 0, 11, '2023-02-16 20:01:53', '2023-02-17 21:46:24', '2023-02-21 20:16:06');
INSERT INTO `admin` VALUES (3, 'ttwfjyizvwo', '$2a$10$6dCOpzKzFcNosPfQB49qO.9nAnVoZHLJKPbyt5Ki5zprS4QhlPmqe', 'http://gynb.jp/xtnhev', '蔡杰', '530000197208233571', 2, 111, '华南', '13890646351', 'f.tiybdqlk@xqk.vn', '间白装始求流受究华人元改建也具济听又。', 2, '1970-01-01 20:01:22', '', 0, 17, '2023-02-16 20:02:37', '2023-02-17 21:47:29', '2023-02-21 20:16:38');
INSERT INTO `admin` VALUES (4, 'uitmqlet', '$2a$10$OZ1PnIiUgfpnLvZuJ3/7guiEbV1BBe8dsbwSR3dheoj15KHoFKVJy', 'http://gdffig.lu/bfvixjoe', '黄强', '430000199707246380', 2, 6, '西北', '18663190178', 'n.jgdjdmrbi@tqfccw.pm', '自率月第叫例效好设节段称明展圆式。', 2, '1970-01-01 20:01:22', '', 0, 0, '2023-02-16 20:04:29', '2023-02-16 20:04:29', '2023-02-21 20:16:38');
INSERT INTO `admin` VALUES (5, 'hzkntv', '$2a$10$xSdVvQEybz1Jhgk/Lv43qehDFC5qrQpJu.xdTa4kg9UF7OQQvBWje', 'http://gndqnt.az/efoc', '廖娜', '41000020170605544X', 1, 70, '东北', '18642829187', 'z.mfloefkh@lsuhbfh.do', '王近去切始议今导算志去长存内铁压流。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 20:06:22', '2023-02-16 20:06:22', '2023-02-21 20:16:49');
INSERT INTO `admin` VALUES (6, 'lylvmw', '$2a$10$fozXR61oorUun9f/2u9L1eInZHpx5HfQ4dWucqbvSRrSqszJ7eWSW', 'http://djiocee.af/csofpjsm', '陈艳', '610000200304059782', 1, 99, '华南', '18120605772', 'l.ismsfkpc@cohgwltwed.ug', '率片段好机质电面战形持加应及。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 20:06:22', '2023-02-16 20:06:22', '2023-02-21 20:27:10');
INSERT INTO `admin` VALUES (7, 'sdqcmjil', '$2a$10$3TRmE6/z/LGVkjsAd/owSOKGGFAGI2xgAmp28FMVUUVZCnu7lrUK.', 'http://roiby.mg/wgigtwt', '于涛', '620000197806223797', 1, 54, '华中', '18164162417', 'e.mrunqtk@smxnnm.gl', '每年步任万验有记界上活为济队必采图拉。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 20:12:37', '2023-02-16 20:12:37', NULL);
INSERT INTO `admin` VALUES (8, 'atbqgxjki', '$2a$10$kvO2zoWLXlPtbDqAYm8n0exJiwH97wjhJQguEKmiBdEivFmt38g6K', 'http://ybk.cr/nislxbdih', '李艳', '63000019941221352X', 1, 77, '华东', '18100733698', 'i.kzzu@dvuqxlyqws.中国互联.公司', '片大设布用习压类不对满严高做。', 2, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 20:27:24', '2023-02-16 20:27:24', NULL);
INSERT INTO `admin` VALUES (9, 'jjrxtw', '$2a$10$h3XSgNDuS/Z/01TspVGfVuN/zOiiZ7k2LD6m1HnoS1cL6BsDTuYy2', 'http://hboaqyb.sl/yquyrr', '龚秀兰', '520000199403193576', 2, 80, '西南', '18658335333', 'i.cqqvxwvf@cbsqehabn.bd', '感酸红化管但火红好越和海新史。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 20:27:24', '2023-02-16 20:27:24', NULL);
INSERT INTO `admin` VALUES (10, 'hjncsnpg', '$2a$10$pDazLwNKDDwAgHR/gCN.B.3lQjCMPAzyrZpcsyjxbEA5AwfeXakBm', 'http://ustvl.sk/uwdhvn', '任超', '120000198904201328', 2, 88, '东北', '13839218035', 'x.ixv@umqeatdlf.rw', '走长建正事形民边铁体人会用了程但导。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 20:28:06', '2023-02-16 20:28:06', NULL);
INSERT INTO `admin` VALUES (13, 'igonku', '$2a$10$7zZgZ676EoVqi5B5PCwKF.VAdwUQs/SfDCqzwVCX0wXDiZ04Y11Um', 'http://mgijdy.cl/vadqdyh', '傅平', '15000019870617079X', 1, 34, '华东', '19824757142', 'u.mzbjjhq@sothtsifrw.aq', '公月况素参低平组飞共拉把最知压义。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 21:53:31', '2023-02-16 21:53:31', NULL);
INSERT INTO `admin` VALUES (31, 'xlenlywixow', '$2a$10$O/hXZCstiO2AnKd8NwM2nu5azOVqnCJwUVlEbkGMTTTuoErbe6b4i', 'http://eefuqq.nt/nvnetqym', '石艳', '310000200905298518', 2, 18, '西北', '18157376338', 'c.gyrg@nkcnimh.dk', '矿必导置众以就族型长因设义。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 22:08:38', '2023-02-16 22:08:38', NULL);
INSERT INTO `admin` VALUES (32, 'nsvjuhjdt', '$2a$10$orcNi2yHQXpvStPkYowikOAOvF/A8OE.5m46sEqtEzpdit/Z0WXjm', 'http://foonvcwx.hr/ookowiwmj', '阎娜', '410000201711036155', 2, 19, '华东', '18611577821', 'z.giiqc@ceykjrb.gt', '商格发温许效此验联你效把所。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin` VALUES (33, 'btoxvnfm', '$2a$10$eVY/rHW0eTmHUABpirbPdeEeCChkhQ3ATXM/Al/0DoKDW2esDvYSe', 'http://covfpje.sd/fvpy', '戴秀兰', '620000197102175248', 1, 64, '华东', '18653768738', 'b.yvjvrcpay@vfqcwbwlm.travel', '话积处合通电生米件规名它做改清感其。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin` VALUES (34, 'olxutuvzgclt', '$2a$10$sdb3kiNCFINeuDOuyAUXbu1j1WoKAI4eA/IIk3as9wih44G2CfeQi', 'http://txsfgqxm.kg/bdxkcurxt', '方平', '340000198303195736', 1, 110, '华东', '13617803669', 't.uueklfwjhs@noblhtqnd.lv', '际直片社感位更那做育族信压且。', 2, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin` VALUES (35, 'srdiprfehfycto', '$2a$10$jku0TPdA6lSw1vtpMDelIeEwcC5DcoS0pwlQOGcE.tUyYzM1JeEYq', 'http://mhcw.ug/bxayom', '姚强', '530000201512126083', 1, 22, '西北', '19806316639', 'q.fhbsxiebj@gtkeviheyi.gf', '强边就人商越通上水总得两面。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin` VALUES (36, 'vrqknyuysx', '$2a$10$sS4jiX8o1VmpICw27rssheFJ7ggde3qnjnda2fuSVJNniMgUONVIK', 'http://ojdueyot.gw/gdwpmayhc', '曹强', '110000200211070714', 1, 35, '华北', '18661320216', 'e.fmapqq@sup.us', '数阶个我构斗共现种非习么改数化受。', 2, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin` VALUES (37, 'tddecwbmwj', '$2a$10$szHc37fuABisemNbTAZgteWQdnZCGz8qnpbY0b8fTfMLYlt3eB06S', 'http://ipccxwnp.ve/vutn', '罗秀兰', '420000199309012456', 1, 18, '华南', '18122717386', 't.lwtpaioty@ktohu.sm', '有增命养于他新近史物九此说放。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin` VALUES (38, 'utmmpsm', '$2a$10$K57/VHwOmzDmN3rLsfQ6NubARXNI6jWLXiBMHuTVrBRZxkaUed.PC', 'http://tuuijlbbt.中国/jlj', '孔芳', '630000200807093150', 2, 23, '华中', '18158758414', 'g.pmckzgdx@zxgucogp.ph', '实队三人经命些了六律处其容个。', 1, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin` VALUES (39, 'aougcdibjqy', '$2a$10$5W0vgFLKqWKON9VEsHLL4egYXQ.wU6OMDtuvYyEtLtj2TwnIW0Kya', 'http://vtiy.sd/nccnl', '崔杰', '410000197904288691', 1, 98, '华南', '18156135832', 'h.ojhgxdasq@edoanci.vg', '以只民易选单除界标维常许新几需好局难。', 2, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin` VALUES (40, 'tcobautsog', '$2a$10$vZHp/FKdpf/WFYD8KaBEGemTdPdGpXdx9mqeXGLov9BYltG1AKr9e', 'http://dwhqq.is/ciexvvd', '赖桂英', '450000199103142357', 1, 45, '华北', '18616756227', 'j.kmhvslvomk@jmpnzrfd.coop', '族地住即面算县除统小命反。', 2, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin` VALUES (43, 'yftdnkihhduw', '$2a$10$BcqWdXA2SfLCcTP9Eh1HseyFJXBPipEwB0cfTgDjakT4xsrF6q5Ye', 'http://tudbtnziq.sm/qfybewmw', '罗静', '350000197010096635', 1, 82, '华中', '18166062772', 'j.nscspolnv@tdefh.mp', '受观阶众程国自处行温细权商外提感。', 2, '1970-01-01 00:00:00', '', 0, 0, '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin` VALUES (44, 'twqatjyftin', '$2a$10$v48NlDJcOx7pB8UFvtiriuhZ1Da//CKUf.TNf7RNQJt64aDUN65HW', 'http://ever.td/nvmp', '崔杰', '520000200010275576', 2, 2, '华南', '19899515746', 'n.hdtzwyjimf@mvriqgjjcf.mn', '清体想元人据三表收类眼边位重重。', 2, '1970-01-01 00:00:00', '', 0, 81, '2023-02-16 22:32:13', '2023-02-16 22:32:13', NULL);
INSERT INTO `admin` VALUES (45, 'dllroumwawevab', '$2a$10$pmi63heqeglG42V1we7O/.xYBIkQMObzBQHGY8Z14uZoc50OHzsSK', 'http://yuujkecimd.bg/glg', '杜秀兰', '130000200901133616', 1, 76, '华北', '18658380217', 'q.ulhlokd@gokhmumwf.cf', '红今立号花所会如高所路众关复委。', 1, '1970-01-01 00:00:00', '', 0, 81, '2023-02-16 22:32:13', '2023-02-16 22:32:13', NULL);
INSERT INTO `admin` VALUES (47, 'avhkgtqfbbk', '$2a$10$gKXe32J.R6knPn6YzexHH.cwUfFuYCVR4AZ0q9hs0xAzxVZktFz1W', 'http://hfkud.tk/kqbgwjhja', '姚勇', '420000197206125539', 2, 39, '西南', '18175758263', 'e.qsvsvo@otmr.ai', '东济事业厂已图来历应应打者清。', 1, '1970-01-01 00:00:00', '', 0, 75, '2023-02-16 23:40:43', '2023-02-16 23:40:43', NULL);
INSERT INTO `admin` VALUES (48, 'cinvdk', '$2a$10$Yvb4YncFpdSdeLquFbSKg.E1xDO.aLzaSK7plkzgmQRNEyl1wEXj6', 'http://yoesvfh.de/njwf', '魏桂英', '500000197611271708', 1, 91, '华北', '18127102824', 'r.gijllbin@nwgtqucqrk.gt', '相较提实她影周重速适信要。', 2, '1970-01-01 00:00:00', '', 0, 38, '2023-02-16 23:48:02', '2023-02-16 23:48:02', NULL);
INSERT INTO `admin` VALUES (49, 'tfpeiwbmbgr', '$2a$10$z.6heVEbbs/ruzcEHlKtl.ik5rFGD0DjsfeRhWt9.bstmflgvzlwO', 'http://tepmws.uy/arrkjtlp', '唐军', '110000200207178706', 2, 27, '华北', '13860853575', 'l.cayirpnd@vxdqff.中国互联.公司', '油山表连八给商特号非济公毛温。', 2, '1970-01-01 00:00:00', '', 0, 113, '2023-02-16 23:52:05', '2023-02-16 23:52:05', NULL);
INSERT INTO `admin` VALUES (50, 'bckveghzbk', '$2a$10$rQ0KFJ3KfaLvrVJzk7wT5.6A7Wp7p4G3qY3HzYnPyMo4gMmzeMZL.', 'http://xaw.ls/glvvxuvfg', '李勇', '50000019741031787X', 2, 50, '华中', '18153022484', 'f.zgvmup@tdrxtxpu.gi', '报命共米着回原用通只由做加如制般。', 1, '1970-01-01 00:00:00', '', 0, 118, '2023-02-16 23:52:33', '2023-02-16 23:52:33', NULL);
INSERT INTO `admin` VALUES (51, 'yjykwfy', '$2a$10$jBv6RSxKX88dYtnqYt8el.qnp8lfObLjwNutHeTwBS26RWPNrsdpi', 'http://lvrfbqt.org.cn/zhdnvfj', '阎超', '360000198508053356', 1, 95, '西南', '18150246528', 'f.ewv@iwtlbs.bs', '马记表九下事严治求所本技内手志。', 2, '1970-01-01 00:00:00', '', 0, 74, '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin` VALUES (52, 'fvwdpmkkwctpwvt', '$2a$10$wu/STikXNtX621gAbEQ0LemvfqnTaTAYJgbf9LxHCduXN/vtNJJ1y', 'http://hswrq.uy/nqroj', '李明', '710000199601268810', 2, 23, '西南', '18671691621', 'u.ibjfhcy@bwb.中国互联.网络', '又此出生集边有手好五运没声义。', 1, '1970-01-01 00:00:00', '', 0, 44, '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin` VALUES (53, 'ebzumrx', '$2a$10$dDuMfASuPc3Wm1CL0Rt7luNl/84A4dyWdfSqC8bzQq1h/IpH9QEr6', 'http://wpvth.bd/rhgv', '侯静', '650000199104196362', 2, 91, '华中', '18643941146', 'o.jucxlufr@dwyy.qa', '用近式强专按律转又利严花历对经。', 2, '1970-01-01 00:00:00', '', 0, 99, '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin` VALUES (54, 'fmeqjnqf', '$2a$10$9tAIcb1BWnlhyxWpbIgzcOTGcA6jdV8Cua6pipB465AYtUwxOUyQy', 'http://avjao.edu/zfsdjqequ', '韩秀英', '520000200605093472', 1, 88, '东北', '18195875825', 'q.jimfi@reade.kh', '事特几眼响书新教展角眼场个音始器。', 1, '1970-01-01 00:00:00', '', 0, 99, '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin` VALUES (55, 'ynmmohmnb', '$2a$10$.vDjUTxIgTogd0OfrD4y.uFz.xfA.aVE.F3PSkYaYga.j23iffJ2q', 'http://pbcjzijk.se/jps', '贺军', '430000202005312915', 1, 37, '西南', '13678472507', 'm.wsyei@vgxeuhtp.eh', '院要求才速中石话观再内义制产物联马。', 1, '1970-01-01 00:00:00', '', 0, 111, '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin` VALUES (56, 'kgdyrzmzyvavex', '$2a$10$XPL26lUKP5AQHPgVHv91ceihqazpxlv5KRY8OKbayydFWoLRMzhzq', 'http://mcjsfi.hm/qdza', '阎芳', '530000201011203060', 1, 86, '华东', '18642570213', 'x.fgosyq@xditjp.so', '铁同查调增元习很亲积每书。', 2, '1970-01-01 00:00:00', '', 0, 75, '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin` VALUES (62, 'walpracm', 'juulqqut', 'http://pwjm.bw/nhkowgd', '吴秀兰', '810000200905082539', 2, 30, '西北', '18168715772', 'f.gszj@fsouvvm.vu', '识今光农才向局许装被文年们多。', 1, '1970-01-01 00:00:00', '', 0, 59, '2023-02-17 21:43:22', '2023-02-17 21:43:22', NULL);
INSERT INTO `admin` VALUES (63, 'cktmrr', 'pgrmoipy', 'http://bngkpywls.中国/qfsrojv', '万艳', '440000197908249185', 1, 91, '华南', '19860082778', 'u.ugtymcp@kiuo.la', '第主所热计新外意公又按真区军。', 1, '1970-01-01 00:00:00', '', 0, 73, '2023-02-17 21:44:11', '2023-02-17 21:44:11', NULL);
INSERT INTO `admin` VALUES (64, 'ytjwmvpvq', 'schtnug', '', '龚伟', '530000200407038343', 2, 17, '东北', '18105011266', 'x.hlzf@sewxksobh.中国', '利权土单步己习集方决近育文。', 2, '1970-01-01 00:00:00', '', 0, 33, '2023-02-17 21:44:11', '2023-02-17 21:44:11', NULL);
INSERT INTO `admin` VALUES (65, 'jkiieg', 'eskucklpuddn', 'http://bukkeih.mc/vxnloeq', '', '530000197702154836', 2, 70, '西南', '18153884441', 'h.xveu@pxznhhud.tg', '果战上物作要阶打它教结当其。', 1, '1970-01-01 00:00:00', '', 0, 75, '2023-02-17 21:44:11', '2023-02-17 21:44:11', NULL);
INSERT INTO `admin` VALUES (66, 'nvrbjddhngcv', 'jyxkuhwpw', 'http://bgpedn.hu/msuny', '谢涛', '530000198005232223', 2, 5, '华北', '18124551346', 'u.tpu@vvntlizv.sn', '亲该整一同便间型张科器片切走安。', 1, '1970-01-01 00:00:00', '', 0, 104, '2023-02-17 21:45:39', '2023-02-17 21:45:39', NULL);
INSERT INTO `admin` VALUES (67, 'afiqpjjy', 'ibnmibeypo', 'http://dkmlsrj.bb/sooifwpsy', '郭艳', '120000202108056157', 1, 63, '华中', '18111626534', 'h.yudgrh@bvdhhf.dk', '增自越局计记步记定斗将元华。', 2, '1970-01-01 00:00:00', '', 0, 41, '2023-02-17 21:45:39', '2023-02-17 21:45:39', NULL);

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
                               `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                               `admin_id` int(11) UNSIGNED NULL DEFAULT NULL COMMENT '管理员ID',
                               `role_key` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色KEY',
                               `created_at` datetime NOT NULL COMMENT '创建时间',
                               `updated_at` datetime NOT NULL COMMENT '更新时间',
                               `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                               PRIMARY KEY (`id`) USING BTREE,
                               INDEX `管理员索引`(`admin_id`) USING BTREE COMMENT '管理员角色索引',
                               INDEX `角色索引`(`role_key`) USING BTREE COMMENT '角色索引',
                               CONSTRAINT `管理员外键` FOREIGN KEY (`admin_id`) REFERENCES `admin` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
                               CONSTRAINT `角色外键` FOREIGN KEY (`role_key`) REFERENCES `roles` (`role`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 63 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员角色中间表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES (1, 31, 'n3toP', '2023-02-16 22:08:38', '2023-02-16 22:08:38', NULL);
INSERT INTO `admin_role` VALUES (2, 31, 'ZtG', '2023-02-16 22:08:38', '2023-02-16 22:08:38', NULL);
INSERT INTO `admin_role` VALUES (3, 32, 'n3toP', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (4, 32, 'ZtG', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (5, 33, 'n3toP', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (6, 33, 'ZtG', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (7, 34, 'n3toP', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (8, 34, 'ZtG', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (9, 35, 'n3toP', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (10, 35, 'ZtG', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (11, 36, 'n3toP', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (12, 36, 'ZtG', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (13, 37, 'n3toP', '2023-02-16 22:27:23', '2023-02-16 22:27:23', '2023-02-26 23:21:57');
INSERT INTO `admin_role` VALUES (14, 37, 'ZtG', '2023-02-16 22:27:23', '2023-02-16 22:27:23', '2023-02-26 23:21:57');
INSERT INTO `admin_role` VALUES (15, 38, 'n3toP', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (16, 38, 'ZtG', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (17, 39, 'n3toP', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (18, 39, 'ZtG', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (19, 40, 'n3toP', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (20, 40, 'ZtG', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (21, 43, 'n3toP', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (22, 43, 'ZtG', '2023-02-16 22:27:23', '2023-02-16 22:27:23', NULL);
INSERT INTO `admin_role` VALUES (23, 44, 'n3toP', '2023-02-16 22:32:13', '2023-02-16 22:32:13', NULL);
INSERT INTO `admin_role` VALUES (24, 44, 'ZtG', '2023-02-16 22:32:13', '2023-02-16 22:32:13', NULL);
INSERT INTO `admin_role` VALUES (25, 45, 'n3toP', '2023-02-16 22:32:13', '2023-02-16 22:32:13', NULL);
INSERT INTO `admin_role` VALUES (26, 45, 'ZtG', '2023-02-16 22:32:13', '2023-02-16 22:32:13', NULL);
INSERT INTO `admin_role` VALUES (27, 47, 'n3toP', '2023-02-16 23:40:43', '2023-02-16 23:40:43', NULL);
INSERT INTO `admin_role` VALUES (28, 47, 'ZtG', '2023-02-16 23:40:43', '2023-02-16 23:40:43', NULL);
INSERT INTO `admin_role` VALUES (29, 48, 'n3toP', '2023-02-16 23:48:02', '2023-02-16 23:48:02', NULL);
INSERT INTO `admin_role` VALUES (30, 48, 'ZtG', '2023-02-16 23:48:02', '2023-02-16 23:48:02', NULL);
INSERT INTO `admin_role` VALUES (31, 49, 'n3toP', '2023-02-16 23:52:05', '2023-02-16 23:52:05', NULL);
INSERT INTO `admin_role` VALUES (32, 49, 'ZtG', '2023-02-16 23:52:05', '2023-02-16 23:52:05', NULL);
INSERT INTO `admin_role` VALUES (33, 50, 'n3toP', '2023-02-16 23:52:33', '2023-02-16 23:52:33', NULL);
INSERT INTO `admin_role` VALUES (34, 50, 'ZtG', '2023-02-16 23:52:33', '2023-02-16 23:52:33', NULL);
INSERT INTO `admin_role` VALUES (35, 51, 'n3toP', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (36, 51, 'ZtG', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (37, 52, 'n3toP', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (38, 52, 'ZtG', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (39, 53, 'n3toP', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (40, 53, 'ZtG', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (41, 54, 'n3toP', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (42, 54, 'ZtG', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (43, 55, 'n3toP', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (44, 55, 'ZtG', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (45, 56, 'n3toP', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (46, 56, 'ZtG', '2023-02-17 19:59:25', '2023-02-17 19:59:25', NULL);
INSERT INTO `admin_role` VALUES (47, 62, 'n3toP', '2023-02-17 21:43:22', '2023-02-17 21:43:22', NULL);
INSERT INTO `admin_role` VALUES (48, 62, 'ZtG', '2023-02-17 21:43:22', '2023-02-17 21:43:22', NULL);
INSERT INTO `admin_role` VALUES (49, 63, 'n3toP', '2023-02-17 21:44:11', '2023-02-17 21:44:11', NULL);
INSERT INTO `admin_role` VALUES (50, 63, 'ZtG', '2023-02-17 21:44:11', '2023-02-17 21:44:11', NULL);
INSERT INTO `admin_role` VALUES (51, 64, 'n3toP', '2023-02-17 21:44:11', '2023-02-17 21:44:11', NULL);
INSERT INTO `admin_role` VALUES (52, 64, 'ZtG', '2023-02-17 21:44:11', '2023-02-17 21:44:11', NULL);
INSERT INTO `admin_role` VALUES (53, 65, 'n3toP', '2023-02-17 21:44:11', '2023-02-17 21:44:11', NULL);
INSERT INTO `admin_role` VALUES (54, 65, 'ZtG', '2023-02-17 21:44:11', '2023-02-17 21:44:11', NULL);
INSERT INTO `admin_role` VALUES (55, 66, 'n3toP', '2023-02-17 21:45:39', '2023-02-17 21:45:39', NULL);
INSERT INTO `admin_role` VALUES (56, 66, 'ZtG', '2023-02-17 21:45:39', '2023-02-17 21:45:39', NULL);
INSERT INTO `admin_role` VALUES (57, 1, 'n3toP', '2023-02-17 21:45:39', '2023-02-17 21:45:39', '2023-02-17 21:46:12');
INSERT INTO `admin_role` VALUES (58, 1, 'n3toP', '2023-02-17 21:45:39', '2023-02-17 21:45:39', '2023-02-17 21:46:37');
INSERT INTO `admin_role` VALUES (59, 2, 'n3toP', '2023-02-17 21:45:39', '2023-02-17 21:45:39', NULL);
INSERT INTO `admin_role` VALUES (60, 1, 'n3toP', '2023-02-17 21:45:39', '2023-02-17 21:45:39', NULL);
INSERT INTO `admin_role` VALUES (61, 67, 'n3toP', '2023-02-17 21:45:39', '2023-02-17 21:45:39', NULL);
INSERT INTO `admin_role` VALUES (62, 67, 'ZtG', '2023-02-17 21:45:39', '2023-02-17 21:45:39', NULL);
INSERT INTO `admin_role` VALUES (63, 3, 'n3toP', '2023-02-17 21:45:39', '2023-02-17 21:45:39', NULL);

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
                                `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
                                `ptype` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                                PRIMARY KEY (`id`) USING BTREE,
                                UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 68 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (5, 'g', 'aassdd123', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (36, 'g', 'dpj1b[', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (37, 'g', 'dpj1b[', 'ZtG', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (52, 'g', 'JfDSf', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (53, 'g', 'JfDSf', 'ZtG', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (24, 'g', 'm3K', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (25, 'g', 'm3K', 'ZtG', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (16, 'g', 'mno6VL', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (17, 'g', 'mno6VL', 'ZtG', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (48, 'g', 'nTs4UA', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (49, 'g', 'nTs4UA', 'ZtG', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (40, 'g', 'tFhLT', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (41, 'g', 'tFhLT', 'ZtG', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (8, 'g', 'uK@cB2', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (9, 'g', 'uK@cB2', 'ZtG', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (28, 'g', 'UV@l', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (29, 'g', 'UV@l', 'ZtG', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (20, 'g', 'VBnDU@M', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (21, 'g', 'VBnDU@M', 'ZtG', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (44, 'g', 'YU%x', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (45, 'g', 'YU%x', 'ZtG', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (32, 'g', '[ptM', 'n3toP', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (33, 'g', '[ptM', 'ZtG', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (65, 'g', '[W19', 'VBnDU@M', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (6, 'p', '9mkM', '/api/a/b', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (7, 'p', '9mkM', '/api/a/c', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (38, 'p', 'dpj1b[', '/api/v1/a/c', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (39, 'p', 'dpj1b[', '/api/v1/a/y', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (54, 'p', 'JfDSf', '/api/v1/a/c', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (55, 'p', 'JfDSf', '/api/v1/a/y', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (26, 'p', 'm3K', '/api/v1/a/c', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (27, 'p', 'm3K', '/api/v1/a/y', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (18, 'p', 'mno6VL', '/api/v1/a/b', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (19, 'p', 'mno6VL', '/api/v1/a/c', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (4, 'p', 'n3toP', '/api/a/c', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (3, 'p', 'n3toP', '/api/v1/a/b', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (50, 'p', 'nTs4UA', '/api/v1/a/c', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (51, 'p', 'nTs4UA', '/api/v1/a/y', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (42, 'p', 'tFhLT', '/api/v1/a/c', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (43, 'p', 'tFhLT', '/api/v1/a/y', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (10, 'p', 'uK@cB2', '/api/v1/a/b', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (11, 'p', 'uK@cB2', '/api/v1/a/c', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (30, 'p', 'UV@l', '/api/v1/a/c', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (31, 'p', 'UV@l', '/api/v1/a/y', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (22, 'p', 'VBnDU@M', '/api/v1/a/c', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (23, 'p', 'VBnDU@M', '/api/v1/a/y', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (46, 'p', 'YU%x', '/api/v1/a/c', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (47, 'p', 'YU%x', '/api/v1/a/y', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (34, 'p', '[ptM', '/api/v1/a/c', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (35, 'p', '[ptM', '/api/v1/a/y', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (66, 'p', '[W19', '/api/v1/a/b', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (67, 'p', '[W19', '/api/v1/a/d', 'POST', '', '', '');

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`  (
                          `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                          `pid` int(11) UNSIGNED NULL DEFAULT NULL COMMENT '菜单父ID',
                          `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '菜单名称',
                          `path` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '路由PATH',
                          `method` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'PATH的请求方法',
                          `sort` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
                          `created_at` datetime NOT NULL COMMENT '创建时间',
                          `updated_at` datetime NOT NULL COMMENT '更新时间',
                          `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                          PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES (1, NULL, '测试路由', '/api/v1/a/b', 'GET', 0, '2023-02-28 22:42:37', '2023-02-28 22:42:41', NULL);

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles`  (
                          `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
                          `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色名称',
                          `role` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色英文KEY',
                          `state` tinyint(4) NOT NULL DEFAULT 2 COMMENT '状态(1已停用,2已启用)',
                          `sort` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
                          `introduction` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '简介',
                          `created_at` datetime NOT NULL COMMENT '创建时间',
                          `updated_at` datetime NOT NULL COMMENT '更新时间',
                          `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `role_index`(`role`) USING BTREE COMMENT '角色名索引',
                          INDEX `id`(`id`, `role`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 48 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO `roles` VALUES (5, '真人上划要已', 'K%xB9m', 2, 0, '干度了圆话中色队式九机加以机。', '2023-02-13 19:51:59', '2023-02-13 19:51:59', '2023-02-27 20:52:04');
INSERT INTO `roles` VALUES (6, '达比少心', 'n3toP', 2, 0, '算特总格了使头品王展办美导。', '2023-02-13 19:51:59', '2023-02-13 19:51:59', NULL);
INSERT INTO `roles` VALUES (8, '界了被最这', '9mkM', 2, 0, '想定很段方许争可到复龙治。', '2023-02-13 19:51:59', '2023-02-13 19:51:59', NULL);
INSERT INTO `roles` VALUES (25, '阎涛', 'n&3^1h', 1, 0, '重出六火料算确来建为照内县领向片观点。', '2023-02-14 23:28:26', '2023-02-14 23:28:26', NULL);
INSERT INTO `roles` VALUES (26, '何秀英', 'ZtG', 2, 0, '片新市线马往打务下适办影向。', '2023-02-14 23:32:26', '2023-02-14 23:32:26', NULL);
INSERT INTO `roles` VALUES (34, '梁涛', 'LF&', 1, 0, '义名看比研龙去边就过界育和求规。', '2023-02-14 23:40:10', '2023-02-14 23:40:10', NULL);
INSERT INTO `roles` VALUES (35, '刘秀英', 'iil8K(i', 1, 0, '式根专千品层且公问色半运数带第。', '2023-02-14 23:55:33', '2023-02-14 23:55:33', NULL);
INSERT INTO `roles` VALUES (36, '于桂英', 'uK@cB2', 1, 0, '文报离格它群形员变进便方历长公因该展。', '2023-02-14 23:59:33', '2023-02-14 23:59:33', NULL);
INSERT INTO `roles` VALUES (37, '苏磊', '[W19', 1, 0, '确委件广都论音及果包始第保单。', '2023-02-14 23:59:33', '2023-02-26 23:29:58', NULL);
INSERT INTO `roles` VALUES (38, '谭静', 'mno6VL', 2, 0, '但因种火些美色族分量年包影。', '2023-02-14 23:59:33', '2023-02-14 23:59:33', NULL);
INSERT INTO `roles` VALUES (39, '冯洋', 'VBnDU@M', 1, 0, '决情立性自很拉标儿细置学律劳场适。', '2023-02-14 23:59:33', '2023-02-14 23:59:33', NULL);
INSERT INTO `roles` VALUES (40, '叶静', 'm3K', 2, 0, '指省现走为张指易转系各强济。', '2023-02-15 20:42:17', '2023-02-15 20:42:17', NULL);
INSERT INTO `roles` VALUES (41, '石静', 'UV@l', 2, 0, '较石该果中用件带候约音区特往查千特它。', '2023-02-15 21:31:19', '2023-02-15 21:31:19', NULL);
INSERT INTO `roles` VALUES (42, '赖刚', '[ptM', 2, 0, '社有中农格容工然研都统温际王支除他军。', '2023-02-15 21:31:19', '2023-02-15 21:31:19', NULL);
INSERT INTO `roles` VALUES (43, '汪平', 'dpj1b[', 2, 0, '精离干压报只难适民程场法。', '2023-02-15 21:46:44', '2023-02-15 21:46:44', NULL);
INSERT INTO `roles` VALUES (45, '苏霞', 'tFhLT', 1, 0, '论形们示口根名义持法平书资内值角。', '2023-02-15 22:57:54', '2023-02-15 22:57:54', NULL);
INSERT INTO `roles` VALUES (46, '董桂英', 'YU%x', 2, 0, '每及子在回电又团持想团众被增化务位。', '2023-02-15 23:50:22', '2023-02-15 23:50:22', NULL);
INSERT INTO `roles` VALUES (47, '张磊', 'nTs4UA', 1, 0, '片受色太每利高切般改油力江金劳并压。', '2023-02-15 23:51:29', '2023-02-15 23:51:29', NULL);
INSERT INTO `roles` VALUES (48, '郝秀英', 'JfDSf', 1, 0, '类口内生压争属委般场斯运展低革。', '2023-02-17 21:35:14', '2023-02-17 21:35:14', NULL);

-- ----------------------------
-- Table structure for schema_migrations
-- ----------------------------
DROP TABLE IF EXISTS `schema_migrations`;
CREATE TABLE `schema_migrations`  (
                                      `version` bigint(20) NOT NULL,
                                      `dirty` tinyint(1) NOT NULL,
                                      PRIMARY KEY (`version`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of schema_migrations
-- ----------------------------
INSERT INTO `schema_migrations` VALUES (1, 0);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
                          `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
                          `user` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
                          `password` char(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户密码',
                          `photo` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户头像',
                          `nickname` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '昵称',
                          `name` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
                          `card` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '身份证号码',
                          `sex` tinyint(4) NOT NULL DEFAULT 3 COMMENT '性别(1为女，2为男，3为保密)',
                          `age` tinyint(4) NOT NULL DEFAULT 0 COMMENT '年龄',
                          `region` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地区',
                          `phone` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号码',
                          `email` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '邮箱',
                          `introduction` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '简介',
                          `state` tinyint(4) NOT NULL DEFAULT 2 COMMENT '状态(1已停用,2已启用)',
                          `login_out_ip` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '最后登录IP地址',
                          `login_total` int(11) NOT NULL DEFAULT 0 COMMENT '登录总数',
                          `login_out_at` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
                          `created_at` datetime NOT NULL COMMENT '创建时间',
                          `updated_at` datetime NOT NULL COMMENT '更新时间',
                          `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
                          PRIMARY KEY (`id`) USING BTREE,
                          UNIQUE INDEX `user`(`user`) USING BTREE COMMENT '用户索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'xiaohang', '$2y$10$MlLdR2kWhol1VeYSxTAJ.uMOff7QBEkBSRO82.NhPAVoLkn.iAove', NULL, '昵称测试', '小白', NULL, 3, 0, NULL, NULL, NULL, NULL, 2, NULL, 0, NULL, '2022-12-18 16:24:07', '2022-12-18 16:24:10', NULL);

SET FOREIGN_KEY_CHECKS = 1;
