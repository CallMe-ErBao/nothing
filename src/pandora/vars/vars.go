package vars

import (
	"pandora/modules/conf"
)

//使用vars.Conf获取配置读取对象
var (
	TemplatePath string
	Conf         conf.Configuration
	///messageTypes
	MsgType_registerModule int = 1

	MsgType_request  int = 3
	MsgType_response int = 4

	//admin account
	AdminUserName string
	AdminPassword string
)
