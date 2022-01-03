package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// 通配符Entry，即导包最后的那种星号的形式
func newWildcardEntry(path string) CompositeEntry {
	// 去除 *
	baseDir := path[:len(path)-1]
	var compositeEntry []Entry

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			//
			jarEntry := newZipEntry(path)
			// g
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	// Walk 会遍历以root为根的文件树，并对每一个文件和目录执行函数
	// walkFn就是需要执行的函数
	err := filepath.Walk(baseDir, walkFn)
	if err != nil {
		return nil
	}
	return compositeEntry
}
