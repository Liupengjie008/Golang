Windows安装Golang环境

对于Windows用户，Go项目提供两种安装选项（从源码安装除外）： 
1、zip压缩包需要你设置一些环境变量
2、MSI安装程序则会自动配置你的安装  	##实验性


ZIP安装程序

1. 根据操作系统选择相应Golang版本( Golang资源站：http://www.golang.mom/ )
	下载：go1.9.1.windows-386.zip

2. 将下载后的文件解压，放到了D:\go目录下 

3. 设置环境变量

$GOROOT  指向golang安装之后的根目录，如果你选择.msi安装的话，windows平台下默认为c:/go，而且在安装过程中由安装程序自动写入系统环境变量。
$GOARCH  目标平台（编译后的目标平台）的处理器架构（386、amd64、arm）
$GOOS      目标平台（编译后的目标平台）的操作系统（darwin、freebsd、linux、windows）
$GOBIN    指向安装之后根目录下的bin目录，即$GOROOT/bin，windows平台下默认为c:/go/bin，会在安装过程中由安装程序自动添加到PATH变量中

对于我们来说只需要配置GOROOT和GOBIN即可(如果你下载的是msi安装后,就会自动写入环境变量,而无需我们配置)
右键我的电脑-->属性-->高级系统设置-->环境变量

设置GOROOT
右键我的电脑-->属性-->高级系统设置-->环境变量-->系统变量-->新建
变量名(N)：GOROOT
变量值(V)：D:\go

将bin路径添加到Path目录中
右键我的电脑-->属性-->高级系统设置-->环境变量-->系统变量-->找到Path双击
变量值(V)：(追加) %GOROOT%\bin;

设置GOPATH目录

go 命令依赖一个重要的环境变量：$GOPATH
注：
a.  这个不是Go安装目录，相当于我们的工作目录，在类似 Unix 环境这样设置： export    GOPATH=/home/apple/mygo
b. GOPATH允许多个目录，当有多个目录时请注意分隔符，多个目录的时候Windows是(;)分号，Linux系统是(:)冒号，当有多个GOPATH时，默认会将go get的内容放在第一个目录下
c.  $GOPATH 目录约定有三个子目录：
 
src 存放源代码（比如：.go .c .h .s等）
pkg 编译后生成的文件（比如：.a）
bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）

我们在D盘新建一个目录GoProject作为我们的gopath目录，并在目录中创建一个文件夹src，用来存放我们的源文件。

将GOPATH加入到环境变量中去：
右键我的电脑-->属性-->高级系统设置-->环境变量-->系统变量-->新建
变量名(N)：GOPATH
变量值(V)：D:\GoProject

打开终端运行 go version命令	( win+r 输入cmd ，输入go version )
如果出现如下图所示，说明安装成功
C:\Users\Administrator>go version
go version go1.9.1 windows/amd64

查看Go语言的环境信息
C:\Users\Administrator>go env


MSI安装程序

1. 根据操作系统选择相应Golang版本( Golang资源站：http://www.golang.mom/ )
	下载：go1.9.1.windows-386.msi
	
2. 打开此MSI文件 并跟随提示来安装Go工具。默认情况下，该安装程序会将Go发行版放到 c:\Go 中。
  此安装程序应该会将 c:\Go\bin 目录放到你的 PATH 环境变量中。 要使此更改生效，你需要重启所有打开的命令行。

3. 检查环境变量，输出版本则安装成功
win+r 输入cmd ，输入go version，
C:\Users\Administrator>go version
go version go1.9.1 windows/amd64
  
4. 在Windows下设置环境变量
  在Windows下，你可以通过在系统“控制面板”中，“高级”标签上的“环境变量”按钮来设置环境变量。 Windows的一些版本通过系统“控制面板”中的“高级系统设置”选项提供此控制板。