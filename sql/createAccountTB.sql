use fchat;

create table accounts(
	id int unsigned not null auto_increment primary key,
	account varchar(16) not null unique,
	password varchar(16) not null,
	nickName varchar(12) null default "",
	registerTime timestamp not null default now(),
	lastLoginTime timestamp not null default now()
)ENGINE=InnoDB DEFAULT CHARSET=gbk;
