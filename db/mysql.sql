/*�û���*/

drop table IF EXISTS`im_user`;
CREATE TABLE `im_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '�û�ID',
  `token` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'token',
  `status` int(2) DEFAULT NULL COMMENT 'status��1��online  2��offline',
  `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `note` varchar(500) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '��ע',
  PRIMARY KEY (`user_id`)
) COMMENT='�û���';


/*����������Ϣ��*/

DROP TABLE IF EXISTS `im_s_offline_msg`;
CREATE TABLE `im_s_offline_msg` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sendId` INT(11) DEFAULT NULL COMMENT '���ͷ�ID',
  `recvId` INT(11) DEFAULT NULL COMMENT '���շ�ID',
  `msgContent` VARCHAR(512) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'msgContent ����',
  `msg_type` INT(2) DEFAULT NULL COMMENT 'type��1��text  2��url 3��ͼƬ 4������',
  `insert_time` DATETIME DEFAULT NULL COMMENT '����ʱ��',
  `note` VARCHAR(500) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '��ע',
  PRIMARY KEY (`id`)
) COMMENT='����������Ϣ��';


/*Ⱥ��������Ϣ��*/
DROP TABLE IF EXISTS `im_g_offline_msg`;
CREATE TABLE `im_g_offline_msg` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sendId` int(11) DEFAULT NULL COMMENT '���ͷ�ID',
  `recvId` int(11) DEFAULT NULL COMMENT '���շ�ID',
  `groupid` int(11) DEFAULT NULL COMMENT 'id',
  `msgContent` varchar(512) COLLATE utf8_unicode_ci DEFAULT NULL,
  `msg_type` int(2) DEFAULT NULL COMMENT 'type��1��text  2��url 3��ͼƬ 4������',
  `insert_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `note` varchar(500) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '��ע',
  PRIMARY KEY (`id`)
) COMMENT='Ⱥ��������Ϣ��';

/*��Ϣ���ͼ�¼��*/
DROP TABLE IF EXISTS `im_msg_send`;
CREATE TABLE `im_msg_send` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `sendId` int(11) DEFAULT NULL COMMENT '���ͷ�ID',
  `target_type` int(2) DEFAULT NULL COMMENT 'type��1��user  2��group',
  `targetId` int(11) DEFAULT NULL COMMENT 'Ŀ��ID���շ�ID��ȺID',
  `msg` varchar(512) COLLATE utf8_unicode_ci DEFAULT NULL,
  `msg_type` int(2) DEFAULT NULL COMMENT 'type��1��text  2��url 3��ͼƬ 4������',
  `send_time` datetime DEFAULT NULL COMMENT '����ʱ��',
  `note` varchar(500) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '��ע',
  PRIMARY KEY (`id`)
) COMMENT='��Ϣ���ͼ�¼��';





