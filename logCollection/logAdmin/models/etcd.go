package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

type CollectionConf struct {
	Id       int
	LogPath  string
	Topic    string
	Status   int
	SendRate int
}

var (
	TypeMap map[string]int = map[string]int{"add": 1, "update": 2, "del": 3}
)

func SyncLogConfigToEtcd(Type string, conf CollectionConf) (err error) {
	TypeValue, ok := TypeMap[Type]
	if !ok {
		err = fmt.Errorf("SyncLogConfigToEtcd type is not exist ")
		return
	}
	etcdKey := ConfigAll.EtcdConf.ConfigKey
	confList, err := loadLogConfigFromEtcd(etcdKey)
	switch TypeValue {
	case 1:
		confList = append(confList, conf)
	case 2:
		for k, v := range confList {
			if v.Id == conf.Id {
				confList[k] = conf
			}
		}
	case 3:
		for k, v := range confList {
			if v.Id == conf.Id {
				confList = append(confList[:k], confList[k+1:]...)
			}
		}
	}

	jsonConfList, err := json.Marshal(confList)
	if err != nil {
		logs.Error("json marshal confList failed, err : %v", err)
		return
	}
	// 设置超时时间
	putTimeOut := time.Second * time.Duration(ConfigAll.EtcdConf.PutTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), putTimeOut)
	defer func(cancel context.CancelFunc) { cancel() }(cancel)
	// 存入ETCD
	stringConfList := string(jsonConfList)
	_, err = EtcdClient.Put(ctx, etcdKey, stringConfList)
	if err != nil {
		logs.Error("put [%s] to etcd [%s] failed, err : %v", stringConfList, etcdKey, err)
		return
	}
	fmt.Println(confList)
	return
}

func loadLogConfigFromEtcd(etcdKey string) (list []CollectionConf, err error) {
	// 设置超时时间
	getTimeOut := time.Second * time.Duration(ConfigAll.EtcdConf.GetTimeout)
	ctx, cancel := context.WithTimeout(context.Background(), getTimeOut)
	defer func(cancel context.CancelFunc) { cancel() }(cancel)
	// 取值
	logConfigInfo, err := EtcdClient.Get(ctx, etcdKey)
	if err != nil {
		logs.Error("get [%s] from etcd failed, err : %v", etcdKey, err)
		return
	}
	for _, v := range logConfigInfo.Kvs {
		err = json.Unmarshal(v.Value, &list)
		if err != nil {
			logs.Error("Unmarshal logConfigInfo failed, err : %v", err)
			return
		}
	}
	return
}
