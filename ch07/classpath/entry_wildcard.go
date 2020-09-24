package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// WildcardEntry实际上也是CompositeEntry，所以就不再定义新的类型

func newWildcardEntry(path string) CompositeEntry {
	// 首先把路径末尾的星号去掉，得到baseDir
	baseDir := path[:len(path)-1] // remove *
	compositeEntry := []Entry{}

	// 在walkFn中，根据后缀名选出JAR文件，并且返回SkipDir跳过子目录（通配符类路径不能递归匹配子目录下的JAR文件）。
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	// 然后调用filepath包的Walk（）函数遍历baseDir创建ZipEntry
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}


