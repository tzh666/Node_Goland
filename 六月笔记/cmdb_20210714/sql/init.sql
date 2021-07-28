-- crette databases cmdb
create  database if not exists cmdb default charset utf8mb4;

--  use cmdb
use cmdb

-- crette table
create table if not exists user(
    id bigint primary key auto_increment,
    staff_id varchar(32) not null default '',
    name varchar(64) not null default '',
    nickname varchar(64) not null default '',
    password varchar(1024) not null default '',
    gender int not null default 0 comment '0: 女、1: 男',
    tel varchar(32) not null default '',
    email varchar(64) not null default '',
    addr varchar(128) not null default '',
    department varchar(128) not null default '',
    status int not null default 0 comment '0: 正常、1: 锁定、2: 离职',
    create_at datetime not null,
    update_at datetime not null,
    delete_at datetime
)engine=innodb default charset utf8mb4;



insert into user(staff_id,name,nickname,password,gender,tel,email,addr,department,status,create_at,update_at) values
("k00001","kk","kk",md5("123456"),1,"12365489654","7777@qq.com","深圳","运维开发",0,now(),now());