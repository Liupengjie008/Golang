package logconfig

import (
	"fmt"
	"logCollection/logAdmin/models"
	"time"

	"github.com/astaxie/beego"
)

type LogConfigController struct {
	beego.Controller
}

var LogConfigModel *models.LogConfig = models.NewLogConfigModel()

func (this *LogConfigController) Error(err interface{}) {
	url := this.Ctx.Request.Referer()
	this.Data["error"] = fmt.Sprintln(err)
	fmt.Println(err)
	this.Redirect(url, 302)
	return
}

func (this *LogConfigController) Success(url string, msg interface{}) {
	this.Data["message"] = msg
	fmt.Println(msg)
	this.Redirect(url, 302)
	return
}

func (this *LogConfigController) Index() {
	LogConfigList, err := LogConfigModel.GetLogConfigList(0, 20)
	if err != nil {
		this.Error(err)
		return
	}
	this.Data["LogConfigList"] = LogConfigList
	this.TplName = "logconfig/index.html"
	return
}

func (this *LogConfigController) AddLogConfig() {
	if this.Ctx.Input.IsGet() {
		this.TplName = "logconfig/add.html"
		return
	}
	if this.Ctx.Input.IsPost() {
		LogPath := this.GetString("LogPath")
		if len(LogPath) == 0 {
			this.Error("请输入日志路径")
			return
		}
		Topic := this.GetString("Topic")
		if len(Topic) == 0 {
			this.Error("请输入 Topic")
			return
		}
		Service := this.GetString("Service")
		if len(Service) == 0 {
			this.Error("请输入 服务名称")
			return
		}
		SendRate, err := this.GetInt("SendRate")
		if err != nil {
			err = fmt.Errorf("get 发送到 kafka 速率 err : %v", err)
			this.Error(err)
			return
		}
		Status, err := this.GetInt("Status")
		if err != nil {
			err = fmt.Errorf("get 是否收集 err : %v", err)
			this.Error(err)
			return
		}
		LogConfig := &models.LogConfig{
			LogPath:  LogPath,
			Topic:    Topic,
			Service:  Service,
			SendRate: SendRate,
			Status:   Status,
			AddTime:  time.Now().Local(),
		}
		_, err = LogConfigModel.InsertLogConfig(LogConfig)
		if err != nil {
			this.Error(err)
			return
		}
		this.Success("index", "添加成功")
		return
	}
	this.Ctx.WriteString("来了")
	return
}

func (this *LogConfigController) UpdateLogConfig() {
	this.Ctx.WriteString("页面输出字符串")
	return
}
