文件打包，文件解压，文件遍历，这些相关的操作基本上都可以参考RemoveAll的方式来进行，就是递归加遍历的方式。
下面是文件压缩的一个实现：

//将文件夹中的内容打包成 .gz.tar 文件
package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

//将fi文件的内容 写入到 dir 目录之下 压缩到tar文件之中
func Filecompress(tw *tar.Writer, dir string, fi os.FileInfo) {

	//打开文件 open当中是 目录名称/文件名称 构成的组合
	filename := dir + "/" + fi.Name()
	fmt.Println("the last one:", filename)
	fr, err := os.Open(filename)
	fmt.Println(fr.Name())
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	hdr, err := tar.FileInfoHeader(fi, "")

	hdr.Name = fr.Name()
	if err = tw.WriteHeader(hdr); err != nil {
		panic(err)
	}
	//bad way
	//	//信息头部 生成tar文件的时候要先写入tar结构体
	//	h := new(tar.Header)
	//	//fmt.Println(reflect.TypeOf(h))

	//	h.Name = fi.Name()
	//	h.Size = fi.Size()
	//	h.Mode = int64(fi.Mode())
	//	h.ModTime = fi.ModTime()

	//	//将信息头部的内容写入
	//	err = tw.WriteHeader(h)
	//	if err != nil {
	//		panic(err)
	//	}

	//copy(dst Writer,src Reader)
	_, err = io.Copy(tw, fr)
	if err != nil {
		panic(err)
	}
	//打印文件名称
	fmt.Println("add the file: " + fi.Name())

}

//将目录中的内容递归遍历 写入tar 文件中
func Dircompress(tw *tar.Writer, dir string) {
	fmt.Println(dir)
	//打开文件夹
	dirhandle, err := os.Open(dir + "/")
	//fmt.Println(dir.Name())
	//fmt.Println(reflect.TypeOf(dir))
	if err != nil {
		panic(err)
	}
	defer dirhandle.Close()

	fis, err := dirhandle.Readdir(0)
	//fis的类型为 []os.FileInfo

	//也可以通过Readdirnames来读入所有子文件的名称
	//但是这样 再次判断是否为文件的时候 需要通过Stat来得到文件的信息
	//返回的就是os.File的类型

	if err != nil {
		panic(err)
	}

	//遍历文件列表 每一个文件到要写入一个新的*tar.Header
	//var fi os.FileInfo
	for _, fi := range fis {
		fmt.Println(fi.Name())

		if fi.IsDir() {

			newname := dir + "/" + fi.Name()
			fmt.Println("using dir")
			fmt.Println(newname)
			//这个样直接continue就将所有文件写入到了一起 没有层级结构了
			//Filecompress(tw, dir, fi)
			Dircompress(tw, newname)

		} else {
			//如果是普通文件 直接写入 dir 后面已经有了 /
			Filecompress(tw, dir, fi)
		}

	}

}

//在tardir目录中创建一个.tar.gz文件 存放压缩之后的文件
func Dirtotar(sourcedir string, tardir string, tarname string) {
	//file write 在tardir目录下创建
	fw, err := os.Create(tardir + "/" + tarname + ".tar.gz")
	//type of fw is *os.File
	//	fmt.Println(reflect.TypeOf(fw))
	if err != nil {
		panic(err)

	}
	defer fw.Close()

	//gzip writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	//tar write
	tw := tar.NewWriter(gw)

	fmt.Println("源目录：", sourcedir)
	Dircompress(tw, sourcedir)

	//通过控制写入流 也可以控制 目录结构 比如将当前目录下的Dockerfile文件单独写在最外层
	fileinfo, err := os.Stat("tarrepo" + "/" + "testDockerfile")
	fmt.Println("the file name:", fileinfo.Name())
	if err != nil {
		panic(err)

	}
	//比如这里将Dockerfile放在 tar包中的最外层 会注册到tar包中的 /tarrepo/testDockerfile 中
	Filecompress(tw, "tarrepo", fileinfo)
	//Filecompress(tw, "systempdir/test_testwar_tar/", fileinfo)

	fmt.Println("tar.gz packaging OK")

}

func main() {
	//	workdir, _ := os.Getwd()
	//	fmt.Println(workdir)
	Dirtotar("testdir", "tarrepo", "testtar")

}