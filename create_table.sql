create table website(
	id int auto_increment primary key,
	name varchar(50) not null default '' comment '标签名称',
	title varchar(255) not null default '' comment '标签title',
	comment varchar(255) not null default '' comment '标签值',
	href varchar(255) not null default '' comment '标签链接',
	unique key href (href)
) engine=innodb default charset=utf8;