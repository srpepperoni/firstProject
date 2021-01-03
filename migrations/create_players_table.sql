CREATE TABLE player
(
    `id`         int(11) not null auto_increment,
    `name`       varchar(150),
    `last_name`  varchar(150),
    `height`     int(10),
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) engine = InnoDB
  DEFAULT charset = utf8;