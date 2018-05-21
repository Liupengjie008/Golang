表单验证

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
	"fmt"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	valid := validation.Validation{}
	// Required 不为空，即各个类型要求不为其零值
	res := valid.Required(nil, "name")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Required 不为空 : ", res.Error.Key, res.Error.Message))
	}

	// Min(min int) 最小值，有效类型：int，其他类型都将不能通过验证
	res = valid.Min(16, 18, "min_age")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Min(min int) 最小值 : ", res.Error.Key, res.Error.Message))
	}
	// Max(max int) 最大值，有效类型：int，其他类型都将不能通过验证
	res = valid.Max(20, 19, "max_age")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Max(max int) 最大值 : ", res.Error.Key, res.Error.Message))
	}
	// Range(min, max int) 数值的范围，有效类型：int，他类型都将不能通过验证
	res = valid.Range(nil, 16, 18, "range_age")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Range(min, max int) 数值的范围 : ", res.Error.Key, res.Error.Message))
	}
	// MinSize(min int) 最小长度，有效类型：string slice，其他类型都将不能通过验证
	res = valid.MinSize(123, 5, "min_size")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("MinSize(min int) 最小长度 : ", res.Error.Key, res.Error.Message))
	}
	// MaxSize(max int) 最大长度，有效类型：string slice，其他类型都将不能通过验证
	res = valid.MaxSize(123, 2, "max_size")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("MaxSize(max int) 最大长度 : ", res.Error.Key, res.Error.Message))
	}
	// Length(length int) 指定长度，有效类型：string slice，其他类型都将不能通过验证
	res = valid.Length(0, 1, "length")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Length(length int) 指定长度 : ", res.Error.Key, res.Error.Message))
	}
	// Alpha alpha字符，有效类型：string，其他类型都将不能通过验证
	// res = valid.Alpha("", "alpha")
	res = valid.Alpha(nil, "alpha")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Alpha alpha字符 : ", res.Error.Key, res.Error.Message))
	}
	// Numeric 数字，有效类型：string，其他类型都将不能通过验证
	// res = valid.Numeric("2", "numeric")
	res = valid.Numeric(2, "numeric")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Numeric 数字 : ", res.Error.Key, res.Error.Message))
	}
	// AlphaNumeric alpha 字符或数字，有效类型：string，其他类型都将不能通过验证
	res = valid.AlphaNumeric(nil, "AlphaNumeric")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("AlphaNumeric alpha 字符或数字 : ", res.Error.Key, res.Error.Message))
	}
	// Match(pattern string) 正则匹配，有效类型：string，其他类型都将被转成字符串再匹配(fmt.Sprintf(“%v”, obj).Match)
	// res = valid.Match("123456789", regexp.MustCompile(`^(\-|\+)?\d+(\.\d+)?$`), "Match")
	res = valid.Match("abc", regexp.MustCompile(`^(\-|\+)?\d+(\.\d+)?$`), "Match")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Match(pattern string) 正则匹配 : ", res.Error.Key, res.Error.Message))
	}
	// AlphaDash alpha字符或数字或横杠-_，有效类型：string，其他类型都将不能通过验证
	res = valid.AlphaDash(nil, "AlphaDash")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("AlphaDash alpha字符或数字或横杠-_，有效类型 : ", res.Error.Key, res.Error.Message))
	}
	// Email邮箱格式，有效类型：string，其他类型都将不能通过验证
	// res = valid.Email("123456@qq.com", "email")
	res = valid.Email("123456qq.com", "email")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Email邮箱格式 : ", res.Error.Key, res.Error.Message))
	}
	// IP IP格式，目前只支持IPv4格式验证，有效类型：string，其他类型都将不能通过验证
	// res = valid.IP("192.168.0.1", "ip")
	res = valid.IP("192.168.300.1", "ip")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("IP格式 : ", res.Error.Key, res.Error.Message))
	}
	// Base64 base64编码，有效类型：string，其他类型都将不能通过验证
	// res = valid.Base64(base64.StdEncoding.EncodeToString([]byte("abc")), "base64")
	res = valid.Base64(nil, "base64")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("base64编码 : ", res.Error.Key, res.Error.Message))
	}
	// Mobile手机号，有效类型：string，其他类型都将不能通过验证
	// res = valid.Mobile("+8615621628869", "mobile")
	// res = valid.Mobile("15621628869", "mobile")
	// res = valid.Mobile(15621628869, "mobile")
	res = valid.Mobile("+861528869", "mobile")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Mobile手机号 : ", res.Error.Key, res.Error.Message))
	}
	// Tel固定电话号，有效类型：string，其他类型都将不能通过验证
	// res = valid.Tel("010-7700008", "tel")
	res = valid.Tel("15621628869", "tel")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Tel固定电话号 : ", res.Error.Key, res.Error.Message))
	}
	// Phone手机号或固定电话号，有效类型：string，其他类型都将不能通过验证
	res = valid.Phone("110", "phone")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Phone手机号或固定电话号 : ", res.Error.Key, res.Error.Message))
	}
	// ZipCode邮政编码，有效类型：string，其他类型都将不能通过验证
	// res = valid.ZipCode("100000", "zipcode")
	res = valid.ZipCode("000000", "zipcode")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("ZipCode邮政编码 : ", res.Error.Key, res.Error.Message))
	}
}


浏览器访问：
http://127.0.0.1:8080/admin/user/index
浏览器输出：
Required 不为空 :  name Can not be empty
Min(min int) 最小值 :  min_age Minimum is 18
Max(max int) 最大值 :  max_age Maximum is 19
Range(min, max int) 数值的范围 :  range_age Range is 16 to 18
MinSize(min int) 最小长度 :  min_size Minimum size is 5
MaxSize(max int) 最大长度 :  max_size Maximum size is 2
Length(length int) 指定长度 :  length Required length is 1
Alpha alpha字符 :  alpha Must be valid alpha characters
Numeric 数字 :  numeric Must be valid numeric characters
AlphaNumeric alpha 字符或数字 :  AlphaNumeric Must be valid alpha or numeric characters
Match(pattern string) 正则匹配 :  Match Must match ^(\-|\+)?\d+(\.\d+)?$
AlphaDash alpha字符或数字或横杠-_，有效类型 :  AlphaDash Must be valid alpha or numeric or dash(-_) characters
Email邮箱格式 :  email Must be a valid email address
IP格式 :  ip Must be a valid ip address
base64编码 :  base64 Must be valid base64 characters
Mobile手机号 :  mobile Must be valid mobile number
Tel固定电话号 :  tel Must be valid telephone number
Phone手机号或固定电话号 :  phone Must be valid telephone or mobile phone number
ZipCode邮政编码 :  zipcode Must be valid zipcode




定制错误信息 

|-- admin
|     |--controllers
|	  		`-- user.go


package admin

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

type User struct {
	Name string
	Age  int
}

func (this *UserController) Index() {
	u := User{"man", 40}
	valid := validation.Validation{}

	valid.Range(u.Age, 0, 18, "age")
	// 定制错误信息
	minAge := 18
	valid.Max(u.Age, minAge, "age").Message("少儿不宜！")
	// 错误信息格式化
	valid.Max(u.Age, minAge, "age").Message("%d不禁", minAge)

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			this.Ctx.WriteString(fmt.Sprintln(err.Key, err.Message))
		}
	}

}

浏览器访问：
http://127.0.0.1:8080/admin/user/index
浏览器输出：
age Range is 0 to 18
age 少儿不宜！
age 18不禁


验证函数结合 struct tag 

|-- admin
|     |--controllers
|	  		`-- user.go


package admin

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

// 验证函数写在 "valid" tag 的标签里
// 各个函数之间用分号 ";" 分隔，分号后面可以有空格
// 参数用括号 "()" 括起来，多个参数之间用逗号 "," 分开，逗号后面可以有空格
// 正则函数(Match)的匹配模式用两斜杠 "/" 括起来
// 各个函数的结果的 key 值为字段名.验证函数名
type user struct {
	Id     int
	Name   string `valid:"Required;Match(/^Bee.*/)"` // Name 不能为空并且以 Bee 开头
	Age    int    `valid:"Range(1, 140)"`            // 1 <= Age <= 140，超出此范围即为不合法
	Email  string `valid:"Email; MaxSize(100)"`      // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
	Mobile string `valid:"Mobile"`                   // Mobile 必须为正确的手机号
	IP     string `valid:"IP"`                       // IP 必须为一个正确的 IPv4 地址
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *user) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}

func (this *UserController) Index() {
	valid := validation.Validation{}
	u := user{Name: "Beego", Age: 2, Email: "dev@beego.me"}
	b, err := valid.Valid(&u)
	if err != nil {
		// handle error
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			this.Ctx.WriteString(fmt.Sprintln(err.Key, err.Message))
		}
	}
}

浏览器访问：
http://127.0.0.1:8080/admin/user/index
浏览器输出：
Mobile.Mobile Must be valid mobile number
IP.IP Must be a valid ip address




XSRF过滤

跨站请求伪造
跨站请求伪造（Cross-site request forgery ），简称为XSRF，是Web应用中常见的一个安全问题。前面的链接也详细讲述了XSRF攻击的实现方式。

当前防范XSRF的一种通用的方法，是对每个用户都记录一个无法预知的cookie数据，然后要求所有提交的请求（POST / PUT / DELETE）中都必须带有这个cookie数据。如果此数据不匹配，那么这个请求就可能是被伪造的。

beego有内建的XSRF的防范机制，要使用此机制，您需要在应用配置文件中加上enablexsrf设定：

enablexsrf = true
xsrfkey = 61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o
xsrfexpire = 3600
或者直接在主入口处这样的设置：

beego.EnableXSRF = true
beego.XSRFKEY = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
beego.XSRFExpire = 3600  //过期时间，默认1小时
如果开启了XSRF，那么beego的Web应用将对所有用户设置一个_xsrf的cookie值（默认过期1小时），如果POST PUT DELET请求中没有这个cookie值，那么这个请求会被直接拒绝。如果你开启了这个机制，那么在所有被提交的表单中，你都需要加上一个域来提供这个值。你可以通过在模板中使用专门的函数XSRFFormHTML()来做到这一点：

过期时间上面我们设置了全局的过期时间beego.XSRFExpire，但是有些时候我们也可以在控制器中修改这个过期时间，专门针对某一类处理逻辑：

func (this *HomeController) Get(){
    this.XSRFExpire = 7200
    this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
}

在表单中使用

|-- conf       
|   `-- app.conf

appname = project
httpport = 8080
runmode = dev

enablexsrf = true
xsrfkey = lalalalallalalallalalallallalla
xsrfexpire = 3600

|-- admin
|     |--controllers
|	  		`-- user.go


package admin

import (
	"fmt"
	"html/template"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	if this.Ctx.Request.Method == "GET" {
		this.XSRFExpire = 7200
		this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
		this.TplName = "admin/user/index.html"
	} else {
		_xsrf := this.GetString("_xsrf")
		xsrf := this.GetString("xsrf")

		this.Ctx.WriteString(fmt.Sprintf("_xsrf : %v\n 你的输入：%v\n", _xsrf, xsrf))
	}

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
    <form action='{{urlfor "UserController.Add"}}' method="post" enctype="multipart/form-data">
        {{ .xsrfdata }}
        <div class="field-content">
            XSRF过滤：<input name="xsrf" type="text" />
        </div> 
        <div class="field-content">
            <input type="submit" value="提交" />
        </div>
    </form>
</body>
</html>


浏览器访问：
http://127.0.0.1:8080/admin/user/index
浏览器输出：
_xsrf : GpUbw6eo4KKOySAlJurxIIE3altPhxB5
 你的输入：XSRF过滤
























