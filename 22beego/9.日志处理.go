日志处理
beego之前介绍的时候说过是基于几个模块搭建的，beego的日志处理是基于logs模块搭建的，内置了一个变量BeeLogger，默认已经是logs.BeeLogger类型，初始了了console，也就是默认输出到console。

使用入门
一般在程序中我们使用如下的方式进行输出：

beego.Emergency("this is emergency")
beego.Alert("this is alert")
beego.Critical("this is critical")
beego.Error("this is error")
beego.Warning("this is warning")
beego.Notice("this is notice")
beego.Informational("this is informational")
beego.Debug("this is debug")

设置输出
我们的程序往往期望把信息输出到log中，现在设置输出到文件很方便，如下所示：

beego.SetLogger("file", `{"filename":"logs/test.log"}`)
详细更多日志的请配置查看日志配置

这个默认情况就会同时输出到两个地方，一个控制台，一个文件，如果只想输出到文件，就需要调用删除操作：

beego.BeeLogger.DelLogger("console")

设置级别
日志的级别如上所示的代码这样分为八个级别：

LevelEmergency
LevelAlert
LevelCritical
LevelError
LevelWarning
LevelNotice
LevelInformational
LevelDebug

级别依次降低，默认全部打印，但是一般我们在部署环境，可以通过设置级别设置日志级别：
beego.SetLevel(beego.LevelInformational)

输出文件名和行号
日志默认不输出调用的文件名和文件行号，如果你期望输出调用的文件名和文件行号，可以如下设置

beego.SetLogFuncCall(true)
开启传入参数true，关闭传入参数false，默认是关闭的。


示例代码：

创建logs文件夹 project.log 文件

project
|
|-- logs
|		`-- project.log 


project
|-- conf       
|   `-- app.conf


appname = project
httpport = 8080
runmode = dev
#关闭自动渲染
autorender = false



|-- routers
|   `-- router.go

package routers

import (
	admin "project/admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 固定路由也就是全匹配的路由
	beego.Router("/admin/user/index", &admin.UserController{}, "*:Index")
}


|-- admin
|     |--controllers
|	  		`-- user.go

package admin

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	// 正式环境 日志配置根据需求在 beego.Run() 之前配置
	
	// 日志配置
	beego.SetLogger("file", `{"filename":"logs/project.log"}`)
	// 设置级别
	beego.SetLevel(beego.LevelDebug)
	// 输出文件名和行号
	beego.SetLogFuncCall(true)

	// 写入日志
	beego.Debug("this is debug")
	beego.Informational("this is informational")
	beego.Notice("this is notice")
	beego.Warning("this is warning")
	beego.Error("this is error")
	beego.Critical("this is critical")
	beego.Alert("this is alert")
	beego.Emergency("this is emergency")

	this.Ctx.WriteString("Run to the end")
}



浏览器访问：
http://127.0.0.1:8080/admin/user/index

查看 logs/project.log 
2018/04/20 11:10:40 [D] [user.go:19] this is debug 
2018/04/20 11:10:40 [I] [user.go:20] this is informational 
2018/04/20 11:10:40 [N] [user.go:21] this is notice 
2018/04/20 11:10:40 [W] [user.go:22] this is warning 
2018/04/20 11:10:40 [E] [user.go:23] this is error 
2018/04/20 11:10:40 [C] [user.go:24] this is critical 
2018/04/20 11:10:40 [A] [user.go:25] this is alert 
2018/04/20 11:10:40 [M] [user.go:26] this is emergency 
2018/04/20 11:10:40 [D] [server.go:2610] |      127.0.0.1|[42m 200 [0m|   1.349317ms|   match|[44m GET     [0m /admin/user/index   r:/admin/user/index



logs 模块
这是一个用来处理日志的库，它的设计思路来自于 database/sql，目前支持的引擎有 file、console、net、smtp，可以通过如下方式进行安装：

go get github.com/astaxie/beego/logs
如何使用
通用方式
首先引入包：

import (
    "github.com/astaxie/beego/logs"
)
然后添加输出引擎（log 支持同时输出到多个引擎），这里我们以 console 为例，第一个参数是引擎名（包括：console、file、conn、smtp、es、multifile）

logs.SetLogger("console")
添加输出引擎也支持第二个参数,用来表示配置信息，详细的配置请看下面介绍：

logs.SetLogger(logs.AdapterFile,`{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
然后我们就可以在我们的逻辑中开始任意的使用了：

package main

import (
    "github.com/astaxie/beego/logs"
)

func main() {
    //an official log.Logger
    l := logs.GetLogger()
    l.Println("this is a message of http")
    //an official log.Logger with prefix ORM
    logs.GetLogger("ORM").Println("this is a message of orm")

    logs.Debug("my book is bought in the year of ", 2016)
    logs.Info("this %s cat is %v years old", "yellow", 3)
    logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
    logs.Error(1024, "is a very", "good game")
    logs.Critical("oh,crash")
}
多个实例
一般推荐使用通用方式进行日志，但依然支持单独声明来使用独立的日志

    package main

    import (
        "github.com/astaxie/beego/logs"
    )

    func main() {
        log := logs.NewLogger()
        log.SetLogger(logs.AdapterConsole)
        log.Debug("this is a debug message")
    }
输出文件名和行号
日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置

logs.EnableFuncCallDepth(true)
开启传入参数 true,关闭传入参数 false,默认是关闭的.

如果你的应用自己封装了调用 log 包,那么需要设置 SetLogFuncCallDepth,默认是 2,也就是直接调用的层级,如果你封装了多层,那么需要根据自己的需求进行调整.

logs.SetLogFuncCallDepth(3)
异步输出日志
为了提升性能, 可以设置异步输出:

logs.Async()
异步输出允许设置缓冲 chan 的大小

logs.Async(1e3)
引擎配置设置
console

可以设置输出的级别，或者不设置保持默认，默认输出到 os.Stdout：

logs.SetLogger(logs.AdapterConsole, `{"level":1}`)
file

设置的例子如下所示：

logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)
主要的参数如下说明：

filename 保存的文件名
maxlines 每个文件保存的最大行数，默认值 1000000
maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
daily 是否按照每天 logrotate，默认是 true
maxdays 文件最多保存多少天，默认保存 7 天
rotate 是否开启 logrotate，默认是 true
level 日志保存的时候的级别，默认是 Trace 级别
perm 日志文件权限
multifile

设置的例子如下所示：

logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
主要的参数如下说明(除 separate 外,均与file相同)：

filename 保存的文件名
maxlines 每个文件保存的最大行数，默认值 1000000
maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
daily 是否按照每天 logrotate，默认是 true
maxdays 文件最多保存多少天，默认保存 7 天
rotate 是否开启 logrotate，默认是 true
level 日志保存的时候的级别，默认是 Trace 级别
perm 日志文件权限
separate 需要单独写入文件的日志级别,设置后命名类似 test.error.log
conn

网络输出，设置的例子如下所示：

logs.SetLogger(logs.AdapterConn, `{"net":"tcp","addr":":7020"}`)
主要的参数说明如下：

reconnectOnMsg 是否每次链接都重新打开链接，默认是 false
reconnect 是否自动重新链接地址，默认是 false
net 发开网络链接的方式，可以使用 tcp、unix、udp 等
addr 网络链接的地址
level 日志保存的时候的级别，默认是 Trace 级别
smtp

邮件发送，设置的例子如下所示：

logs.SetLogger(logs.AdapterMail, `{"username":"beegotest@gmail.com","password":"xxxxxxxx","host":"smtp.gmail.com:587","sendTos":["xiemengjun@gmail.com"]}`)
主要的参数说明如下：

username smtp 验证的用户名
password smtp 验证密码
host 发送的邮箱地址
sendTos 邮件需要发送的人，支持多个
subject 发送邮件的标题，默认是 Diagnostic message from server
level 日志发送的级别，默认是 Trace 级别
ElasticSearch

输出到 ElasticSearch:

logs.SetLogger(logs.AdapterEs, `{"dsn":"http://localhost:9200/","level":1}`)
简聊

输出到简聊：

logs.SetLogger(logs.AdapterJianLiao, `{"authorname":"xxx","title":"beego", "webhookurl":"https://jianliao.com/xxx", "redirecturl":"https://jianliao.com/xxx","imageurl":"https://jianliao.com/xxx","level":1}`)
slack

输出到slack:

logs.SetLogger(logs.AdapterSlack, `{"webhookurl":"https://slack.com/xxx","level":1}`)



示例代码：

创建logs文件夹 project.log 文件

project
|
|-- logs
|		`-- project.log 


project
|-- conf       
|   `-- app.conf


appname = project
httpport = 8080
runmode = dev
#关闭自动渲染
autorender = false



|-- routers
|   `-- router.go

package routers

import (
	admin "project/admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 固定路由也就是全匹配的路由
	beego.Router("/admin/user/index", &admin.UserController{}, "*:Index")
}


|-- admin
|     |--controllers
|	  		`-- user.go

package admin

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	config := make(map[string]interface{})
	// config["filename"] 日志保存路径，文件必须存在。（不存在不会报错，但是不会自动创建）
	config["filename"] = "./logs/project.log"
	// config["level"] 日志级别
	config["level"] = logs.LevelDebug

	configJson, err := json.Marshal(config)
	if err != nil {
		this.Ctx.WriteString(fmt.Sprintf("json marshal config err : %v", err))
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configJson))

	logs.Debug("this is a test, my name is %s", "debug")
	logs.Trace("this is a trace, my name is %s", "trace")
	logs.Warn("this is a warn, my name is %s", "warn")

	this.Ctx.WriteString("Run to the end")
}


浏览器访问：
http://127.0.0.1:8080/admin/user/index

查看 logs/project.log 
2018/04/20 10:26:02 [D] [asm_amd64.s:509] this is a test, my name is debug
2018/04/20 10:26:02 [D] [asm_amd64.s:509] this is a trace, my name is trace
2018/04/20 10:26:02 [W] [asm_amd64.s:509] this is a warn, my name is warn
2018/04/20 10:26:02 [D] [server.go:2610] |      127.0.0.1|[42m 200 [0m|   7.935733ms|   match|[44m GET     [0m /admin/user/index   r:/admin/user/index





beego logs包使用：


package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	// config["filename"] 日志保存路径，文件必须存在。（不存在不会报错，但是不会自动创建）
	config["filename"] = "./test.log"
	// config["level"] 日志级别
	config["level"] = logs.LevelDebug

	configJson, err := json.Marshal(config)
	if err != nil {
		fmt.Printf("json marshal config err : %v", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configJson))

	logs.Debug("this is a test, my name is %s", "debug")
	logs.Trace("this is a trace, my name is %s", "trace")
	logs.Warn("this is a warn, my name is %s", "warn")

}


















