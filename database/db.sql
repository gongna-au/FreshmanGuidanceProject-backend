use  `Test`;



CREATE TABLE `persons` (
  `id`                  int               AUTO_INCREMENT,
  `student_id`          varchar(10)           UNIQUE,
  `password`            varchar(10)     DEFAULT NULL,
  `num_of_spot`         int             DEFAULT NULL,
  `num_of_know`         int             DEFAULT NULL,




PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;










/*use  `FreshmanGuidance`;
-- 用户信息表(user)


CREATE TABLE `users` (
  `id`                  int  null    AUTO_INCREMENT  comment      "id",
  `username`            varchar(10)    DEFAULT NULL     comment      "学号",
  `password`            varchar(10)    DEFAULT NULL     comment      "密码",
  `num_of_landmarks`    int            DEFAULT NULL     comment      "地标数",
  `num_of_knowledge`    int            DEFAULT NULL     comment      "见闻数",


PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;


 /*`id`                    int not null    AUTO_INCREMENT  comment      "用户id",
    `student_id`            varchar(10)     UNIQUE          comment      "学号",
    `password`              varchar(20)     null            comment      "密码",
    `number_of_landmarks`   int             null            comment      "地标数",
    `number_of_knowledge`   int             null            comment      "见闻数",
-- 添加约束
primary key (`id`),
key `student_id` (`student_id`),
key `password` (`password`)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;*/

