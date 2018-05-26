package server

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
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

type ConsumerInfo struct {
	ExitSign    bool
	Consumer    sarama.Consumer
	Topic       string
	MessageChan chan string
}

var (
	LogConfList []CollectionConf
	ConsumerMap map[int]ConsumerInfo
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
	// 将消费者 存入map
	ConsumerMap = SwitchoverConsumerMap(list)

	for _, v := range ConsumerMap {
		// 启动消费者
		ConsumerStart(v)
	}
	return
}

// 监听 ETCD 变化
func WatchLogConfigEtcd(EtcdKey string) {
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
				newMap := SwitchoverConsumerMap(logConfList)
				oldMap := ConsumerMap
				// 开启新的日志搜集
				for k, v := range newMap {
					if _, ok := oldMap[k]; ok {
						continue
					}
					// 启动新增的日志搜集
					ConsumerStart(v)
				}
				// 结束日志收集
				for k, v := range oldMap {
					if _, ok := newMap[k]; !ok {
						v.ExitSign = true
					}
				}

				MutexLock.Lock()
				ConsumerMap = newMap
				MutexLock.Unlock()
			}
		}
	}
}

// 将 LogConfList 转成 ConsumerMap
func SwitchoverConsumerMap(list []CollectionConf) (consumerMap map[int]ConsumerInfo) {
	consumerMap = make(map[int]ConsumerInfo)
	for _, v := range list {
		consumerMap[v.Id] = ConsumerInfo{
			ExitSign:    false,
			Consumer:    KafkaConsumer,
			Topic:       v.Topic,
			MessageChan: make(chan string, 100000),
		}
	}
	return
}

func ConsumerStart(c ConsumerInfo) {
	ConsumerStartNode(c)

}

func ConsumerStartNode(c ConsumerInfo) {
	go ReadKafka(c)
	go SendToEs(c)
}
