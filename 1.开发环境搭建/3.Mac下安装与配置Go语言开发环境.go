Mac下安装与配置Go语言开发环境
1. 下载安装包安装 ( go1.9.1.darwin-amd64.pkg )

　　官网 (需翻墙)：https://storage.googleapis.com/golang/
	
	Golang资源站：http://www.golang.mom/ 

2. 配置Go环境变量GOPATH和GOBIN

　　（1）打开终端，cd ~

　　（2）查看是否有.bash_profile文件：

　　　　  ls -all

　　（3）有则跳过此步，没有则：

　　　　1）创建：touch .bash_profile

　　　　2）编辑：open -e .bash_profile

　　　　3）自定义GOPATH和GOBIN位置：

			export GOPATH=/Users/hopkings/www/Go
			export GOBIN=$GOPATH/bin
			export PATH=$PATH:$GOBIN

　　（4）编译：source .bash_profile

3. 查看Go环境变量：go env