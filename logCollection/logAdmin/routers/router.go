package routers

import (
	"admin/controllers" //自身业务包
	"logCollection/logAdmin/controllers/logconfig"
	_ "logCollection/logAdmin/models"

	"github.com/astaxie/beego" //beego 包
	"github.com/beego/admin"   //admin 包
)

func init() {
	admin.Run()
	beego.Router("/", &controllers.MainController{})

	beego.Router("/elk/LogConfig/index", &logconfig.LogConfigController{}, "*:Index")
	beego.Router("/elk/LogConfig/add", &logconfig.LogConfigController{}, "*:AddLogConfig")
	beego.Router("/elk/LogConfig/update", &logconfig.LogConfigController{}, "*:UpdateLogConfig")
}
