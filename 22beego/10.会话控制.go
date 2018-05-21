session 操作
// 设置 session
this.SetSession("key", value)
// 获取 session
value := this.GetSession("key")
// 清空 session ，清空后 key 对应的 session value 是 nil
this.DelSession("key")

cookie 操作
// 设置 cookie
this.Ctx.SetCookie("key", value)
// 获取 cookie
this.Ctx.GetCookie("key")

过滤器 验证用户是否已经登录

beego.InsertFilter(pattern string, postion int, filter FilterFunc, skip ...bool)
InsertFilter 函数的三个必填参数，一个可选参数

pattern 路由规则，可以根据一定的规则进行路由，如果你全匹配可以用 *
postion 执行 Filter 的地方，四个固定参数如下，分别表示不同的执行过程
	BeforeRouter 寻找路由之前
	BeforeExec 找到路由之后，开始执行相应的 Controller 之前
	AfterExec 执行完 Controller 逻辑之后执行的过滤器
	FinishRouter 执行完逻辑之后执行的过滤器
filter filter 函数 type FilterFunc func(*context.Context)



代码实例：


project
|-- conf       
|   `-- app.conf

appname = project
httpport = 8080
runmode = dev
#开启 session
sessionon = true




|-- routers
|   `-- router.go

package routers

import (
	admin "project/admin/controllers"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func init() {
	// 固定路由也就是全匹配的路由
	beego.Router("/admin/user/login", &admin.UserController{}, "*:Login")
	beego.Router("/admin/user/index", &admin.UserController{}, "*:Index")
	beego.Router("/admin/user/exit", &admin.UserController{}, "*:Exit")
	// 验证用户是否已经登录
	beego.InsertFilter("/*", beego.BeforeExec, FilterUser)
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("user_name").(string)

	if !ok && ctx.Request.RequestURI != "/admin/user/login" {
		ctx.Redirect(302, "login")
	}
}



|-- admin
|     |--controllers
|	  		`-- user.go

package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Login() {
	if this.Ctx.Input.IsGet() {
		// 获取 session
		userName := this.GetSession("user_name")
		userPwd := this.GetSession("user_pwd")

		_, nameOk := userName.(string)
		_, pwdOk := userPwd.(string)
		if nameOk && pwdOk {
			// 重定向
			this.Redirect("index", 302)
		} else {
			// 获取 cookie
			this.Data["user_name"] = this.Ctx.GetCookie("user_name")
			this.Data["user_pwd"] = this.Ctx.GetCookie("user_pwd")
			this.TplName = "admin/user/login.html"
		}
	} else {
		userName := this.GetString("user_name")
		userPwd := this.GetString("user_pwd")
		// 表单验证
		valid := validation.Validation{}
		resName := valid.Required(userName, "user_name")
		resPwd := valid.Required(userPwd, "user_pwd")
		if !resName.Ok || !resPwd.Ok {
			// 重定向
			this.Redirect("login", 302)
		}
		// 设置 cookie
		this.Ctx.SetCookie("user_name", userName)
		this.Ctx.SetCookie("user_pwd", userPwd)
		// 设置 session
		this.SetSession("user_name", userName)
		this.SetSession("user_pwd", userPwd)
		this.Redirect("index", 302)
	}

}

func (this *UserController) Index() {
	user_name := this.GetSession("user_name")
	this.Data["user_name"] = user_name
	this.TplName = "admin/user/index.html"
}

func (this *UserController) Exit() {
	// 清空 session ，清空后 key 对应的 session value 是 nil
	this.DelSession("user_name")
	this.DelSession("user_pwd")
	this.Data["json"] = nil
	this.ServeJSON()
	// this.Redirect("login", 302)
}

|-- views
|     |--admin
|			|--user
|	  			  `-- index.html


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>admin/user/add</title>
</head>
<body>
    <h3>Welcome {{.user_name}}</h3>
    <!-- js post 请求 前端 js 跳转 -->
    <a href="javascript:void(0)" onclick="do_exit()">退出</a>
    <!-- a 标签 get 访问 后台重定向跳转 -->
    <!-- <a href='{{urlfor "UserController.Exit"}}' onclick="do_exit()">退出</a> -->
</body>
</html>
<script src="http://code.jquery.com/jquery-1.8.0.min.js"></script>
<script>
    function do_exit(){
        $.ajax({
            url:'{{urlfor "UserController.Exit"}}',
            data:{},
            type:"post",
            dataType:'json',
            success:function(){
                window.location.href = '/admin/user/login'
            }
        });
    }
</script>




|-- views
|     |--admin
|			|--user
|	  			  `-- login.html


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>admin/user/add</title>
</head>
<body>
    <form action='{{urlfor "UserController.Login"}}' method="post" enctype="multipart/form-data">
        <div class="field-content">
            User Name：<input name="user_name" value="{{.user_name}}" type="text" />
        </div>
        <div class="field-content">
            Password：<input name="user_pwd" value="{{.user_pwd}}" type="password" />
        </div>
        <div class="field-content">
            <input type="submit" value="提交" />
        </div>
    </form>
</body>
</html>


测试：
浏览器访问：
http://127.0.0.1:8080/admin/user/index
在未登录情况下跳转到
http://127.0.0.1:8080/admin/user/login


beego session包

代码实例：

main.go

package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

func init() {
	managerConfig := session.ManagerConfig{
		CookieName:      "cookie",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "",
	}
	// 初始化 session
	globalSessions, _ = session.NewManager("memory", &managerConfig)
	go globalSessions.GC()
}

var myTemplate *template.Template

type Person struct {
	Name string
	Pwd  string
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	// 根据当前请求返回 session 对象
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)

	if r.Method == "GET" {
		// 获取 session
		username := sess.Get("user_name")
		userpwd := sess.Get("user_pwd")
		name, errName := username.(string)
		pwd, errPwd := userpwd.(string)
		if errName && errPwd {
			p := Person{Name: name, Pwd: pwd}
			myTemplate.Execute(w, p)
		} else {
			myTemplate.Execute(w, nil)
		}

	} else {
		username := r.FormValue("user_name")
		userpwd := r.FormValue("user_pwd")
		// 设置 session
		sess.Set("user_name", username)
		sess.Set("user_pwd", userpwd)
		p := Person{Name: username, Pwd: userpwd}
		io.WriteString(w, "")
		myTemplate.Execute(w, p)
	}

}

func initTemplate(fileName string) (err error) {
	myTemplate, err = template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

func main() {
	initTemplate("./index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}


index.html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <form action='#' method="post" enctype="multipart/form-data">
        <div class="field-content">
            User Name：<input name="user_name" value="{{.Name}}" type="text" />
        </div>
        <div class="field-content">
            Password：<input name="user_pwd" value="{{.Pwd}}" type="password" />
        </div>
        <div class="field-content">
            <input type="submit" value="提交" />
        </div>
    </form>
</body>
</html>









