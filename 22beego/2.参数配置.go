配置文件目录：

project
|-- conf       
|   `-- app.conf

beego的默认参数
AppName

应用名称，默认是 beego。通过bee new创建的是创建的项目名。

AppPath

当前应用的路径，默认会通过设置os.Args[0]获得执行的命令的第一个参数，所以你在使用 supervisor 管理进程的时候记得采用全路径启动。

AppConfigPath

配置文件所在的路径，默认是应用程序对应的目录下的 conf/app.conf，用户可以修改该值配置自己的配置文件。

EnableHttpListen

是否启用HTTP监听，默认是true

beego的默认参数
HttpAddr

应用监听地址，默认为空，监听所有的网卡 IP。

HttpPort

应用监听端口，默认为 8080。

EnableHttpTLS

是否启用 HTTPS，默认是关闭。

HttpsPort

应用监听https端口，默认为 10443。

HttpCertFile

开启 HTTPS 之后，certfile 的路径。

beego的默认参数
HttpKeyFile

开启 HTTPS 之后，keyfile 的路径。

HttpServerTimeOut

设置 HTTP 的超时时间，默认是 0，不超时。

RunMode

应用的模式，默认是 dev，为开发模式，在开发模式下出错会提示友好的出错页面，如前面错误描述中所述。

AutoRender

是否模板自动渲染，默认值为 true，对于 API 类型的应用，应用需要把该选项设置为 false，不需要渲染模板。

RecoverPanic

是否异常恢复，默认值为 true，即当应用出现异常的情况，通过 recover 恢复回来，而不会导致应用异常退出。

beego的默认参数
ViewsPath

模板路径，默认值是 views。

SessionOn

session 是否开启，默认是 false。

SessionProvider

session 的引擎，默认是 memory。

SessionName

存在客户端的 cookie 名称，默认值是 beegosessionID。

SessionGCMaxLifetime

session 过期时间，默认值是 3600 秒。

beego的默认参数
SessionSavePath

session 保存路径，默认是空。

SessionHashFunc

sessionID 生成函数，默认是 sha1。

SessionHashKey

session hash 的 key。

SessionCookieLifeTime

session 默认存在客户端的 cookie 的时间，默认值是 3600 秒。

UseFcgi

是否启用 fastcgi，默认是 false。

beego的默认参数
MaxMemory

文件上传默认内存缓存大小，默认值是 1 << 26(64M)。

EnableGzip

是否开启 gzip 支持，默认为 false 不支持 gzip，一旦开启了 gzip，那么在模板输出的内容会进行 gzip 或者 zlib 压缩，根据用户的 Accept-Encoding 来判断。

DirectoryIndex

是否开启静态目录的列表显示，默认不显示目录，返回 403 错误。

BeegoServerName

beego 服务器默认在请求的时候输出 server 为 beego。

EnableAdmin

是否开启进程内监控模块，默认关闭。

beego的默认参数
AdminHttpAddr

监控程序监听的地址，默认值是 localhost。

AdminHttpPort

监控程序监听的端口，默认值是 8088。

TemplateLeft

模板左标签，默认值是{{。

TemplateRight

模板右标签，默认值是}}。

ErrorsShow

是否显示错误，默认显示错误信息。

beego的默认参数
XSRFKEY

XSRF 的 key 信息，默认值是 beegoxsrf。

XSRFExpire

XSRF 过期时间，默认值是 0。

FlashName

Flash数据设置时Cookie的名称，默认是BEEGO_FLASH

FlashSeperator

Flash数据的分隔符，默认是BEEGOFLASH

StaticDir

静态文件目录设置，默认是static

app.conf的说明
app.conf里面定义的是上面这些默认值的可覆盖值，app.conf是ini解析
[section]
key = value


beego.AppConfig.Bool("key")

Set(key, val string) error   
String(key string) string  
Strings(key string) []string
Int(key string) (int, error)
Int64(key string) (int64, error)
Bool(key string) (bool, error)
Float(key string) (float64, error)    
beego的自定义参数
如下所示的配置文件
;comment one
#comment two
appname = beeapi
httpport = 8080
mysqlport = 3600
PI = 3.1415976
runmode = "dev"
autorender = false
copyrequestbody = true
[demo]
key1="asta"
key2 = "xie"
CaseInsensitive = true
peers = one;two;three
自定义配置的读取
beego.AppConfig.Bool("autorender")
beego.AppConfig.Float("PI")
beego.AppConfig.Int("mysqlport")
beego.AppConfig.String("appname")
beego.AppConfig.String("demo::key1")
beego.AppConfig.Bool("demo::CaseInsensitive")  
beego.AppConfig.Strings("demo::peers")    [one two three]



配置文件解析
这是一个用来解析文件的库，它的设计思路来自于 database/sql，目前支持解析的文件格式有 ini、json、xml、yaml，可以通过如下方式进行安装：

go get github.com/astaxie/beego/config
如果你使用xml 或者 yaml 驱动就需要手工安装引入包

go get -u github.com/astaxie/beego/config/xml
而且需要在使用的地方引入包

import _ "github.com/astaxie/beego/config/xml"
如何使用
首先初始化一个解析器对象

iniconf, err := NewConfig("ini", "testini.conf")
if err != nil {
    t.Fatal(err)
}
然后通过对象获取数据

iniconf.String("appname")
解析器对象支持的函数有如下：

Set(key, val string) error
String(key string) string
Int(key string) (int, error)
Int64(key string) (int64, error)
Bool(key string) (bool, error)
Float(key string) (float64, error)
DIY(key string) (interface{}, error)
ini 配置文件支持 section 操作，key通过 section::key 的方式获取

例如下面这样的配置文件

[demo]
key1 = "asta"
key2 = "xie"
那么可以通过 iniconf.String("demo::key2") 获取值.

如何获取环境变量
config 模块支持环境变量配置，对应配置项 Key 格式为 ${环境变量名} ，则 Value = os.Getenv(‘环境变量名’)。
同时可配置默认值，当环境变量无此配置或者环境变量值为空时，则优先使用默认值。包含默认值的 Key 格式为 ${GOPATH||/home/workspace/go/} ，使用||分割环境变量和默认值。

注意 获取环境变量值仅仅是在配置文件解析时处理，而不会在调用函数获取配置项时实时处理。




beego config包使用：


文件目录：


|-- main.go       
|-- test.conf       

test.conf

;comment one
#comment two
appname = beeapi
httpport = 8080
mysqlport = 3600
PI = 3.1415976
runmode = "dev"
autorender = false
copyrequestbody = true
[demo]
key1="asta"
key2 = "xie"
CaseInsensitive = true
peers = one;two;three


main.go 

package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	// 初始化一个 config 解析器对象
	conf, err := config.NewConfig("ini", "test.conf")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	appname := conf.String("appname")

	fmt.Printf("get appname : %v\n", appname)

	httpport, err := conf.Int("httpport")
	if err != nil {
		fmt.Printf("get httpport err : %v\n", err)
	}
	fmt.Printf("get httpport : %v\n", httpport)

	mysqlport, err := conf.Int64("mysqlport")
	if err != nil {
		fmt.Printf("get mysqlport err : %v\n", err)
	}
	fmt.Printf("get mysqlport : %v\n", mysqlport)

	pi, err := conf.Float("PI")
	if err != nil {
		fmt.Printf("get PI err : %v\n", err)
	}
	fmt.Printf("get PI : %v\n", pi)
	autorender, err := conf.Bool("autorender")
	if err != nil {
		fmt.Printf("get autorender err : %v\n", err)
	}
	fmt.Printf("get autorender : %v\n", autorender)

	key1 := conf.String("demo::key1")
	fmt.Printf("get key1 : %v\n", key1)
	peers := conf.Strings("demo::peers")
	fmt.Println("get peers : ", peers)
}


运行结果：
get appname : beeapi
get httpport : 8080
get mysqlport : 3600
get PI : 3.1415976
get autorender : false
get key1 : asta
get peers :  [one two three]










