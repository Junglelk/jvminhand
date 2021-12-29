package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathSeparator)

// Entry 接口定义语法
type Entry interface {
	// readClass 函数定义方式: functionName(paramName paramType) (returnValue01, returnValue02, ...)
	readClass(className string) ([]byte, Entry, error)
	// String 由上可知，此方法返回一个string，无参数
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
