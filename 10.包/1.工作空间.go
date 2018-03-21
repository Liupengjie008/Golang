Golang 工作空间 ：编译工具对源码目录有严格要求，每个工作空间 (workspace) 必须由 bin、pkg、src 三个目录组成。

workspace
    |
	+--- bin     			// go install 安装目录。
	|	|
	| 	+--- learn
	|
	+--- pkg。             // go build 生成静态库 (.a) 存放目录。
	|	  |
	| 	  +--- darwin_amd64 
	|				|
	| 				+--- mylib.a 
	|				|
	| 				+--- mylib 
	|					   |
	| 					   +--- sublib.a
	|	
	+--- src 			// 项目源码目录。
	|
	+--- learn
	|	   |
	| 	   +--- main.go
	|  
	+--- mylib
		   |		
		   +--- mylib.go
		   |		
		   +--- sublib
				   |
				   +--- sublib.go


可在 GOPATH 环境变量列表中添加多个工作空间，但不能和 GOROOT 相同。 

export GOPATH=$HOME/projects/golib:$HOME/projects/go

通常 go get 使用第一个工作空间保存下载的第三方库。











