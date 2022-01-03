package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry golang中声明数组的语法与Java不一致，是先有中括号后跟类型
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	// 定义一个变量
	var compositeEntry []Entry
	// strings 和Java中的Arrays方法类很像的样子，提供一些针对字符串的操作
	// pathListSeparator 在main包中定义的一个常量，使用系统文件分隔符
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (c CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		data, from, err := entry.readClass(className)
		if err != nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (c CompositeEntry) String() string {
	// make 函数别调用来初始化map、slice和chan
	strs := make([]string, len(c))
	for i, entry := range c {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
