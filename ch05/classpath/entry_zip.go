package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	// 存放ZIP或JAR文件的绝对路径
	absPath string
}
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 首先打开ZIP文件
	r, err := zip.OpenReader(self.absPath)
	// 如果这一步出错的话，直接返回
	if err != nil {
		return nil, nil, err
	}
	// 有两处使用了defer语句来确保打开的文件得以关闭
	defer r.Close()
	// 然后遍历ZIP压缩包里的文件，看能否找到class文件
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			// 如果能找到，则打开class文件，把内容读取出来，并返回
			return data, self, nil
		}
	}
	// 如果找不到，或者出现其他错误，则返回错误信息
	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
