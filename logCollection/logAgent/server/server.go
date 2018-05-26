package server

import (
	"fmt"
	"logCollection/common/initall"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/coreos/etcd/clientv3"
)

var (
	err         error
	EtcdClient  *clientv3.Client
	KafkaClient sarama.SyncProducer
	ConfigAll   initall.ConfAll
	MutexLock   sync.Mutex
	wg          sync.WaitGroup
)

func init() {
	if err = InitServer(); err != nil {
		panic(fmt.Sprintf("init Server failed, err:%v", err))
	}
}

func InitServer() (err error) {
	// 初始化配置
	if ConfigAll, err = initall.InitConf(); err != nil {
		return
	}
	// 初始化 etcd
	if EtcdClient, err = initall.InitEtcd(); err != nil {
		return
	}
	// 初始化 kafka
	if KafkaClient, err = initall.InitKafka(); err != nil {
		return
	}
	// 初始化 log
	if err = initall.InitLogs(); err != nil {
		return
	}
	// 读取日志搜集信息
	if LogConfList, err = LoadLogConfigFromEtcd(ConfigAll.EtcdConf.ConfigKey); err != nil {
		err = fmt.Errorf("LoadLogConfigFromEtcd err : %v", err)
		return
	}
	return
}

func Run() {
	// 监听 ETCD 变化
	go WatchLogConfigEtcd(ConfigAll.EtcdConf.ConfigKey)
	wg.Wait()
}
