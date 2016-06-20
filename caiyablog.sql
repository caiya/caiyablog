/*
Navicat MySQL Data Transfer

Source Server         : 192.168.20.132
Source Server Version : 50711
Source Host           : localhost:3306
Source Database       : caiyablog

Target Server Type    : MYSQL
Target Server Version : 50711
File Encoding         : 65001

Date: 2016-06-20 20:40:43
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for t_blog
-- ----------------------------
DROP TABLE IF EXISTS `t_blog`;
CREATE TABLE `t_blog` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `content` text COLLATE utf8mb4_unicode_ci,
  `type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '博客类别',
  `tag` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标签（可能多个）',
  `createtime` int(11) DEFAULT NULL COMMENT '创建时间',
  `uptime` int(11) DEFAULT NULL COMMENT '更新时间',
  `readnum` int(11) DEFAULT NULL COMMENT '阅读数',
  `istop` int(1) DEFAULT NULL COMMENT '是否置顶',
  `abstract` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '摘要',
  `userid` int(11) DEFAULT NULL,
  `isdelete` int(11) DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='博客表';

-- ----------------------------
-- Records of t_blog
-- ----------------------------
INSERT INTO `t_blog` VALUES ('1', '测试标题1', '测试内容1', '1', '1', '1466060616', '1466060616', '3', '0', '摘要测试1', '1', '0');
INSERT INTO `t_blog` VALUES ('2', '测试标题2', '测试内容2', '2', '2', '1466060616', '1466060616', '10', '0', '摘要测试2', '1', '0');
INSERT INTO `t_blog` VALUES ('3', '测试标题3', '测试内容3', '2', '2', '1466060616', '1466060616', '2', '0', '摘要测试2', '1', '0');
INSERT INTO `t_blog` VALUES ('4', '测试标题4', '测试内容4', '2', '2', '1466060616', '1466060616', '2', '0', '摘要测试2', '1', '0');
INSERT INTO `t_blog` VALUES ('5', '测试标题5', '测试内容5', '2', '2', '1466060616', '1466060616', '2', '0', '摘要测试2', '1', '0');
INSERT INTO `t_blog` VALUES ('6', '测试标题6', '测试内容6', '2', '2', '1466060616', '1466060616', '2', '0', '摘要测试2', '1', '0');
INSERT INTO `t_blog` VALUES ('7', '测试标题7', '##MaHua是什么?\r\n一个在线编辑markdown文档的编辑器\r\n\r\n向Mac下优秀的markdown编辑器mou致敬\r\n\r\n##MaHua有哪些功能？\r\n\r\n* 方便的`导入导出`功能\r\n    *  直接把一个markdown的文本文件拖放到当前这个页面就可以了\r\n    *  导出为一个html格式的文件，样式一点也不会丢失\r\n* 编辑和预览`同步滚动`，所见即所得（右上角设置）\r\n* `VIM快捷键`支持，方便vim党们快速的操作 （右上角设置）\r\n* 强大的`自定义CSS`功能，方便定制自己的展示\r\n* 有数量也有质量的`主题`,编辑器和预览区域\r\n* 完美兼容`Github`的markdown语法\r\n* 预览区域`代码高亮`\r\n* 所有选项自动记忆\r\n\r\n##有问题反馈\r\n在使用中有任何问题，欢迎反馈给我，可以用以下联系方式跟我交流\r\n\r\n* 邮件(dev.hubo#gmail.com, 把#换成@)\r\n* QQ: 287759234\r\n* weibo: [@草依山](http://weibo.com/ihubo)\r\n* twitter: [@ihubo](http://twitter.com/ihubo)\r\n\r\n##捐助开发者\r\n在兴趣的驱动下,写一个`免费`的东西，有欣喜，也还有汗水，希望你喜欢我的作品，同时也能支持一下。\r\n当然，有钱捧个钱场（右上角的爱心标志，支持支付宝和PayPal捐助），没钱捧个人场，谢谢各位。\r\n\r\n##感激\r\n感谢以下的项目,排名不分先后\r\n\r\n* [mou](http://mouapp.com/) \r\n* [ace](http://ace.ajax.org/)\r\n* [jquery](http://jquery.com)\r\n\r\n##关于作者\r\n\r\n```javascript\r\n  var ihubo = {\r\n    nickName  : \"草依山\",\r\n    site : \"http://jser.me\"\r\n  }\r\n##MaHua是什么?\r\n一个在线编辑markdown文档的编辑器\r\n\r\n向Mac下优秀的markdown编辑器mou致敬\r\n\r\n##MaHua有哪些功能？\r\n\r\n* 方便的`导入导出`功能\r\n    *  直接把一个markdown的文本文件拖放到当前这个页面就可以了\r\n    *  导出为一个html格式的文件，样式一点也不会丢失\r\n* 编辑和预览`同步滚动`，所见即所得（右上角设置）\r\n* `VIM快捷键`支持，方便vim党们快速的操作 （右上角设置）\r\n* 强大的`自定义CSS`功能，方便定制自己的展示\r\n* 有数量也有质量的`主题`,编辑器和预览区域\r\n* 完美兼容`Github`的markdown语法\r\n* 预览区域`代码高亮`\r\n* 所有选项自动记忆\r\n\r\n##有问题反馈\r\n在使用中有任何问题，欢迎反馈给我，可以用以下联系方式跟我交流\r\n\r\n* 邮件(dev.hubo#gmail.com, 把#换成@)\r\n* QQ: 287759234\r\n* weibo: [@草依山](http://weibo.com/ihubo)\r\n* twitter: [@ihubo](http://twitter.com/ihubo)\r\n\r\n##捐助开发者\r\n在兴趣的驱动下,写一个`免费`的东西，有欣喜，也还有汗水，希望你喜欢我的作品，同时也能支持一下。\r\n当然，有钱捧个钱场（右上角的爱心标志，支持支付宝和PayPal捐助），没钱捧个人场，谢谢各位。\r\n\r\n##感激\r\n感谢以下的项目,排名不分先后\r\n\r\n* [mou](http://mouapp.com/) \r\n* [ace](http://ace.ajax.org/)\r\n* [jquery](http://jquery.com)\r\n\r\n##关于作者\r\n\r\n```javascript\r\n  var ihubo = {\r\n    nickName  : \"草依山\",\r\n    site : \"http://jser.me\"\r\n  }\r\n##MaHua是什么?\r\n一个在线编辑markdown文档的编辑器\r\n\r\n向Mac下优秀的markdown编辑器mou致敬\r\n\r\n##MaHua有哪些功能？\r\n\r\n* 方便的`导入导出`功能\r\n    *  直接把一个markdown的文本文件拖放到当前这个页面就可以了\r\n    *  导出为一个html格式的文件，样式一点也不会丢失\r\n* 编辑和预览`同步滚动`，所见即所得（右上角设置）\r\n* `VIM快捷键`支持，方便vim党们快速的操作 （右上角设置）\r\n* 强大的`自定义CSS`功能，方便定制自己的展示\r\n* 有数量也有质量的`主题`,编辑器和预览区域\r\n* 完美兼容`Github`的markdown语法\r\n* 预览区域`代码高亮`\r\n* 所有选项自动记忆\r\n\r\n##有问题反馈\r\n在使用中有任何问题，欢迎反馈给我，可以用以下联系方式跟我交流\r\n\r\n* 邮件(dev.hubo#gmail.com, 把#换成@)\r\n* QQ: 287759234\r\n* weibo: [@草依山](http://weibo.com/ihubo)\r\n* twitter: [@ihubo](http://twitter.com/ihubo)\r\n\r\n##捐助开发者\r\n在兴趣的驱动下,写一个`免费`的东西，有欣喜，也还有汗水，希望你喜欢我的作品，同时也能支持一下。\r\n当然，有钱捧个钱场（右上角的爱心标志，支持支付宝和PayPal捐助），没钱捧个人场，谢谢各位。\r\n\r\n##感激\r\n感谢以下的项目,排名不分先后\r\n\r\n* [mou](http://mouapp.com/) \r\n* [ace](http://ace.ajax.org/)\r\n* [jquery](http://jquery.com)\r\n\r\n##关于作者\r\n\r\n```javascript\r\n  var ihubo = {\r\n    nickName  : \"草依山\",\r\n    site : \"http://jser.me\"\r\n  }\r\n```', '2', '2', '1466060616', '1466060616', '2', '1', '摘要测试2', '1', '0');

-- ----------------------------
-- Table structure for t_category
-- ----------------------------
DROP TABLE IF EXISTS `t_category`;
CREATE TABLE `t_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '类别名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='博客类别表';

-- ----------------------------
-- Records of t_category
-- ----------------------------
INSERT INTO `t_category` VALUES ('4', 'java');
INSERT INTO `t_category` VALUES ('5', 'nodejs');
INSERT INTO `t_category` VALUES ('6', 'golang');
INSERT INTO `t_category` VALUES ('7', 'js');
INSERT INTO `t_category` VALUES ('8', 'other');

-- ----------------------------
-- Table structure for t_tag
-- ----------------------------
DROP TABLE IF EXISTS `t_tag`;
CREATE TABLE `t_tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标签名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签表，即关键字表';

-- ----------------------------
-- Records of t_tag
-- ----------------------------

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `pass` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `avatar` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
  `desc` text COLLATE utf8mb4_unicode_ci COMMENT '描述',
  `age` int(11) DEFAULT NULL,
  `sex` int(11) DEFAULT NULL,
  `major` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '专业',
  `mobile` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `nickname` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
  `regtime` int(11) DEFAULT NULL COMMENT '注册（创建）时间',
  `addr` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `birth` int(11) DEFAULT NULL,
  `qq` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `motto` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '座右铭',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES ('1', '晁州', 'e10adc3949ba59abbe56e057f20f883e', '', '个人描述部分', '23', '1', '软件开发', '18679189528', 'caiya928@aliyun.com', 'caiya', '1466156092', '陕西省西安市', '1466156092', '1419901425', '2121');
