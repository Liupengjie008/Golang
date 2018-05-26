package models

import (
	"fmt"
	"logCollection/common/initall"

	"github.com/astaxie/beego/orm"
	"github.com/beego/admin/src/models"
	"github.com/coreos/etcd/clientv3"
)

var (
	DB         orm.Ormer
	EtcdClient *clientv3.Client
	ConfigAll  initall.ConfAll
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(LogConfig))
	//数据库连接
	models.Connect()
	// 自动建表
	orm.RunSyncdb("default", false, true)

	DB = orm.NewOrm()

	if err := initAll(); err != nil {
		panic(fmt.Sprintln("init database failed, err:%v", err))
	}
}

func initAll() (err error) {
	if ConfigAll, err = initall.InitConf(); err != nil {
		return
	}

	if EtcdClient, err = initall.InitEtcd(); err != nil {
		return
	}
	if err = initall.InitLogs(); err != nil {
		return
	}
	return
}
