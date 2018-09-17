use fchat;

create table accounts(
	id int unsigned not null auto_increment primary key,
	account varchar(16) not null,
	password varchar(16) not null,
	nickName varchar(12) null default "",
	registerTime datetime null,
	lastLoginTime datetime null
);
