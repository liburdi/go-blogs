CREATE TABLE `mentality_review` (
      `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
      `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '评测名称',
      `desc` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '评测描述',
      `type` TINYINT(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类。1情感;2性格;3能力;4健康;5人际;6亲子;7职场',
      `is_nice` TINYINT(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否设置为精品。1是;0不是',
      `question_ids` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '对应的题目',
      `created_at` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
      `updated_at` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '操作时间',
      `creator_id` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建人',
      `operator_id` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '操作人',
      PRIMARY KEY (`id`)
) ENGINE = INNODB DEFAULT CHARSET=utf8 COMMENT='评测'

CREATE TABLE `mentality_question` (
      `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
      `name` varchar(255) NOT NULL DEFAULT '' COMMENT '题目名称',
      `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '题目描述',
      `type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '分类。1情感;2性格;3能力;4健康;5人际;6亲子;7职场',
      `is_nice` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否设置为精品。1是;0不是',
      `created_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
      `updated_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '操作时间',
      `creator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
      `operator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '操作人',
      `is_more` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否设置为多选。1是;0不是',
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='题目表'

ALTER TABLE mentality_review ADD COLUMN pic_url VARCHAR(255) NOT NULL DEFAULT '' COMMENT '测试的图';