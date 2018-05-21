beego ORM 使用

已支持数据库驱动：

MySQL：github.com/go-sql-driver/mysql
PostgreSQL：github.com/lib/pq
Sqlite3：github.com/mattn/go-sqlite3

安装 ORM：

go get github.com/astaxie/beego/orm


链接数据库

package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 链接数据库
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	// orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        mysql
	// 参数3        对应的链接字符串 "账号:密码@tcp(ip:端口)/数据库"
	if err != nil {
		fmt.Println("connect mysql err : ", err)
	} else {
		fmt.Println("connect mysql success")
	}

}


ORM 使用 golang 自己的连接池

package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 链接数据库连接池
	maxIdle := 30
	maxConn := 30
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", maxIdle, maxConn)
	// orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", maxIdle, maxConn)
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        mysql
	// 参数3        对应的链接字符串 "账号:密码@tcp(ip:端口)/数据库"
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	if err != nil {
		fmt.Println("connect mysql err : ", err)
	} else {
		fmt.Println("connect mysql success")
	}

}


根据数据库的别名，设置数据库的最大空闲连接

orm.SetMaxIdleConns("default", 30)

根据数据库的别名，设置数据库的最大数据库连接 (go >= 1.2)

orm.SetMaxOpenConns("default", 30)

时区设置
ORM 默认使用 time.Local 本地时区

作用于 ORM 自动创建的时间
从数据库中取回的时间转换成 ORM 本地时间
如果需要的话，你也可以进行更改

// 设置为 UTC 时间
orm.DefaultTimeLoc = time.UTC
ORM 在进行 RegisterDataBase 的同时，会获取数据库使用的时区，然后在 time.Time 类型存取时做相应转换，以匹配时间系统，从而保证时间不会出错。

Using 切换为其他数据库

orm.RegisterDataBase("db1", "mysql", "root:root@/orm_db2?charset=utf8")
orm.RegisterDataBase("db2", "sqlite3", "data.db")

o1 := orm.NewOrm()
o1.Using("db1")

o2 := orm.NewOrm()
o2.Using("db2")

// 切换为其他数据库以后
// 这个 Ormer 对象的其下的 api 调用都将使用这个数据库
默认使用 default 数据库，无需调用 Using


对象的 CRUD（增删改查） 操作

var DB orm.Ormer
DB = orm.NewOrm() 创建一个数据库链接 Ormer
DB.Insert() 插入一条数据，返回自增 id
DB.InsertMulti() 同时插入多条数据，返回插入的条数
DB.Update() 修改操作，返回值为受影响的行数
DB.Read() 从数据库读取数据
DB.ReadOrCreate() 尝试从数据库读取，不存在的话就创建一个
DB.Delete() 删除操作，返回值为受影响的行数


代码实例：

CREATE TABLE `person` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(260) DEFAULT NULL,
  `sex` varchar(260) DEFAULT NULL,
  `email` varchar(260) DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8


package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
+----------+--------------+------+-----+---------+----------------+
| Field    | Type         | Null | Key | Default | Extra          |
+----------+--------------+------+-----+---------+----------------+
| user_id  | int(11)      | NO   | PRI | NULL    | auto_increment |
| username | varchar(260) | YES  |     | NULL    |                |
| sex      | varchar(260) | YES  |     | NULL    |                |
| email    | varchar(260) | YES  |     | NULL    |                |
+----------+--------------+------+-----+---------+----------------+
*/
// struct 字段，首字母必须大写
type Person struct {
	User_id  int `orm:"pk"`
	Username string
	Sex      string
	Email    string
}

func init() {
	maxIdle := 30
	maxConn := 30
	// 创建连接池
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", maxIdle, maxConn)

	if err != nil {
		fmt.Println("connect mysql err : ", err)
	}
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Person))
}


func main() {

	var DB orm.Ormer
	// 创建一个数据库链接 Ormer
	DB = orm.NewOrm()

	p0 := Person{
		Username: "person0",
		Sex:      "man",
		Email:    "person0@golang.com",
	}
	// 插入一条数据，返回自增 id
	id, err := DB.Insert(&p0)
	if err != nil {
		fmt.Println("insert p0 err : ", err)
	}
	fmt.Println("id :", id)
	p1 := []Person{
		{Username: "person1", Sex: "man", Email: "person1@golang.com"},
		{Username: "person2", Sex: "man", Email: "person2@golang.com"},
		{Username: "person3", Sex: "man", Email: "person3@golang.com"},
		{Username: "person4", Sex: "man", Email: "person4@golang.com"},
		{Username: "person5", Sex: "man", Email: "person5@golang.com"},
		{Username: "person6", Sex: "man", Email: "person6@golang.com"},
	}
	// 同时插入多条数据，返回插入的条数
	ids, err := DB.InsertMulti(len(p1), p1)
	if err != nil {
		fmt.Println("insert p1 err : ", err)
	}
	fmt.Println("ids : ", ids)

	p2 := Person{User_id: 1}
	// 查询
	err = DB.Read(&p2)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(p2)
	}

	p3 := Person{Username: "person7"}
	// 尝试从数据库读取，不存在的话就创建一个
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	if created, id, err := DB.ReadOrCreate(&p3, "Username"); err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", id)
		} else {
			fmt.Println("Get an object. Id:", id)
		}
	}

	p4 := Person{User_id: 4}
	if DB.Read(&p4) == nil {
		p4.Username = "Update"
		p4.Sex = "woman"
		p4.Email = "update@golang.com"
		// 修改操作，返回值为受影响的行数
		if num, err := DB.Update(&p4); err == nil {
			fmt.Println("update return num : ", num)
		}
	}

	p5 := Person{User_id: 5}
	// 删除操作，返回值为受影响的行数
	if num, err := DB.Delete(&p5); err == nil {
		fmt.Println("delete return num : ", num)
	}

}

原生 SQL 查询

DB := orm.NewOrm() 创建一个数据库链接 Ormer
var ret orm.RawSeter 创建一个 RawSeter
ret = o.Raw("SQl语句")
QueryRow() QueryRow 和 QueryRows 提供高级 sql mapper 功能
QueryRows() QueryRows 支持的对象还有 map 规则是和 QueryRow 一样的，但都是 slice
Values() 返回结果集的 key => value 值
ValuesList() 返回结果集 slice
ValuesFlat() 返回单一字段的平铺 slice 数据
RowsToMap() 查询结果匹配到 map 里
RowsToStruct() 查询结果匹配到 struct 里

Exec() 执行 sql 语句，返回 sql.Result 对象
SetArgs() 改变 Raw(sql, args…) 中的 args 参数，返回一个新的 RawSeter 。用于单条 sql 语句，重复利用，替换参数然后执行。
Prepare() 用于一次 prepare 多次 exec，以提高批量执行的速度。


package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// struct 字段，首字母必须大写
type Person struct {
	User_id  int `orm:"pk"`
	Username string
	Sex      string
	Email    string
}

func init() {
	maxIdle := 30
	maxConn := 30
	// 创建连接池
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", maxIdle, maxConn)

	if err != nil {
		fmt.Println("connect mysql err : ", err)
	}
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Person))
}

func main() {

	// 创建一个数据库链接 Ormer
	DB := orm.NewOrm()
	var r0 orm.RawSeter
	// 创建一个 RawSeter
	r0 = DB.Raw("select * from person where user_id = ?", 1)
	var p1 Person
	// QueryRow 执行 sql 语句，返回 err ，查询结果存入 p1
	err := r0.QueryRow(&p1)
	if err != nil {
		fmt.Println("QueryRow err : ", err)
	}
	fmt.Println(p1)

	ids := []int{1, 2, 3}
	var p2 []Person
	// QueryRows 执行 sql 语句，返回 num,err ，查询结果 slice 存入 p2
	num, _ := DB.Raw("select * from person where user_id in (?,?,?)", ids).QueryRows(&p2)
	fmt.Println(num)
	fmt.Println(p2)

	// Values 返回结果集的 key => value 值
	var maps []orm.Params
	num, err = DB.Raw("select * from person where user_id = ?", 1).Values(&maps)
	if err == nil && num > 0 {
		fmt.Println(maps)
	}

	// ValuesList 返回结果集 slice
	var lists []orm.ParamsList
	num, err = DB.Raw("select * from person where user_id < ?", 3).ValuesList(&lists)
	if err == nil && num > 0 {
		fmt.Println(lists) // slene
	}

	// ValuesFlat 返回单一字段的平铺 slice 数据
	var list orm.ParamsList
	num, err = DB.Raw("select * from person where user_id < ?", 3).ValuesFlat(&list)
	if err == nil && num > 0 {
		fmt.Println(list) // []{"1","2","3",...}
	}

	// RowsToMap 查询结果匹配到 map 里
	res := make(orm.Params)
	num, err = DB.Raw("select * from person where user_id < ?", 3).RowsToMap(&res, "username", "email")
	if err == nil && num > 0 {
		fmt.Println("RowsToMap", res)
	}

	// RowsToStruct 查询结果匹配到 struct 里
	type Email struct {
		Person0 string
		Person1 string
	}
	res1 := new(Email)
	num, err = DB.Raw("select username,email from person where user_id < ?", 3).RowsToStruct(res1, "username", "email")
	if err == nil && num > 0 {
		fmt.Println(err, num)
		fmt.Println("RowsToStruct", res1)
	}

	// Exec() 执行 sql 语句，返回 sql.Result 对象
	act := DB.Raw("UPDATE person SET username = ? where user_id = ?", "your", 1)
	res2, err := act.Exec()
	if err == nil {
		num, _ := res2.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	// SetArgs() 用于单条 sql 语句，重复利用，替换参数然后执行。
	res3, err := act.SetArgs("your2", 2).Exec()
	if err == nil {
		num, _ := res3.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	// Prepare() 用于一次 prepare 多次 exec，以提高批量执行的速度。
	p, err := DB.Raw("UPDATE person SET username = ? WHERE user_id = ?").Prepare()
	_, _ = p.Exec("testing", 3)
	_, _ = p.Exec("testing", 4)
	p.Close() // 别忘记关闭 statement

}



构造查询

Select(fields ...string) QueryBuilder
From(tables ...string) QueryBuilder
InnerJoin(table string) QueryBuilder
LeftJoin(table string) QueryBuilder
RightJoin(table string) QueryBuilder
On(cond string) QueryBuilder
Where(cond string) QueryBuilder
And(cond string) QueryBuilder
Or(cond string) QueryBuilder
In(vals ...string) QueryBuilder
OrderBy(fields ...string) QueryBuilder
Asc() QueryBuilder
Desc() QueryBuilder
Limit(limit int) QueryBuilder
Offset(offset int) QueryBuilder
GroupBy(fields ...string) QueryBuilder
Having(cond string) QueryBuilder
Subquery(sub string, alias string) string
String() string


package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// struct 字段，首字母必须大写
type Person struct {
	User_id  int `orm:"pk"`
	Username string
	Sex      string
	Email    string
}

func init() {
	maxIdle := 30
	maxConn := 30
	// 创建连接池
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", maxIdle, maxConn)

	if err != nil {
		fmt.Println("connect mysql err : ", err)
	}
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Person))
}

func main() {
	var p0 []Person

	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	DB, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	DB.Select("*").
		From("person").
		Where("user_id > 1").
		OrderBy("user_id").Desc().
		Limit(10)

	// 导出 SQL 语句
	sql := DB.String()
	fmt.Println(sql)

	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&p0)
	fmt.Println(p0)
}



事务处理
ORM 可以简单的进行事务操作

o := NewOrm()
err := o.Begin()
// 事务处理过程
...
...
// 此过程中的所有使用 o Ormer 对象的查询都在事务处理范围内
if SomeError {
    err = o.Rollback()
} else {
    err = o.Commit()
}


beego orm 自动建表：

package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// struct 字段，首字母必须大写
type PersonCopy struct {
	User_id  int    `orm:"pk;auto"`
	Username string `orm:"size(260);null;default(NULL)"`
	Sex      string `orm:"size(260);null;default(NULL)"`
	Email    string `orm:"size(260);null;default(NULL)"`
}

func init() {
	maxIdle := 30
	maxConn := 30
	// 创建连接池
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", maxIdle, maxConn)

	if err != nil {
		fmt.Println("connect mysql err : ", err)
	}
	// 需要在init中注册定义的model
	orm.RegisterModel(new(PersonCopy))
}

func main() {
	// 自动建表
	orm.RunSyncdb("default", false, true)
}








