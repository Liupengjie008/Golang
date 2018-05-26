package server

import (
	"fmt"
	"logCollection/common/initall"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/coreos/etcd/clientv3"
	elastic "gopkg.in/olivere/elastic.v2"
)

var (
	err           error
	EtcdClient    *clientv3.Client
	KafkaConsumer sarama.Consumer
	EsClient      *elastic.Client
	ConfigAll     initall.ConfAll
	MutexLock     sync.Mutex

	wg sync.WaitGroup
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
	// 初始化 kafka 消费者
	if KafkaConsumer, err = initall.InitKafkaConsumer(); err != nil {
		return
	}
	// 初始化 ES
	if EsClient, err = initall.InitES(); err != nil {
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
	WatchLogConfigEtcd(ConfigAll.EtcdConf.ConfigKey)
	wg.Wait()
}
