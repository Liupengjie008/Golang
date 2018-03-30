文件删除
文件删除的时候，不管是普通文件还是目录文件，都可以用err:=os.Remove(filename)这样的操作来执行。当然要是想移除整个文件夹，直接使用RemoveAll(path string)操作即可。可以看一下RemoveAll函数的内部实现，整体上就是遍历，递归的操作过程，其他的类似的文件操作都可以用类似的模板来实现，下面以RemoveAll函数为模板，进行一下具体的分析，注意考虑到各种情况：

func RemoveAll(path string) error {
// Simple case: if Remove works, we're done.
//先尝试一下remove如果是普通文件 直接删掉 报错 则可能是目录中还有子文件
err := Remove(path)
//没错或者路径不存在 直接返回 nil
if err == nil || IsNotExist(err) {
	return nil
}

// Otherwise, is this a directory we need to recurse into?
// 目录里面还有文件 需要递归处理
// 注意Lstat和stat函数的区别，两个都是返回文件的状态信息
//Lstat多了处理Link文件的功能，会返回Linked文件的信息，而state直接返回的是Link文件所指向的文件的信息
dir, serr := Lstat(path)
if serr != nil {
	if serr, ok := serr.(*PathError); ok && (IsNotExist(serr.Err) || serr.Err == syscall.ENOTDIR) {
		return nil
	}
	return serr
}
//不是目录
if !dir.IsDir() {
	// Not a directory; return the error from Remove.
	return err
}

// Directory.
fd, err := Open(path)
if err != nil {
	if IsNotExist(err) {
		// Race. It was deleted between the Lstat and Open.
		// Return nil per RemoveAll's docs.
		return nil
	}
	return err
}

// Remove contents & return first error.
err = nil
//递归遍历目录中的文件 如果参数n<=0则将全部的信息存入到一个slice中返回
//如果参数n>0则至多返回n个元素的信息存入到slice当中
//还有一个类似的函数是Readdir 这个返回的是 目录中的内容的Fileinfo信息

for {
	names, err1 := fd.Readdirnames(100)
	for _, name := range names {
		err1 := RemoveAll(path + string(PathSeparator) + name)
		if err == nil {
			err = err1
		}
	}
	//遍历到最后一个位置
	if err1 == io.EOF {
		break
	}
	// If Readdirnames returned an error, use it.
	if err == nil {
		err = err1
	}
	if len(names) == 0 {
		break
	}
}

// Close directory, because windows won't remove opened directory.
fd.Close()
//递归结束 当前目录下位空 删除当前目录
// Remove directory.
err1 := Remove(path)
if err1 == nil || IsNotExist(err1) {
	return nil
}
if err == nil {
	err = err1
}
return err
}    