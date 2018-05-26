package initall

import (
	"errors"
	"fmt"
	"strings"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

type ConfAll struct {
	LogConf   LogConf
	MysqlConf MysqlConf
	EtcdConf  EtcdConf
	KafkaConf KafkaConf
	EsConf    EsConf
}

type LogConf struct {
	LogPath  string
	LogLevel int
}

type MysqlConf struct {
	MysqlHost string
	MysqlPort int
	DbName    string
	UserName  string
	Password  string
}

type EtcdConf struct {
	EtcdAddr    []string
	ConfigKey   string
	DailTimeout int
	PutTimeout  int
	GetTimeout  int
}

type KafkaConf struct {
	KafkaAddr []string
}

type EsConf struct {
	EsAddr  []string
	EsSniff bool
}

var conf config.Configer

func InitConf() (confAll ConfAll, err error) {
	conf, err = config.NewConfig("ini", "../common/conf/common.conf")
	if err != nil {
		err = errors.New(fmt.Sprintf("new config failed, err:", err))
		return
	}
	if confAll.EsConf, err = GetEsConf(); err != nil {
		return
	}
	if confAll.EtcdConf, err = GetEtcdConf(); err != nil {
		return
	}
	if confAll.KafkaConf, err = GetKafkaConf(); err != nil {
		return
	}
	if confAll.LogConf, err = GetLogConf(); err != nil {
		return
	}
	if confAll.MysqlConf, err = GetMysqlConf(); err != nil {
		return
	}
	LogConfAll = confAll
	logs.Error("load config success")
	return
}

func GetLogConf() (logConf LogConf, err error) {
	logConf.LogPath = conf.String("logs::log_path")
	if len(logConf.LogPath) == 0 {
		err = errors.New("load config of log_path failed, is null")
		return
	}
	LogLevel := conf.String("logs::log_level")
	if len(LogLevel) == 0 {
		err = errors.New("load config of log_path failed, is null")
		return
	}
	logConf.LogLevel = ConvertLogLevel(LogLevel)
	return
}
func GetMysqlConf() (mysqlConf MysqlConf, err error) {
	mysqlConf.MysqlHost = conf.String("mysql::mysql_host")
	if len(mysqlConf.MysqlHost) == 0 {
		err = errors.New("load config of mysql_host failed, is null")
		return
	}
	mysqlConf.MysqlPort, err = conf.Int("mysql::mysql_port")
	if err != nil {
		err = fmt.Errorf("load config of mysql_port err : ", err)
		return
	}
	mysqlConf.DbName = conf.String("mysql::mysql_db_name")
	if len(mysqlConf.DbName) == 0 {
		err = errors.New("load config of mysql_db_name failed, is null")
		return
	}
	mysqlConf.UserName = conf.String("mysql::mysql_user_name")
	if len(mysqlConf.UserName) == 0 {
		err = errors.New("load config of mysql_user_name failed, is null")
		return
	}
	mysqlConf.Password = conf.String("mysql::mysql_pass")
	if len(mysqlConf.Password) == 0 {
		err = errors.New("load config of mysql_pass failed, is null")
		return
	}
	return
}
func GetEtcdConf() (etcdConf EtcdConf, err error) {
	etcdAddr := conf.String("etcd::etcd_addr")
	if len(etcdAddr) == 0 {
		err = errors.New("load config of etcd_addr failed, is null")
		return
	}
	etcdConf.EtcdAddr = strings.Split(etcdAddr, ",")
	etcdConf.ConfigKey = conf.String("etcd::etcd_config_key")
	if len(etcdConf.ConfigKey) == 0 {
		err = errors.New("load config of etcd_config_key failed, is null")
		return
	}
	if strings.HasSuffix(etcdConf.ConfigKey, "/") == false {
		etcdConf.ConfigKey = etcdConf.ConfigKey + "/"
	}
	etcdConf.DailTimeout, err = conf.Int("etcd::etcd_dail_timeout")
	if err != nil {
		err = fmt.Errorf("load config of etcd_dail_timeout err : ", err)
		return
	}
	etcdConf.GetTimeout, err = conf.Int("etcd::etcd_get_timeout")
	if err != nil {
		err = fmt.Errorf("load config of etcd_get_timeout err : ", err)
		return
	}
	etcdConf.PutTimeout, err = conf.Int("etcd::etcd_put_timeout")
	if err != nil {
		err = fmt.Errorf("load config of etcd_put_timeout err : ", err)
		return
	}
	return
}
func GetKafkaConf() (kafkaConf KafkaConf, err error) {
	KafkaAddress := conf.String("kafka::kafka_addr")
	if len(KafkaAddress) == 0 {
		err = errors.New("load config of kafka_addr failed, is null")
		return
	}
	kafkaConf.KafkaAddr = strings.Split(KafkaAddress, ",")
	return
}

func GetEsConf() (esConf EsConf, err error) {
	esAddr := conf.String("ElasticSearch::es_addr")
	if len(esAddr) == 0 {
		err = errors.New("load config of es_addr failed, is null")
		return
	}
	esConf.EsAddr = strings.Split(esAddr, ",")
	esConf.EsSniff, err = conf.Bool("ElasticSearch::es_sniff")
	if err != nil {
		err = fmt.Errorf("load config of es_sniff err : ", err)
		return
	}

	return
}

func ConvertLogLevel(level string) int {
	switch level {
	case "LevelDebug":
		return logs.LevelDebug
	case "LevelInformational":
		return logs.LevelInformational
	case "LevelNotice":
		return logs.LevelNotice
	case "LevelWarning":
		return logs.LevelWarning
	case "LevelError":
		return logs.LevelError
	case "LevelCritical":
		return logs.LevelCritical
	case "LevelAlert":
		return logs.LevelAlert
	case "LevelEmergency":
		return logs.LevelEmergency
	}
	return logs.LevelDebug
}
