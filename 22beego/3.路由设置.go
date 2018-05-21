在项目中
新建admin模块，在新建模块下
新建controllers文件夹，在新建文件夹下
新建router.go

目录结构：

project
|
|-- admin
|     |--controllers
|	  		`-- user.go
|-- conf     
|   `-- app.conf
|-- controllers   
|   `-- default.go
|-- main.go
|-- models    
|-- routers
|   `-- router.go
|-- static     
|   |-- css
|   |-- img
|   `-- js
|-- tests
|   `-- default_test.go
`-- views       
    `-- index.tpl
自动匹配

|-- routers
|   `-- router.go
代码：

package routers

import (
	admin "project/admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 自动匹配
	beego.AutoRouter(&admin.UserController{})

}
|-- admin
|     |--controllers
|	  		`-- user.go
代码：

package admin

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	this.Ctx.WriteString("这是自动匹配路由 user/index")
}

func (this *UserController) Test() {
	values := this.Ctx.Input.Params()
	str := fmt.Sprintln(values)

	this.Ctx.WriteString("这是自动匹配路由 user/test , values is " + str)
}
浏览器访问：
http://127.0.0.1:8080/user/index
浏览器输出：
这是自动匹配路由 user/index

浏览器访问：
http://127.0.0.1:8080/user/test/user1/123456/28
浏览器输出：
这是自动匹配路由 user/test , values is map[:splat:user1/123456/28 0:user1 1:123456 2:28]

固定路由也就是全匹配的路由
典型的 RESTful 方式

|-- routers
|   `-- router.go
代码：

package routers

import (
	admin "project/admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 固定路由也就是全匹配的路由
	beego.Router("/admin", &admin.UserController{}, "*:Index")
	beego.Router("/admin/add/:user_name/:user_pwd:/:mobile", &admin.UserController{}, "get:Insert")
	beego.Router("/admin/index", &admin.UserController{}, "get:Index")

}
|-- admin
|     |--controllers
|	  		`-- user.go
代码：

package admin

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	this.Ctx.WriteString("这是自动匹配路由 user/index")
}

func (this *UserController) Insert() {
	user_name := this.Ctx.Input.Param(":user_name")
	values := this.Ctx.Input.Params()
	str := fmt.Sprintln(values)

	this.Ctx.WriteString("这是自动匹配路由 user/add , values is " + str + " user name is " + user_name)
}
浏览器访问：
http://127.0.0.1:8080/admin
浏览器输出：
这是自动匹配路由 user/index

浏览器访问：
http://127.0.0.1:8080/admin/index
浏览器输出：
这是自动匹配路由 user/index

浏览器访问：
http://127.0.0.1:8080/admin/add/user1/md5pwd/15688889999
浏览器输出：
这是自动匹配路由 user/add , values is map[:user_name:user1 :user_pwd:md5pwd :mobile:15688889999]
user name is user1

正则路由

|-- routers
|   `-- router.go
代码：

package routers

import (
	admin "project/admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 正则路由
	beego.Router("/admin/index", &admin.UserController{}, "*:Index")

	beego.Router("/admin/test/?:user_name", &admin.UserController{}, "get:Test")
	beego.Router("/admin/test2/:user_name", &admin.UserController{}, "get:Test")

	beego.Router("/admin/test3/:user_name:/:userpwd:/:mobile", &admin.UserController{}, "get:Test")

}
|-- admin
|     |--controllers
|	  		`-- user.go
代码：

package admin

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	this.Ctx.WriteString("这是正则路由 user/index")
}

func (this *UserController) Test() {
	user_name := this.Ctx.Input.Param(":user_name")
	values := this.Ctx.Input.Params()
	str := fmt.Sprintln(values)

	this.Ctx.WriteString("这是正则路由 user/test , values is " + str + " user name is " + user_name)
}
浏览器访问：
http://127.0.0.1:8080/admin/index
浏览器输出：
这是正则路由 user/index

浏览器访问：
http://127.0.0.1:8080/admin/test/user1
浏览器输出：
这是正则路由 user/test , values is map[:user_name:user1]
user name is user1

浏览器访问：
http://127.0.0.1:8080/admin/test2/user2
浏览器输出：
这是正则路由 user/test , values is map[:user_name:user2]
user name is user2

浏览器访问：
http://127.0.0.1:8080/admin/test3/user3/md5pwd3/15688889999
浏览器输出：
这是正则路由 user/test , values is map[:user_name:user3 :userpwd:md5pwd3 :mobile:15688889999]
user name is user3

注解路由

注意：
beego 自动会进行源码分析，注意只会在 dev 模式下进行生成，生成的路由放在 “/routers/commentsRouter.go” 文件中。
将配置文件中的 runmode 更改为 dev 模式 。
或者 在 main函数添加：beego.BConfig.RunMode ="dev"

|-- routers
|   `-- router.go
代码：

package routers

import (
	admin "project/admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 注解路由
	beego.Include(&admin.UserController{})

}
|-- admin
|     |--controllers
|	  		`-- user.go
代码：

package admin

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Test", c.Test)
}

// @router /admin/index/
func (this *UserController) Index() {
	this.Ctx.WriteString("这是注释路由 user/index")
}

// @router /admin/test/user_name/:user_name/user_id/:user_id [get]
func (this *UserController) Test() {
	user_name := this.Ctx.Input.Param(":user_name")
	values := this.Ctx.Input.Params()
	str := fmt.Sprintln(values)

	this.Ctx.WriteString("这是注释路由 user/test , values is " + str + " user name is " + user_name)
}
路由目录下自动生成 ：commentsRouter_admin_controllers.go

|-- routers
|   `-- router.go
|	`-- commentsRouter_admin_controllers.go
commentsRouter_admin_controllers.go 自动生成的代码：

package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["project/admin/controllers:UserController"] = append(beego.GlobalControllerRouter["project/admin/controllers:UserController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/admin/index/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["project/admin/controllers:UserController"] = append(beego.GlobalControllerRouter["project/admin/controllers:UserController"],
		beego.ControllerComments{
			Method: "Test",
			Router: `/admin/test/user_name/:user_name/user_id/:user_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
浏览器访问：
http://127.0.0.1:8080/admin/index
浏览器输出：
这是注释路由 user/index

浏览器访问：
http://127.0.0.1:8080/admin/test/user_name/user1/user_id/123
浏览器输出：
这是注释路由 user/test , values is map[:user_name:user1 :user_id:123]
user name is user1

namespace

|-- routers
|   `-- router.go
代码：

package routers

import (
	admin "project/admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//初始化 namespace
	ns :=
		beego.NewNamespace("/admin",
			beego.NSRouter("/user/add/:user_name:/:user_pwd:/:mobile", &admin.UserController{}, "*:Insert"),
			beego.NSNamespace("/user",
				beego.NSInclude(
					&admin.UserController{},
				),
			),
		)
	//注册 namespace
	beego.AddNamespace(ns)

}
|-- admin
|     |--controllers
|	  		`-- user.go
代码：

package admin

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Test", c.Test)
}

// @router /index/
func (this *UserController) Index() {
	this.Ctx.WriteString("这是注释路由 user/index")
}

// @router /test/user_name/:user_name/user_id/:user_id [get]
func (this *UserController) Test() {
	user_name := this.Ctx.Input.Param(":user_name")
	values := this.Ctx.Input.Params()
	str := fmt.Sprintln(values)

	this.Ctx.WriteString("这是注释路由 user/test , values is " + str + " user name is " + user_name)
}

func (this *UserController) Insert() {
	user_name := this.Ctx.Input.Param(":user_name")
	values := this.Ctx.Input.Params()
	str := fmt.Sprintln(values)

	this.Ctx.WriteString("这是namespace user/add , values is " + str + " user name is " + user_name)
}
浏览器访问：
http://127.0.0.1:8080/admin/user/index
浏览器输出：
这是注释路由 user/index

浏览器访问：
http://127.0.0.1:8080/admin/user/test/user_name/user1/user_id/123
浏览器输出：
这是注释路由 user/test , values is map[:user_name:user1 :user_id:123]
user name is user1

浏览器输出：
http://127.0.0.1:8080/admin/user/add/user1/md5pwd/15688889999
浏览器输出：
这是namespace user/add , values is map[:user_name:user1 :user_pwd:md5pwd :mobile:15688889999]
user name is user1