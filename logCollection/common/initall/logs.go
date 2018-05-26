package initall

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func InitLogs() (err error) {
	logPath := LogConfAll.LogConf.LogPath
	config := fmt.Sprintf(`{"filename":"%s"}`, logPath)
	// 日志配置
	beego.SetLogger("file", config)
	// 设置级别
	beego.SetLevel(LogConfAll.LogConf.LogLevel)
	// 输出文件名和行号
	beego.SetLogFuncCall(true)

	logs.Error("init logs success")
	return
}
