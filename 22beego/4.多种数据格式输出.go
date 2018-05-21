beego 多种数据输出格式


直接输出字符串
通过beego.Controller.Ctx.WriteString()方法可以直接向http response body中输出字符串

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
	this.Ctx.WriteString("页面输出字符串")
}


浏览器访问：
http://127.0.0.1:8080/admin/user/index

浏览器输出：
页面输出字符串


模板数据输出

1)静态模板数据输出
通过简单的指定beego.Controller.TplName模板文件，http response body将输出模板文件对应的内容。

模板目录
beego中默认的模板目录是views，用户可以把你的模板文件放到该目录下，beego会自动在该目录下的所有模板文件进行解析并缓存，开发模式下会每次重新解析，不做缓存。当然用户可以通过如下的方式改变模板的目录：

beego.ViewsPath = "/myviewpath"

自动渲染
beego中用户无需手动的调用渲染输出模板，beego会自动的在调用玩相应的method方法之后调用Render函数，当然如果你的应用是不需要模板输出的，那么你可以在配置文件或者在main.go中设置关闭自动渲染。

AutoRender 是否模板自动渲染，默认值为 true，对于 API 类型的应用，应用需要把该选项设置为 false，不需要渲染模板。

配置文件配置如下：
autorender = true

main.go文件中设置如下：
beego.AutoRender = true

模板名称
beego采用了Go语言内置的模板引擎，所有模板的语法和Go的一模一样，至于如何写模板文件，详细的请参考模板教程。

用户通过在Controller的对应方法中设置相应的模板名称，beego会自动的在viewpath目录下查询该文件并渲染，例如下面的设置，beego会在admin下面找add.tpl文件进行渲染：

this.TplName = "admin/user/index.tpl"

我们看到上面的模板后缀名是tpl，beego默认情况下支持tpl和html后缀名的模板文件，如果你的后缀名不是这两种，请进行如下设置：

beego.AddTemplateExt("你文件的后缀名")

当你设置了自动渲染，然后在你的Controller中没有设置任何的TplName，那么beego会自动设置你的模板文件如下：

c.TplNames = c.ChildName + "/" + c.Ctx.Request.Method + "." + c.TplExt
也就是你对应的Controller名字+请求方法名.模板后缀，也就是如果你的Controller名是AddController，请求方法是POST，默认的文件后缀是tpl，那么就会默认请求/viewpath/AddController/POST.tpl文件。

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
	this.TplName = "admin/user/index.html"
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
    <title>Document</title>
</head>
<body>
    <h3>this is view admin/user/index.html</h3>
</body>
</html>

浏览器访问：
http://127.0.0.1:8080/admin/user/index

浏览器输出：
this is view admin/user/index.html


2)动态模板数据输出
在web中大部分的内容是静态的，只有少部分数据是动态的。为了复用模板的代码，需要能够把动态的数据插入到模板中，这需要特出的语法。

beego中模板通过{{}}包含需要被替换的字段，同时需要把要替换的内容添加到Controller的Data中，这样Controller执行时会自动匹配渲染模板。

模板标签
Go语言的默认模板采用了{{和}}作为左右标签，但是我们有时候在开发中可能界面是采用了AngularJS开发，他的模板也是这个标签，故而引起了冲突。在beego中你可以通过配置文件或者直接设置配置变量修改：

~~~
beego.TemplateLeft = "<<<"
beego.TemplateRight = ">>>"
~~~


模板数据
模板中的数据是通过在Controller中this.Data获取的，所以如果你想在模板中获取内容{{.Content}},那么你需要在Controller中如下设置：

this.Data["Context"] = "value"

如何使用各种类型的数据渲染：

~~~
|-- admin
|     |--controllers
|	  		`-- user.go
~~~


~~~
package admin

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	// int控制器数据赋值
	this.Data["int"] = 123456
	// float控制器数据赋值
	this.Data["float"] = 0.123456
	// string控制器数据赋值
	this.Data["string"] = "this type is string"
	// array控制器数据赋值
	this.Data["array"] = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	// map控制器数据赋值
	this.Data["map"] = map[string]string{"this": "type", "is": "map"}
	// slice控制器数据赋值
	this.Data["slice"] = []string{"s", "l", "i", "c", "e"}
	// 结构体控制器数据赋值
	type stu struct {
		Name string
		Age  int
	}
	this.Data["struct"] = &stu{Name: "stu1", Age: 18}

	this.TplName = "admin/user/index.html"
}
~~~

~~~
|-- views
|     |--admin
|			|--user
|	  			  `-- index.html
~~~

~~~
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    this is view admin/user/index.html
    <br/>
    int 数据模板渲染数据如下：{{.int}}
    <br/>
    float 数据模板渲染数据如下：{{.float}}
    <br/>
    string 数据模板渲染数据如下：{{.string}}
    <br/>
    array 数据模板渲染数据如下：<br/>
    {{range $key, $val := .array}}
    {{$key}} => {{$val}}<br/>
    {{end}}
    <br/>
    
    map 数据模板渲染数据如下：<br/>
    {{range $key, $val := .map}}
    {{$key}} => {{$val}}<br/>
    {{end}}

    {{.map.this}}=>{{.map.is}}
    <br/>
    slice 数据模板渲染数据如下：<br/>
    {{range $key, $val := .slice}}
    {{$key}} => {{$val}}<br/>
    {{end}}
    <br/>
    struct 数据模板渲染数据如下：<br/>
    the username is {{.struct.Name}}<br/>
    the age is {{.struct.Age}}
    <br/>
</body>
</html>
~~~

浏览器访问：
http://127.0.0.1:8080/admin/user/index

浏览器返回：
~~~
this is view admin/user/index.html 
int 数据模板渲染数据如下：123456 
float 数据模板渲染数据如下：0.123456 
string 数据模板渲染数据如下：this type is string 
array 数据模板渲染数据如下：
0 => 1
1 => 2
2 => 3
3 => 4
4 => 5
5 => 6
6 => 7
7 => 8

map 数据模板渲染数据如下：
is => map
this => type
type=>map 
slice 数据模板渲染数据如下：
0 => s
1 => l
2 => i
3 => c
4 => e

struct 数据模板渲染数据如下：
the username is stu1
the age is 18 
~~~

json格式数据输出

通过把要输出的数据放到Data["json"]中，然后调用ServeJSON()进行渲染，就可以把数据进行JSON序列化输出。


~~~
|-- admin
|     |--controllers
|	  		`-- user.go
~~~

package admin

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type stu struct {
	Name string
	Age  int
	Addr string
}

func (this *UserController) Index() {
	user := &stu{"Murphy", 28, "帝都"}
	this.Data["json"] = user
	this.ServeJSON()
	this.TplName = "admin/user/index.html"
}

~~~
|-- views
|     |--admin
|			|--user
|	  			  `-- index.html
~~~


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <h3>{{.json}}</h3>
</body>
</html>

浏览器访问：
http://127.0.0.1:8080/admin/user/index

浏览器输出：
{
  "Name": "Murphy",
  "Age": 28,
  "Addr": "帝都"
}

xml格式数据输出
通过把要输出的数据放到Data["xml"]中，然后调用ServeXML()进行渲染，就可以把数据进行XML序列化输出。

~~~
|-- admin
|     |--controllers
|	  		`-- user.go
~~~

package admin

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type stu struct {
	Name string
	Age  int
	Addr string
}

func (this *UserController) Index() {
	user := &stu{"Murphy", 28, "帝都"}
	this.Data["xml"] = user
	this.ServeXML()
	this.TplName = "admin/user/index.html"
}


~~~
|-- views
|     |--admin
|			|--user
|	  			  `-- index.html
~~~


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <h3>{{.xml}}</h3>
</body>
</html>

浏览器访问：
http://127.0.0.1:8080/admin/user/index

浏览器输出：
This XML file does not appear to have any style information associated with it. The document tree is shown below.
<stu>
	<Name>Murphy</Name>
	<Age>28</Age>
	<Addr>帝都</Addr>
</stu>



jsonp调用
通过把要输出的数据放到Data["jsonp"]中，然后调用ServeJSONP()进行渲染，会设置content-type为application/javascript，然后同时把数据进行JSON序列化，然后根据请求的callback参数设置jsonp输出。



~~~
|-- admin
|     |--controllers
|	  		`-- user.go
~~~

package admin

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type stu struct {
	Name string
	Age  int
	Addr string
}

func (this *UserController) Index() {
	user := &stu{"Murphy", 28, "帝都"}
	this.Data["jsonp"] = user
	this.ServeJSONP()
	this.TplName = "admin/user/index.html"
}



~~~
|-- views
|     |--admin
|			|--user
|	  			  `-- index.html
~~~


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <h3>{{.jsonp}}</h3>
</body>
</html>

浏览器访问：
http://127.0.0.1:8080/admin/user/index

浏览器输出：

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <h3>{Murphy 28 帝都}</h3>
</body>
</html>




























