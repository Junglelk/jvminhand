package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func (e *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(e.absPath)
	if err != nil {
		return nil, nil, err
	}

	// defer 关键字会让随后的代码在函数返回后执行，这样可以保证这个函数一定会被调用
	defer r.Close()

	// 此处是遍历zip压缩包内的文件
	// 使用for range 关键字来迭代文件，range关键字可以迭代迭代数组、字符串、切片、映射和通道
	// 迭代数组时，会返回下标值和对应对象，由于此处我们并不需要下标值，且go中不允许存在未使用变量，所以使用下划线_来代替掉
	// 此外如果在导包语句中使用下划线打头，则仅执行该包内的init()方法
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
			return data, e, nil
		}
	}
	return nil, nil, errors.New("class not found" + className)
}

func (e *ZipEntry) String() string {
	return e.absPath
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	// return 语句用absPath新建了一个ZipEntry实例，实际上花括号中的内容类似于Java里的构造函数的参数，形制上又类似于C语言的数组，所以多个需要用逗号隔开
	// 不加逗号GoLand 会提示 “Need a trailing comma before a newline in the composite literal” 即为：在复合文字的换行符中间需要尾随一个逗号
	return &ZipEntry{absPath}
}
