package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	// 用于存放目录的绝对路径
	absDir string
}
/**
 * Go没有专门的构造函数，
 * 本书统一使用new开头的函数来创建结构体实例，并把这类函数称为构造函数
 */
func newDirEntry(path string) *DirEntry {
	// 先把参数转换成绝对路径
	absDir, err := filepath.Abs(path)
	// 如果转换过程出现错误，则调用panic（）函数终止程序执行
	if err != nil {
		panic(err)
	}
	// 否则创建DirEntry实例并返回
	return &DirEntry{absDir}
}
func (self *DirEntry) readClass(className string) ([]byte , Entry, error) {
	// 先把目录和class文件名拼成一个完整的路径
	fileName := filepath.Join(self.absDir, className)
	// 然后调用ioutil包提供的ReadFile（）函数读取class文件内容
	data, err := ioutil.ReadFile(fileName)
	// 最后返回
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}