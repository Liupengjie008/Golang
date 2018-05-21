首先 安装Go语言环境 + Git工具：

beego的框架环境搭建:

1）安装beego：
go get github.com/astaxie/beego

2）安装bee工具（框架生成工具）
go get github.com/beego/bee

3）使用bee工具生成 项目 代码
注意：先进入到GOPATH的src路径下，再输入" bee new 项目名"

查看 GOPATH：
go env
GOPATH="/Users/***/Desktop/go"

GOPATH的src路径下:

cd /Users/***/Desktop/go/src/

输入" bee new 项目名"：

bee new project
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v1.9.0
2018/04/13 10:20:24 INFO     ▶ 0001 Creating application...
	create	 /Users/liupengjie/Desktop/go/src/project/
	create	 /Users/liupengjie/Desktop/go/src/project/conf/
	create	 /Users/liupengjie/Desktop/go/src/project/controllers/
	create	 /Users/liupengjie/Desktop/go/src/project/models/
	create	 /Users/liupengjie/Desktop/go/src/project/routers/
	create	 /Users/liupengjie/Desktop/go/src/project/tests/
	create	 /Users/liupengjie/Desktop/go/src/project/static/
	create	 /Users/liupengjie/Desktop/go/src/project/static/js/
	create	 /Users/liupengjie/Desktop/go/src/project/static/css/
	create	 /Users/liupengjie/Desktop/go/src/project/static/img/
	create	 /Users/liupengjie/Desktop/go/src/project/views/
	create	 /Users/liupengjie/Desktop/go/src/project/conf/app.conf
	create	 /Users/liupengjie/Desktop/go/src/project/controllers/default.go
	create	 /Users/liupengjie/Desktop/go/src/project/views/index.tpl
	create	 /Users/liupengjie/Desktop/go/src/project/routers/router.go
	create	 /Users/liupengjie/Desktop/go/src/project/tests/default_test.go
	create	 /Users/liupengjie/Desktop/go/src/project/main.go
2018/04/13 10:20:24 SUCCESS  ▶ 0002 New application successfully created!

一个Beego框架的项目就生成成功！

目录结构如下所示:

project
|-- conf       --配置文件
|   `-- app.conf
|-- controllers    --控制器
|   `-- default.go
|-- main.go
|-- models    --模型
|-- routers
|   `-- router.go
|-- static        --静态文件
|   |-- css
|   |-- img
|   `-- js
|-- tests
|   `-- default_test.go
`-- views       --模版页面
    `-- index.tpl


4）运行项目：(bee run ：代码热编译功能)
cd /Users/***/Desktop/go/src/project
bee run 
______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v1.9.0
2018/04/13 10:26:55 INFO     ▶ 0001 Using 'project' as 'appname'
2018/04/13 10:26:55 INFO     ▶ 0002 Initializing watcher...
2018/04/13 10:26:56 SUCCESS  ▶ 0003 Built Successfully!
2018/04/13 10:26:56 INFO     ▶ 0004 Restarting 'project'...
2018/04/13 10:26:56 SUCCESS  ▶ 0005 './project' is running...
2018/04/13 10:26:56 [I] [asm_amd64.s:2337] http server Running on http://:8080

5）打开浏览器，输入 "http://127.0.0.1:8080" 查看运行的结果。
出现 welcome to beego 界面
(如果你本地装有其他应用占用了8080端口，可以更换下端口)


6）为什么没有 nginx 和 apache 居然可以自己干这个事情？
   Go 已经做了网络层的东西，beego 只是封装了一下，所以可以做到不需要 nginx 和 apache。
