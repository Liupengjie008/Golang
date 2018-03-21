Go语言的主要特征：
	1.自动立即回收。
	2.更丰富的内置类型。
	3.函数多返回值。
	4.错误处理。
	5.匿名函数和闭包。
	6.类型和接口。
	7.并发编程。
	8.反射。
	9.语言交互性。

Golang文件名：所有的go源码都是以 .go 结尾。

Go语言命名：
    1.Go的函数、变量、常量、自定义类型、包(Package)的命名方式遵循以下规则：
            1）首字符可以是任意的Unicode字符或者下划线
            2）剩余字符可以是Unicode字符、下划线、数字
            3）字符长度不限
    2.go只有25个关键字
            break        default      func         interface    select
            case         defer        go           map          struct
            chan         else         goto         package      switch
            const        fallthrough  if           range        type
            continue     for          import       return       var
    3.Go还有37个保留字
            Constants:    true  false  iota  nil

            Types:    int  int8  int16  int32  int64  
                      uint  uint8  uint16  uint32  uint64  uintptr
                      float32  float64  complex128  complex64
                      bool  byte  rune  string  error

            Functions:   make  len  cap  new  append  copy  close  delete
                         complex  real  imag
                         panic  recover
    4.可见性
            1）声明在函数内部，是函数的本地值，类似private
            2）声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值，类似protect
            3）声明在函数外部且首字母大写是所有包可见的全局值,类似public

Go语言声明：
    有四种主要声明方式：var（声明变量）, const（声明常量）, type（声明类型） ,func（声明函数）。
    Go的程序是保存在多个.go文件中，文件的第一行就是package XXX声明，用来说明该文件属于哪个包(package)，package声明下来就是import声明，再下来是无关吮吸的类型，变量，常量，函数的声明。


Go项目构建及编译
一个GO工程中主要包含以下三个目录：
src：源代码文件
pkg：包文件
bin：相关bin文件

step1: 建立工程文件夹 goproject
step2: 在工程文件夹中建立src,pkg,bin文件夹
step3: 在GOPATH中添加projiect路径    例 e:/goproject
step4: 如工程中有自己的包examplepackage，那在src文件夹下建立以包名命名的文件夹 例 examplepackage
step5：在src文件架下编写主程序代码代码 goproject.go
step6：在examplepackage文件夹中编写 examplepackage.go 和 包测试文件 examplepackage_test.go
step7：编译调试包
           go build examplepackage
           go test examplepackage
           go install examplepackage
	       这时在pkg文件夹中可以发现会有一个相应的操作系统文件夹如windows_386z, 在这个文件夹中会有examplepackage文件夹，在该文件中有examplepackage.a文件
step8：编译主程序
           go build goproject
           成功后会生成goproject.exe文件

至此一个GO工程编辑成功。


go 编译问题

golang的编译使用命令 go build , go install;除非仅写一个main函数，否则还是准备好目录结构；
GOPATH=工程根目录；其下应创建src，pkg，bin目录，bin目录中用于生成可执行文件，pkg目录中用于生成.a文件；
golang中的import name，实际是到GOPATH中去寻找name.a, 使用时是该name.a的源码中生命的package 名字；这个在前面已经介绍过了。

注意点：
1. 系统编译时 go install abc_name时，系统会到GOPATH的src目录中寻找abc_name目录，然后编译其下的go文件；
2. 同一个目录中所有的go文件的package声明必须相同，所以main方法要单独放一个文件，否则在eclipse和liteide中都会报错；
    编译报错如下：（假设test目录中有个main.go 和mymath.go,其中main.go声明package为main，mymath.go声明packag 为test);
    $ go install test
    can't load package: package test: found packages main (main.go) and test (mymath.go) in /home/wanjm/go/src/test
    报错说 不能加载package test（这是命令行的参数），因为发现了两个package，分别时main.go 和 mymath.go;
3. 对于main方法，只能在bin目录下运行 go build path_tomain.go; 可以用-o参数指出输出文件名；
4. 可以添加参数 go build -gcflags "-N -l"  ****,可以更好的便于gdb；详细参见http://golang.org/doc/gdb
5. gdb全局变量主一点。 如有全局变量 a；则应写为 p 'main.a'；注意但引号不可少；