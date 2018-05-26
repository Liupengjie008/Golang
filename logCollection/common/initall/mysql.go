package initall

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() (client orm.Ormer, err error) {
	// 链接数据库连接池
	maxIdle := 30
	maxConn := 30
	fmt.Println(LogConfAll.MysqlConf)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", LogConfAll.MysqlConf.UserName, LogConfAll.MysqlConf.Password, LogConfAll.MysqlConf.MysqlHost, LogConfAll.MysqlConf.MysqlPort, LogConfAll.MysqlConf.DbName)
	fmt.Println(dataSource)
	err = orm.RegisterDataBase("default", "mysql", dataSource, maxIdle, maxConn)
	// orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", maxIdle, maxConn)
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        mysql
	// 参数3        对应的链接字符串 "账号:密码@tcp(ip:端口)/数据库"
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	if err != nil {
		err = fmt.Errorf("connect mysql err : ", err)
		return
	}
	// 创建一个数据库链接 Ormer
	client = orm.NewOrm()
	logs.Error("connect mysql success")

	return
}
