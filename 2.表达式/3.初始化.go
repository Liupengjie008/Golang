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

