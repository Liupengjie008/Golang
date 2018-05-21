beego admin:一个简单的基于beego的后台管理系统


安装

go get github.com/beego/admin


创建应用

1  首先,使用bee工具创建一个应用程序


bee new admin

2  创建成功以后，你能得到一个名叫admin的应用程序， 现在开始可以使用它了。找到到刚刚新建的程序admin/routers/router.go这个文件


import (
    "admin/controllers"         //自身业务包
    "github.com/astaxie/beego"  //beego 包
    "github.com/beego/admin"  //admin 包
)


3  引入admin代码，再init函数中使用它



func init() {
    admin.Run()
    beego.Router("/", &controllers.MainController{})
}


4 配置文件，这里是mysql的配置文件举例


db_host = localhost
db_port = 3306
db_user = root
db_pass = root
db_name = admin
db_type = mysql

/*
   
    数据库目前仅支持mysql,postgresql,sqlite3,后续会添加更多的数据库支持。

    数据库的配置信息需要填写，程序会根据配置自动建库 mysql数据库链接信息

    db_host = localhost
    db_port = 3306
    db_user = root
    db_pass = root
    db_name = admin
    db_type = mysql

    postgresql数据库链接信息

    db_host = localhost
    db_port = 5432
    db_user = postgres
    db_pass = postgres
    db_name = admin
    db_type = postgres
    db_sslmode=disable

    sqlite3数据库链接信息

    ###db_path 是指数据库保存的路径，默认是在项目的根目录
    db_path = ./
    db_name = admin
    db_type = sqlite3

    把以上信息配置成你自己数据库的信息。

*/

 

 

5 权限系统的配置



sessionon = true
rbac_role_table = role
rbac_node_table = node
rbac_group_table = group
rbac_user_table = user
#admin用户名 此用户登录不用认证
rbac_admin_user = admin

#默认不需要认证模块
not_auth_package = public,static
#默认认证类型 0 不认证 1 登录认证 2 实时认证
user_auth_type = 1
#默认登录网关
rbac_auth_gateway = /public/login
#默认模版
template_type=easyui

6  复制静态文件



cd $GOPATH/src/admin
cp -R ../github.com/beego/admin/static ./
cp -R ../github.com/beego/admin/views ./

目录结构如下所示:
admin
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
|   |-- easyui
|   |       |-- jquery-easyui   
|   |-- img
|   |-- js
|       `-- reload.min.js
|-- tests
|   `-- default_test.go
`-- views    
    |-- easyui
    |       |-- public   
    |       |-- rbac
    |       
    `-- index.tpl


编译运行：

 go build


./admin -syncdb

好了，现在可以通过浏览器地址访问了http://localhost:8080/

默认得用户名密码都是admin

github网址https://github.com/beego/admin










基于 beego admin 开发




 














