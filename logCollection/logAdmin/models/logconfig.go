package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

type LogConfig struct {
	Id       int       `orm:"pk;auto"`
	LogPath  string    `orm:"unique;size(255);default('')" form:"LogPath"  valid:"Required"`
	Topic    string    `orm:"unique;size(60);default('')" form:"Topic"  valid:"Required"`
	Status   int       `orm:"default(2)" form:"Status" valid:"Range(1,3)"`
	AddTime  time.Time `orm:"size(10);default(0)" form:"AddTime" valid:"Min(0)"`
	Service  string    `orm:"unique;size(60);default('')" form:"Service"  valid:"Required"`
	SendRate int       `orm:"size(10);default(0)" form:"BuyRate" valid:"Min(0)"`
}

/*
	Id       int
	LogPath  string			// 日志路径
	Topic    string			// kafka topic
	Status   int			// 是否收集 1,是 2,否
	AddTime  time.Time		// 添加时间
	Service  string			// 服务名称
	SendRate int			// 发送到 kafka 速率  条/秒
*/

func NewLogConfigModel() *LogConfig {
	return &LogConfig{}
}

func (this *LogConfig) GetLogConfigList(slimit, elimit int) (lists []LogConfig, err error) {
	if slimit == 0 && elimit == 0 {
		slimit = 0
		elimit = 20
	}
	num, err := DB.Raw("select * from log_config limit ?,?", slimit, elimit).QueryRows(&lists)
	if err != nil && num < 0 {
		logs.Warn("get log_config list err : %v ", err)
		return
	}
	return
}

func (this *LogConfig) InsertLogConfig(logConfig *LogConfig) (id int64, err error) {
	id, err = DB.Insert(logConfig)
	if err != nil {
		err = fmt.Errorf("add logConfig err : %v", err)
		logs.Warn(err)
		return
	}
	conf := CollectionConf{
		Id:       int(id),
		LogPath:  logConfig.LogPath,
		Topic:    logConfig.Topic,
		Status:   logConfig.Status,
		SendRate: logConfig.SendRate,
	}
	err = SyncLogConfigToEtcd("add", conf)
	if err != nil {
		return
	}
	return
}
