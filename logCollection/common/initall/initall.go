package initall

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/coreos/etcd/clientv3"
	elastic "gopkg.in/olivere/elastic.v2"
)

var (
	MysqlClient orm.Ormer
	EtcdClient  *clientv3.Client
	KafkaClient sarama.SyncProducer
	EsClient    *elastic.Client
	LogConfAll  ConfAll
)

func InitAll() (err error) {
	if LogConfAll, err = InitConf(); err != nil {
		return
	}
	if err = InitLogs(); err != nil {
		return
	}
	if MysqlClient, err = InitMysql(); err != nil {
		return
	}
	if EtcdClient, err = InitEtcd(); err != nil {
		return
	}
	if KafkaClient, err = InitKafka(); err != nil {
		return
	}
	if EsClient, err = InitES(); err != nil {
		return
	}
	logs.Error("init all success")
	return
}
