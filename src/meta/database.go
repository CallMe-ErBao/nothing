package meta

//Status 0 代表正常，-1 代表账户异常(禁用)
//note 备注
var TUsers string = `
create table IF not exists users(
     id int(11) not null  auto_increment,
	 token varchar(128),
     status int(11) default 1,
     name varchar(45),
     role int(11) default 0,
     mail varchar(255),
     note varchar(512),
     icon varchar(255),
     pass  varchar(45),
     phone varchar(45),
     create_time bigint(20),
	 extra1 varchar(255),
	 extra2 varchar(255),
     primary key (id)
    )
`
var EUsers string = `
create table IF not exists event(
     id int(11) not null,
	 eType int(11),
     uid int(11),
     location_x varchar(45),
     location_y varchar(45),
     uip varchar(20),
     dType varchar(20),
	 dMan varchar(20),
	 dTypeVersion varchar(20),
	 OsVersion varchar(20),
	 appVersion varchar(20),
	 appChannel varchar(20),
	 createTime bigint(20),
	 extra1 varchar(50),
	 extra2 varchar(50),
	 primary key (id)
    )
`
var Project =`
create table IF not exists project(
     id int(64)  auto_increment,
	 name varchar(128),
	 ownerID  int(11) not null,
     status int(10),
     deleteFlag varchar(2) default 1,
     balance varchar(64),
	 extra0 varchar(128),
	 extra1 varchar(128),
	 extra2 varchar(128),
primary key (id)
    )
`