/*用户表*/

drop table IF EXISTS`im_user`;
CREATE TABLE `im_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `token` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'token',
  `status` int(2) DEFAULT NULL COMMENT 'status：1、online  2、offline',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `note` varchar(500) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`user_id`)
) COMMENT='用户表';


/*单聊离线消息表*/

DROP TABLE IF EXISTS `im_s_offline_msg`;
CREATE TABLE `im_s_offline_msg` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sendId` INT(11) DEFAULT NULL COMMENT '发送方ID',
  `recvId` INT(11) DEFAULT NULL COMMENT '接收方ID',
  `msgContent` VARCHAR(512) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'msgContent 内容',
  `msg_type` INT(2) DEFAULT NULL COMMENT 'type：1、text  2、url 3、图片 4、语音',
  `insert_time` DATETIME DEFAULT NULL COMMENT '插入时间',
  `note` VARCHAR(500) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) COMMENT='单聊离线消息表';


/*群聊离线消息表*/
DROP TABLE IF EXISTS `im_g_offline_msg`;
CREATE TABLE `im_g_offline_msg` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sendId` int(11) DEFAULT NULL COMMENT '发送方ID',
  `recvId` int(11) DEFAULT NULL COMMENT '接收方ID',
  `groupid` int(11) DEFAULT NULL COMMENT 'id',
  `msgContent` varchar(512) COLLATE utf8_unicode_ci DEFAULT NULL,
  `msg_type` int(2) DEFAULT NULL COMMENT 'type：1、text  2、url 3、图片 4、语音',
  `insert_time` datetime DEFAULT NULL COMMENT '插入时间',
  `note` varchar(500) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) COMMENT='群聊离线消息表';

/*消息发送记录表*/
DROP TABLE IF EXISTS `im_msg_send`;
CREATE TABLE `im_msg_send` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sendId` int(11) DEFAULT NULL COMMENT '发送方ID',
  `target_type` int(2) DEFAULT NULL COMMENT 'type：1、user  2、group',
  `targetId` int(11) DEFAULT NULL COMMENT '目标ID：收方ID或群ID',
  `msg` varchar(512) COLLATE utf8_unicode_ci DEFAULT NULL,
  `msg_type` int(2) DEFAULT NULL COMMENT 'type：1、text  2、url 3、图片 4、语音',
  `send_time` datetime DEFAULT NULL COMMENT '发送时间',
  `note` varchar(500) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) COMMENT='消息发送记录表';





