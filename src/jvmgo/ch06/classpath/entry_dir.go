package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

// newDirEntry 先把参数转换为绝对路径，如果转换过程中出现错误，则调用panic终止执行，否则返回DirEntry实例
// 可以称这类函数为构造函数，但感觉应该不是正式的名字
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

// func (self pointer) funcName(param01 param01Type,...) (returnType01,...)
// readClass 这是结构体实现方法的写法，如果该结构体实现了某个接口的所有方法，则说该结构体实现了该方法
func (e *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(e.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, e, err
}

// String 只有结构体DirEntry同时实现了String()和readClass()方法，才说DirEntry实现了接口Entry
func (e *DirEntry) String() string {
	return e.absDir
}
