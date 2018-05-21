from表单get请求：

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
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/get_add", &admin.UserController{}, "*:GetAdd")
}



|-- admin
|     |--controllers
|	  		`-- user.go

package admin

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	this.TplName = "admin/user/index.html"
}

func (this *UserController) GetAdd() {
	this.TplName = "admin/user/add.html"
}

func (this *UserController) Add() {
	// Get方式的请求，GetString 获取数据
	var str string = this.GetString("get_string")
	this.Ctx.WriteString(fmt.Sprintf("get string : %s\n", str))
	// Get方式的请求，GetStrings 获取数据
	var strs []string = this.GetStrings("get_strings")
	this.Ctx.WriteString(fmt.Sprintf("get strings : %v\n", strs))
	// Get方式的请求，GetInt 获取数据
	int_ret, _ := this.GetInt("get_int")
	this.Ctx.WriteString(fmt.Sprintf("get int64 : %v\n", int_ret))
	// Get方式的请求，GetBool 获取数据
	bool_ret, _ := this.GetBool("get_bool")
	this.Ctx.WriteString(fmt.Sprintf("get bool : %v\n", bool_ret))
	// Get方式的请求，GetFloat 获取数据
	float_ret, _ := this.GetFloat("get_float")
	this.Ctx.WriteString(fmt.Sprintf("get float : %v\n", float_ret))

}


|-- views
|     |--admin
|			|--user
|	  			  `-- add.html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>admin/user/add</title>
</head>
<body>
    this is admin/user/add
    <form action='{{urlfor "UserController.Add"}}' method="get">
        <div class="field-content">
            GetString：<input name="get_string" type="text" />
        </div>
        <div class="field-content">
            GetStrings：<input name="get_strings" type="text" />
            GetStrings：<input name="get_strings" type="text" />
        </div>
        <div class="field-content">
            GetInt：<input name="get_int" type="text" />
        </div>
        <div class="field-content">
            GetBool：<input name="get_bool" type="text" />
        </div>
        <div class="field-content">
            GetFloat：<input name="get_float" type="text" />
        </div>
        <div class="field-content">
            <input type="submit" value="提交" />
        </div>
    </form>
</body>
</html>

from表单post请求:

|-- admin
|     |--controllers
|	  		`-- user.go

package admin

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	this.TplName = "admin/user/index.html"
}

func (this *UserController) Add() {
	// this.Ctx.Request.Method 获取请求方式
	if this.Ctx.Request.Method == "GET" {
		this.TplName = "admin/user/add.html"
	} else {
		// Post方式的请求，GetString 获取数据
		var str string = this.GetString("post_string")
		this.Ctx.WriteString(fmt.Sprintf("post string : %s\n", str))
		// Post方式的请求，GetStrings 获取数据
		var strs []string = this.GetStrings("post_strings")
		this.Ctx.WriteString(fmt.Sprintf("post strings : %v\n", strs))
		// Post方式的请求，GetInt 获取数据
		int_ret, _ := this.GetInt("post_int")
		this.Ctx.WriteString(fmt.Sprintf("post int64 : %v\n", int_ret))
		// Post方式的请求，GetBool 获取数据
		bool_ret, _ := this.GetBool("post_bool")
		this.Ctx.WriteString(fmt.Sprintf("post bool : %v\n", bool_ret))
		// Post方式的请求，GetFloat 获取数据
		float_ret, _ := this.GetFloat("post_float")
		this.Ctx.WriteString(fmt.Sprintf("post float : %v\n", float_ret))
	}

}



|-- views
|     |--admin
|			|--user
|	  			  `-- add.html


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>admin/user/add</title>
</head>
<body>
    this is admin/user/add
    <form action='{{urlfor "UserController.Add"}}' method="post" enctype="multipart/form-data">
        <div class="field-content">
            PostString：<input name="post_string" type="text" />
        </div>
        <div class="field-content">
            PostStrings：<input name="post_strings" type="text" />
            PostStrings：<input name="post_strings" type="text" />
        </div>
        <div class="field-content">
            PostInt：<input name="post_int" type="text" />
        </div>
        <div class="field-content">
            PostBool：<input name="post_bool" type="text" />
        </div>
        <div class="field-content">
            PostFloat：<input name="post_float" type="text" />
        </div>
        <div class="field-content">
            <input type="submit" value="提交" />
        </div>
    </form>
</body>
</html>


from表单文件上传：

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
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
}


|-- admin
|     |--controllers
|	  		`-- user.go

package admin

import (
	"path"
	"strings"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	this.TplName = "admin/user/index.html"
}

func (this *UserController) Add() {
	if this.Ctx.Request.Method == "GET" {
		this.TplName = "admin/user/add.html"
	} else {
		//image，这是一个key值，对应的是html中input type-‘file’的name属性值
		f, h, _ := this.GetFile("image")
		//得到文件的名称
		fileName := h.Filename
		arr := strings.Split(fileName, ":")
		if len(arr) > 1 {
			index := len(arr) - 1
			fileName = arr[index]
		}
		//关闭上传的文件，不然的话会出现临时文件不能清除的情况
		f.Close()
		//保存文件到指定的位置
		//static/img,这个是文件的地址,路径必须存在，第一个static前面不要有/
		this.SaveToFile("image", path.Join("static/img/", fileName))
		//显示在本页面，不做跳转操作
		this.TplName = "admin/user/index.html"
	}
}


|-- views
|     |--admin
|			|--user
|	  			  `-- add.html

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>admin/user/add</title>
</head>
<body>
    this is admin/user/add
    <form action='{{urlfor "UserController.Add"}}' method="post" enctype="multipart/form-data">
        <div class="field-content">
            文件上传：<input name="image" type="file" />
        </div>
        <div class="field-content">
            <input type="submit" value="提交" />
        </div>
    </form>
</body>
</html>



























