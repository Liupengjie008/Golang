package server

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type CollectionConf struct {
	Id       int
	LogPath  string
	Topic    string
	Status   int
	SendRate int
}

var (
	LogConfList []CollectionConf
	TailInfoMap map[int]TailInfo
)

func LoadLogConfigFromEtcd(etcdKey string) (list []CollectionConf, err error) {
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
	// 将日志搜集信息 存入map
	TailInfoMap = SwitchoverTailInfoMap(list)

	for _, v := range TailInfoMap {
		// 启动日志搜集
		CollectionStart(v)
	}
	return
}

// 监听 ETCD 变化
func WatchLogConfigEtcd(EtcdKey string) {
	fmt.Println("watch etcd success")
	var err error
	for {
		// watch key 监听节点
		var (
			logConfList  []CollectionConf
			watchSuccess = true
		)
		WatchChan := EtcdClient.Watch(context.Background(), EtcdKey)
		for WatchResponse := range WatchChan {
			for _, Event := range WatchResponse.Events {
				if Event.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config deleted", EtcdKey)
					continue
				}
				if Event.Type == mvccpb.PUT && string(Event.Kv.Key) == EtcdKey {
					err = json.Unmarshal(Event.Kv.Value, &logConfList)
					if err != nil {
						logs.Error("key [%s], Unmarshal[%s], err:%v ", err)
						watchSuccess = false
						continue
					}
				}
			}
			if watchSuccess {
				logs.Debug("get config from etcd succ, %v", logConfList)
				MutexLock.Lock()
				LogConfList = logConfList
				MutexLock.Unlock()
				newMap := SwitchoverTailInfoMap(logConfList)
				oldMap := TailInfoMap
				// 开启新的日志搜集
				for k, v := range newMap {
					if _, ok := oldMap[k]; ok {
						continue
					}
					// 启动新增的日志搜集
					CollectionStart(v)
				}
				// 结束日志收集
				for k, v := range oldMap {
					if _, ok := newMap[k]; !ok {
						v.ExitSign = true
					}
				}

				MutexLock.Lock()
				TailInfoMap = newMap
				MutexLock.Unlock()
			}
		}
	}
}

// 生成 TailInfoMap
func SwitchoverTailInfoMap(logConfList []CollectionConf) (tailInfoMap map[int]TailInfo) {
	tailInfoMap = make(map[int]TailInfo)
	for _, v := range logConfList {
		tailInfoMap[v.Id] = TailInfo{
			LogConf: v,
		}
	}
	return
}

// 开启日志收集
func CollectionStart(t TailInfo) {

	CollectionStartNode(t)

}

func CollectionStartNode(t TailInfo) {
	// 启动 tail
	wg.Add(1)
	go TailStart(t)

	// 启动kafka
	// wg.Add(1)
	go t.SendToKafka()

}
