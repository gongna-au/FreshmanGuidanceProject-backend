drop database if exists `FreshmanGuidanceProject`;
create database `FreshmanGuidanceProject`;
use  `FreshmanGuidanceProject`;



DROP TABLE IF EXISTS persons;
CREATE TABLE `persons` (
  `id`                  int               AUTO_INCREMENT,
  `student_id`          varchar(10)           UNIQUE,
  `password`            varchar(10)     DEFAULT NULL,
  `num_of_spot`         int             DEFAULT NULL,
  `num_of_know`         int             DEFAULT NULL,




PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;



DROP TABLE IF EXISTS spots;
CREATE TABLE `spots` (
  `id`              int   ,
  `name`            varchar(30)              DEFAULT NULL,
  `introduction`    varchar(255)             DEFAULT NULL,





PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;


DROP TABLE IF EXISTS explorations;
CREATE TABLE `explorations` (
  `id`              int                      AUTO_INCREMENT,
  `student_id`      varchar(10)              DEFAULT NULL,
  `spot_id`         int                      DEFAULT NULL,




PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;




INSERT INTO spots (id,name,introduction)VALUES( 1,"华师北门","坐落于珞瑜路152号，华师er入学打卡第一站。出门即地铁站，紧邻商圈，校名由邓小平先生亲笔题写。");
INSERT INTO spots (id,name,introduction)VALUES( 2,"图书馆","桂子山上的自习圣地，白日座无虚席，夜晚亦是灯火通明。在这里，你可以尽情地采撷知识、畅游书海，感受知识的魅力。");
INSERT INTO spots (id,name,introduction)VALUES( 3,"恽代英广场","鲜花围簇中，恽代英先生的雕像正伫立于此，一袭长衫、清秀儒雅，目光更是炯炯有神，实可谓是一张鲜活的红色名片。");
INSERT INTO spots (id,name,introduction)VALUES( 4,"文华公书林","曾经的老图书馆，现在校史馆和学生事务大厅入驻其中。清晨时分，还可以看到国旗护卫队整齐有素的训练。");

INSERT INTO spots (id,name,introduction)VALUES( 5,"佑铭操场","从运动会上的挥洒汗水，到军训时期的不负青春，佑铭始终洋溢着热闹的气氛，即便是夜晚也不例外。");
INSERT INTO spots (id,name,introduction)VALUES( 6,"露天电影场","始建于20世纪50年代，如今依旧掩映在葱葱郁林之间，各大晚会均在此举办，也不要忘记每周五晚的电影放映哦。");
INSERT INTO spots (id,name,introduction)VALUES( 7,"南湖综合楼","以高科技现代化的舒适环境，赢得了南湖er的青睐和本部er的眼馋。");


INSERT INTO spots (id,name,introduction)VALUES( 8,"小电驴","华师人通勤必备单品。");
INSERT INTO spots (id,name,introduction)VALUES( 9,"牡丹园","华师的海，是三、四月份的牡丹花海。");
INSERT INTO spots (id,name,introduction)VALUES( 10,"大活","是木犀团队的常驻地哦！");
INSERT INTO spots (id,name,introduction)VALUES( 11,"西区宿舍1—6","武汉市一级保护历史建筑。");

INSERT INTO spots (id,name,introduction)VALUES( 12,"东一食堂：","临近考试周总是人从众从人。");
INSERT INTO spots (id,name,introduction)VALUES( 13,"博雅广场","忠诚博雅，朴实刚毅。");
INSERT INTO spots (id,name,introduction)VALUES( 14,"东区小树林","充满生活气息的小树林。");








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

