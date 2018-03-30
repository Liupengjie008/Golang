文件状态
从文件中写入写出内容
这一部分较多的涉及I/O的相关操作，系统的介绍放在I/O那部分来整理，大体上向文件中读写内容的时候有三种方式：

1、在使用f, err := os.Open(file_path)打开文件之后直接使用 f.read() f.write() 结合自定义的buffer每次从文件中读入/读出固定的内容

2、使用ioutl的readFile和writeFile方法

3、使用bufio采用带有缓存的方式进行读写，比如通过info:=bufio.NewReader(f)将实现了io.Reader的接口的实例加载上来之后，就可以使用info.ReadLine（）来每次实现一整行的读取，直到err信息为io.EOF时，读取结束

这个blog对三种文件操作的读入速度进行了比较，貌似读取大文件的时候采用ioutil的时候效率要高些。

每种方式都有不同的适用情况，下面是分别用三种方式进行读出操作的例子，对于写入文件的操作，可以参考读出操作来进行：

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//查看当前的工作目录路径 得到测试文件的绝对路径
	current_dir, _ := os.Getwd()
	fmt.Println(current_dir)
	file_path := current_dir + "/temp.txt"

	//方式一：
	//通过ioutil直接通过文件名来加载文件
	//一次将整个文件加载进来 粒度较大 err返回为nil的时候 文件会被成功加载
	dat, err := ioutil.ReadFile(file_path)
	//若加载的是一个目录 会返回[]os.FileInfo的信息
	//ioutil.ReadDir()
	check(err)
	//the type of data is []uint
	fmt.Println(dat)
	//将文件内容转化为string输出
	fmt.Println(string(dat))

	//方式二：
	//通过os.Open的方式得到 *File 类型的变量
	//貌似是一个指向这个文件的指针 通过这个指针 可以对文件进行更细粒度的操作
	f, err := os.Open(file_path)
	check(err)
	//手工指定固定大小的buffer 每次通过buffer来 进行对应的操作
	buffer1 := make([]byte, 5)
	//从文件f中读取len(buffer1)的信息到buffer1中 返回值n1是读取的byte的长度
	n1, err := f.Read(buffer1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(buffer1))

	//通过f.seek进行更精细的操作 第一个参数表示offset为6 第二个参数表示文件起始的相对位置
	//之后再读就从o2位置开始往后读信息了
	o2, err := f.Seek(6, 0)
	check(err)
	buffer2 := make([]byte, 2)
	//读入了n2长度的信息到buffer2中
	n2, err := f.Read(buffer2)
	check(err)
	fmt.Printf("%d bytes after %d position : %s\n", n2, o2, string(buffer2))

	//通过io包种的函数 也可以实现类似的功能
	o3, err := f.Seek(6, 0)
	check(err)
	buffer3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, buffer3, len(buffer3))
	check(err)
	fmt.Printf("%d bytes after %d position : %s\n", n3, o3, string(buffer3))

	//方式三
	//通过bufio包来进行读取 bufio中又许多比较有用的函数 比如一次读入一整行的内容

	//调整文件指针的起始位置到最开始的地方
	_, err = f.Seek(10, 0)
	check(err)
	r4 := bufio.NewReader(f)

	//读出从头开始的5个字节
	b4, err := r4.Peek(5)
	check(err)
	//fmt.Println(string(b4))
	fmt.Printf("5 bytes : %s\n", string(b4))

	//调整文件到另一个地方
	_, err = f.Seek(0, 0)
	check(err)
	r5 := bufio.NewReader(f)
	//读出从指针所指位置开始的5个字节
	b5, err := r5.Peek(5)
	check(err)
	//fmt.Println(string(b4))
	fmt.Printf("5 bytes : %s\n", string(b5))

	//测试bufio的其他函数

	for {
		//读出内容保存为string 每次读到以'\n'为标记的位置
		line, err := r5.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF {
			break
		}
	}
	//ReadLine() ReadByte() 的用法都是类似 一般都是当err为io.EOF的时候
	//读入内容就结束
	//感觉实际用的时候 还是通过方式三比较好 粒度正合适 还有多种处理输入的方式

	f.Close()

}