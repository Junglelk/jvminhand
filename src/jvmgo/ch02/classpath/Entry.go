package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathSeparator)

// Entry 接口定义语法
// 关于大小写，在Go 语言里，标识符要么从包里公开，要么不从包里公开。当代码导入了一个包时，程序可以直接访问这个包中任意一个公开的标识符。
// 这些标识符以大写字母开头。以小写字母开头的标识符是不公开的，不能被其他包中的代码直接访问。
// 但是，其他包可以间接访问不公开的标识符。
// 例如，一个函数可以返回一个未公开类型的值，那么这个函数的任何调用者，哪怕调用者不是在这个包里声明的，都可以访问这个值。
type Entry interface {
	// readClass 函数定义方式: functionName(paramName paramType) (returnValue01, returnValue02, ...)
	// go语言的参数可以拥有多个返参
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
