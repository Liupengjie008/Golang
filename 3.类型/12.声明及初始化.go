声明及初始化

Go语言声明：
    有四种主要声明方式：var（声明变量）, const（声明常量）, type（声明类型） ,func（声明函数）。
    Go的程序是保存在多个.go文件中，文件的第一行就是package XXX声明，用来说明该文件属于哪个包(package)，package声明下来就是import声明，再下来是无关吮吸的类型，变量，常量，函数的声明。

初始化复合对象，必须使用类型标签，且左大括号必须在类型尾部。


//varastruct{xint}={100} 		//syntax error

// var b []int = { 1, 2, 3 } 	//syntax error

// c := struct {x int; y string} 	// syntax error: unexpected semicolon or newline
// {
// }


var a = struct{ x int }{100}
var b = []int{1, 2, 3}

初始化值以 "," 分隔。可以分多行，但最后一行必须以 "," 或 "}" 结尾。 

a := []int{		
	1,
	2 		// Error: need trailing comma before newline in composite literal
}

a := []int{ 	
	1,
	2, 		// ok
}

b := []int{ 
	1,
	2 } 	// ok

