模板处理
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

代码实现：
~~~
|-- routers
|   `-- router.go
~~~
package routers

import (
	admin "project/admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 固定路由也就是全匹配的路由
	beego.Router("/admin/user/index", &admin.UserController{}, "*:Index")
}

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

func (this *UserController) Index() {
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
    this is view admin/user/index.html
</body>
</html>

模板标签
Go语言的默认模板采用了{{和}}作为左右标签，但是我们有时候在开发中可能界面是采用了AngularJS开发，他的模板也是这个标签，故而引起了冲突。在beego中你可以通过配置文件或者直接设置配置变量修改：

beego.TemplateLeft = "<<<"
beego.TemplateRight = ">>>"



模板数据
模板中的数据是通过在Controller中this.Data获取的，所以如果你想在模板中获取内容{{.Content}},那么你需要在Controller中如下设置：

this.Data["Context"] = "value"

如何使用各种类型的数据渲染：

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

浏览器访问：
http://127.0.0.1:8080/admin/user/index

浏览器返回：
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

模板语法

go 统一使用了 {{ 和 }} 作为左右标签，没有其他的标签符号。

使用 . 来访问当前位置的上下文

使用 $ 来引用当前模板根级的上下文

使用 $var 来访问创建的变量

模板中支持的 go 语言符号

{{"string"}} // 一般 string
{{`raw string`}} // 原始 string
{{'c'}} // byte
{{print nil}} // nil 也被支持

基本语法代码：
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
	// if … else … end
	this.Data["if"] = true
	this.Data["else"] = "else"
	this.Data["elseif"] = false
	/*
		range … end
		支持的类型为 array, slice, map, channel
		range 循环内部的 . 改变为以上类型的子元素
		对应的值长度为 0 时，range 不会执行，. 不会改变
	*/
	pages := []struct {
		Num int
	}{{10}, {20}, {30}}

	this.Data["Total"] = 100
	this.Data["Pages"] = pages
	// with … end
	type stu struct {
		Name string
		Age  int
	}
	this.Data["struct"] = &stu{Name: "Murphy", Age: 28}

	// define 可以用来定义自模板，可用于模块定义和模板嵌套
	this.Data["define"] = "this is define"

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
    this is view admin/user/index.html
    <br/>
    语法 if … else … end :<br/>
    {{if .if}}
        this if = true<br/>
    {{end}}

    嵌套的循环 :<br/>
    {{if .if}}
        this if = true<br/>
        {{if .else}}
            this else = "else"
        {{end}}
    {{end}}<br/>

    else if 使用 :
    {{if .elseif}}
        this elseif = false
    {{else if .if}}
        this if = true<br/>
    {{else}}
        else is ???
    {{end}}<br/>

    语法 range … end :<br/>
    使用 .Num 输出子元素的 Num 属性，使用 $. 引用模板中的根级上下文 :<br/>
    {{range .Pages}}
        {{.Num}} of {{$.Total}}<br/>
    {{end}}

    使用创建的变量，在这里和 go 中的 range 用法是相同的 :<br/>
    {{range $index, $elem := .Pages}}
        {{$index}} - {{$elem.Num}} - {{.Num}} of {{$.Total}}<br/>
    {{end}}

    range else :<br/>
    {{range .Page}}
        {{.Num}}...
    {{else}}
        当 .Pages 为空 或者 长度为 0 时会执行这里
    {{end}}<br/>

    语法 with … end :<br/>
    with 用于重定向 :<br/>
    {{with .struct}}
        ---{{.Name}}
    {{end}}<br/>

    with 对变量赋值操作 : <br/>
    {{with $value := "My name is %s"}}
        {{printf . "Murphy"}}
    {{end}}<br/>

    with else :
    {{with .struct1}}
        this is with else true
    {{else}}
        this is with else false
        {{/* 当 pipeline 为空时会执行这里 */}}
    {{end}}<br/>

    语法 define 用来定义自模板，可用于模块定义和模板嵌套 :<br/>
    {{define "loop"}}
        define:<li>{{.define}}</li>
    {{end}}
    
    语法 template 调用模板 :<br/>
    <ul>
            template:{{template "loop" .}}
    </ul>

    <!-- 模板注释:允许多行文本注释，不允许嵌套 -->
    {{/* comment content
    support new line */}}
</body>
</html>

浏览器访问：
http://127.0.0.1:8080/admin/user/index
浏览器返回：
this is view admin/user/index.html 
语法 if … else … end :
this if = true
嵌套的循环 :
this if = true
this else = "else" 
else if 使用 : this if = true

语法 range … end :
使用 .Num 输出子元素的 Num 属性，使用 $. 引用模板中的根级上下文 :
10 of 100
20 of 100
30 of 100
使用创建的变量，在这里和 go 中的 range 用法是相同的 :
0 - 10 - 10 of 100
1 - 20 - 20 of 100
2 - 30 - 30 of 100
range else :
当 .Pages 为空 或者 长度为 0 时会执行这里 
语法 with … end :
with 用于重定向 :
---Murphy 
with 对变量赋值操作 : 
My name is Murphy 
with else : this is with else false 
语法 define 用来定义自模板，可用于模块定义和模板嵌套 :
语法 template 调用模板 :
template: define:
this is define


基本函数
变量可以使用符号 | 在函数间传递

{{.Con | markdown | addlinks}}
{{.Name | printf "%s"}}
使用括号

{{printf "nums is %s %d" (printf "%d %d" 1 2) 3}}

代码实现：

|-- admin
|     |--controllers
|	  		`-- user.go


package admin

import (
	"fmt"
	"net/url"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	// 函数 and , 函数 or
	this.Data["x"] = "x"
	this.Data["y"] = "y"
	this.Data["z"] = "z"
	// 函数 call
	this.Data["dump"] = fmt.Println
	this.Data["arg1"] = 1
	this.Data["arg2"] = 2
	// 函数 index
	this.Data["index"] = map[string]string{"this": "is", "function": "index"}
	// 函数 len
	this.Data["len"] = [10]int{}
	// 函数 not
	this.Data["not1"] = true
	this.Data["not2"] = false
	this.Data["not3"] = "not"
	this.Data["not4"] = "true"
	// 函数 urlquery
	urlencode, _ := url.ParseQuery("http://beego.me")
	this.Data["urlencode"] = urlencode.Encode()

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
    this is view admin/user/index.html
    <br/>
    函数 and :<br/>
    <!-- and 会逐一判断每个参数，将返回第一个为空的参数，否则就返回最后一个非空参数 -->
    { {and .a .b .c} } 返回空 {{and .a .b .c}}<br/>
    { {and .x .y .z} } 返回最后一个非空值 {{and .x .y .z}}<br/>

    函数 or :<br/>
    <!-- or 会逐一判断每个参数，将返回第一个非空的参数，否则就返回最后一个参数 -->
    { {and .a .b .c} } 返回最后一个参数 {{and .a .b .c}}<br/>
    { {and .x .y .z} } 返回第一个非空的参数 {{and .x .y .z}}<br/>

    函数 call :<br/>
    <!-- call 可以调用函数，并传入参数
    调用的函数需要返回 1 个值 或者 2 个值，返回两个值时，第二个值用于返回 error 类型的错误。返回的错误不等于 nil 时，执行将终止。 -->
    {{call .dump .arg .arg}}<br/>

    函数 index :<br/>
    <!-- index 支持 map, slice, array, string，读取指定类型对应下标的值 -->
    { {index .index "function"} } 下标对应的值是： {{index .index "function"}}<br/>

    函数 len :<br/>
    <!-- 返回对应类型的长度，支持类型：map, slice, array, string, chan -->
    {{printf "The len length is %d" (.len|len)}}<br/>

    函数 not :<br/>
    <!-- not 返回输入参数的否定值，if true then false else true -->
    not1 : {{not .not1}}<br/>
    not2 : {{not .not2}}<br/>
    not3 : {{not .not3}}<br/>
    not4 : {{not .not4}}<br/>

    函数 print 对应 fmt.Sprint<br/>
    函数 printf 对应 fmt.Sprintf<br/>
    函数 println 对应 fmt.Sprintln<br/>
    {{printf "x => %s , y => %s , z => %v\n" .x  .y  .z}}

    函数 urlquery :<br/>
    {{urlquery "http://beego.me"}}<br/>

    url.Encode() => {{.urlencode}}<br/>

    函数 eq / ne / lt / le / gt / ge 
    这类函数一般配合在 if 中使用

</body>
</html>

模板函数
beego 支持用户定义模板函数，但是必须在 beego.Run() 调用之前，设置如下：

func hello(in string)(out string){
    out = in + "world"
    return
}

beego.AddFuncMap("hi",hello)
定义之后你就可以在模板中这样使用了：

{{.Content | hi}}

beego 内置的模板函数：

|-- admin
|     |--controllers
|	  		`-- user.go


package admin

import (
	"time"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	// 函数 dateformat , 函数 date
	this.Data["time"] = time.Now()
	// 函数 compare
	this.Data["A"] = "A"
	this.Data["B"] = "B"
	// 函数 substr
	this.Data["str"] = "hello world！"
	// 函数 html2str
	this.Data["htmlInof"] = `<? W3S?h??！??>this is function html2str <dfdjfdk>`
	// 函数 str2html
	this.Data["strHtml"] = `<h3>this is function html2str </h3>`
	// 函数 htmlquote
	this.Data["quote"] = `<h3>this is function html2str </h3>`
	// 函数 htmlunquote
	this.Data["unquote"] = `&lt;h3&gt;this&nbsp;is&nbsp;function&nbsp;html2str&nbsp;&lt;/h3&gt;`
	// 函数 renderform
	type stu struct {
		Name  string `form:"user_name"`
		Age   int    `form:"user_age"`
		Class int
	}
	// <input name="user_name" type="text" value="">
	// <input name="user_age" type="text" value="0">
	// <input name="Class" type="text" value="0">
	this.Data["struct"] = &stu{}

	// 函数 assets_js
	this.Data["js_src"] = "./public/js/test.js"
	// <script src="./public/js/test.js"></script>

	// 函数 assets_css
	this.Data["css_src"] = "./public/css/test.css"
	// <link href="./public/css/test.css" rel="stylesheet">

	// 函数 map_get
	this.Data["map"] = map[string]interface{}{
		"key1": "value1",
		"key2": map[string]string{"key3": "value2"},
	}

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
    this is view admin/user/index.html
    <br/>
    函数 dateformat :<br/>
    <!-- 实现了时间的格式化，返回字符串 -->
    {{dateformat .time "2006-01-02 15:04:05"}} <br/>   

    函数 date :<br/>
    <!-- 实现了类似 PHP 的 date 函数，可以很方便的根据字符串返回时间 -->
    {{date .time "Y-m-d H:i:s"}}<br/>

    函数 compare :<br/>
    <!-- 实现了比较两个对象的比较，如果相同返回 true，否者 false -->
    { {compare .A .B} } => {{compare .A .B}}<br/>

    函数 substr :<br/>
    <!-- 实现了字符串的截取，支持中文截取的完美截取 -->
    {{substr .str 0 6}}<br/>
    {{substr .str 6 20}}<br/>

    函数 html2str :<br/>
    <!-- 实现了把 html 转化为字符串，剔除一些 script、css 之类的元素，返回纯文本信息 -->
    {{html2str .htmlInof}}<br/>

    函数 str2html :<br/>
    <!-- 实现了把相应的字符串当作 HTML 来输出，不转义 -->
    {{str2html .strHtml}}<br/>

    函数 htmlquote :<br/>
    <!-- 实现了基本的 html 字符转义 -->
    {{htmlquote .quote}}<br/>

    函数 htmlunquote :<br/>
    <!-- 实现了基本的反转移字符 -->
    {{htmlunquote .unquote}}<br/>

    函数 renderform :<br/>
    <!-- 根据 StructTag 直接生成对应的表单 -->
    {{.struct | renderform}}<br/>

    函数 assets_js :<br/>
    <!-- 为 js 文件生成一个 <script> 标签.  -->
    {{assets_js .js_src}}<br/>

    函数 assets_css :<br/>
    <!-- 为 css 文件生成一个 <link> 标签.  -->
    {{assets_css .css_src}}

    函数 config :<br/>
	<!-- 获取 AppConfig 的值. 可选的 configType 有 String, Bool, Int, Int64, Float, DIY -->
	{{config "String" "appname" "default"}}<br/>

    函数 map_get :<br/>
    <!-- 获取 map 的值 -->
    {{map_get .map "key1"}}<br/>
    {{map_get .map "key2" "key3"}}<br/>

    函数 urlfor :<br/>
    <!-- 获取控制器方法的 URL -->
    {{urlfor "UserController.Index"}}
</body>
</html>

lauout设计
beego支持layout设计，例如你在管理系统中，其实整个的管理界面是固定的，支会变化中间的部分，那么你可以通过如下的设置：

this.Layout = "admin/layout.html"
this.TplNames = "admin/add.tpl"
在layout.html中你必须设置如下的变量：

{{.LayoutContent}}
beego就会首先解析TplNames指定的文件，获取内容赋值给LayoutContent，然后最后渲染layout.html文件。

目前采用首先把目录下所有的文件进行缓存，所以用户还可以通过类似这样的方式实现layout：

{{template "header.html" .}}
Logic code 处理逻辑
{{template "footer.html" .}}

