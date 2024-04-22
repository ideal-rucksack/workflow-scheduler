package entities

import "time"

//	 name varchar(255) not null comment '数据源名称',
//		type varchar(32) not null comment '数据源类型',
//		connect_type varchar(32) not null comment '连接类型',
//		host varchar(255) not null comment '数据源地址',
//		port int not null comment '数据源端口',
//		username varchar(255) not null comment '数据源用户名',
//		password varchar(255) not null comment '数据源密码',
//		database_name varchar(255) not null comment '数据库名称',
//		extra text comment '额外配置',
//		protocol varchar(32) not null comment '数据源协议',
//		status varchar(32) not null comment '数据源状态',
//		testing_connect_at datetime comment '测试连接时间',
type Datasource struct {
	GlobalEntity
	Name        *string    `json:"name"`
	Type        *string    `json:"type"`
	ConnectType *string    `json:"connect_type"`
	Host        *string    `json:"host"`
	Port        *int       `json:"port"`
	Username    *string    `json:"username"`
	Password    *string    `json:"password"`
	Database    *string    `json:"database_name"`
	Extra       *string    `json:"extra"`
	Protocol    *string    `json:"protocol"`
	Status      *string    `json:"status"`
	TestingAt   *time.Time `json:"testing_connect_at"`
}
