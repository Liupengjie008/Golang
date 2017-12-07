linux安装Golang环境

1. SSH远程登录你的linux服务器

2. [root@localhost ~]# yum install mercurial 		##安装 mercurial包

3. [root@localhost ~]# yum install git 		##安装git包

4. [root@localhost ~]# yum install gcc 		##安装gcc

5. 下载golang的压缩包:(可选择最新的Golang版本)

[root@localhost ~]# cd /usr/local/

[root@localhost local]# wget https://go.googlecode.com/files/go1.9.1.linux-amd64.tar.gz

注意：如果不能翻墙，去go语言资源站  http://www.golang.mom/ 下载相应的包。然后通过ftp上传到此目录。

6. 下载完成 or ftp上传完成，用tar 命令来解压压缩包。
[root@localhost local]# tar -zxvf go1.9.1.linux-amd64.tar.gz

7. 建立Go的工作空间（workspace，也就是GOPATH环境变量指向的目录）
GO代码必须在工作空间内。工作空间是一个目录，其中包含三个子目录：
src ---- 里面每一个子目录，就是一个包。包内是Go的源码文件
pkg ---- 编译后生成的，包的目标文件
bin ---- 生成的可执行文件
这里，我们在/home目录下, 建立一个名为go(可以不是go, 任意名字都可以)的文件夹，
然后再建立三个子文件夹(子文件夹名必须为src、pkg、bin)。

[root@localhost local]# cd /home/
[root@localhost home]# mkdir go
[root@localhost home]# cd go/
[root@localhost go]# mkdir bin
[root@localhost go]# mkdir src
[root@localhost go]# mkdir pkg

8. 添加PATH环境变量and设置GOPATH环境变量

[root@localhost go]# vi /etc/profile 

加入下面这三行:
export GOROOT=/usr/local/go		##Golang安装目录
export PATH=$GOROOT/bin:$PATH
export GOPATH=/home/go  ##Golang项目目录

保存后，执行以下命令，使环境变量立即生效:

[root@localhost go]# source /etc/profile		##刷新环境变量

至此，Go语言的环境已经安装完毕。

9. 验证一下是否安装成功，如果出现下面的信息说明安装成功了
[root@localhost go]# go version		##查看go版本
go version go1.9.1 linux/amd64

10. 查看Go语言的环境信息
[root@localhost go]# go env

编译运行一个简单的程序：

1. [root@localhost ~]# cd /home/go/src/		##进入项目目录

2. 创建项目目录及文件：( project/test/main/main.go )
[root@localhost src]# mkdir project
[root@localhost src]# cd project/
[root@localhost project]# mkdir test
[root@localhost project]# cd test/
[root@localhost test]# mkdir main
[root@localhost test]# cd main/
[root@localhost main]# touch main.go

3. 进入main.go编写第一个程序
[root@localhost main]# vi main.go 

package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}

4. 使用go install工具编译此程序
[root@localhost main]# cd /home/go/
[root@localhost go]# go install project/test/main
5. 进入bin目标查看编译好的二进制文件 ( cd /home/go/bin )
[root@localhost go]# cd bin/
[root@localhost bin]# ls
main

6. 使用shell命令 ./ 运行编译好的二进制文件
[root@localhost bin]# ./main 
Hello World



