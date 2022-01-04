package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	// 启动类路径 使用-Xjre解析启动类路径和扩展类路径
	bootClasspath Entry
	// 扩展类路径
	extClasspath Entry
	// 用户类路径 使用-classpath/-cp解析用户类路径
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (e *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	e.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	e.extClasspath = newWildcardEntry(jreExtPath)

}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	// Stat returns a FileInfo describing the named file.
	// If there is an error, it will be of type *PathError.
	// if中含有分号的写法是指声明了一个局部变量，分号之后是针对该局部变量的布尔表达式
	// 可以在外部定义一个err变量，全局变量与该处的局部变量不会相互影响；当然如果该处使用的是“=”而非“:=”，就只是单纯的赋值语句，会相互影响
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// parseUserClasspath 如果用户没有提供-classpath/-cp选项，使用当前目录作为用户类路径
func (e *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	e.userClasspath = newEntry(cpOption)
}

// ReadClass 方法依次从启动类路径，扩展类路径和用户路径中搜索class文件
func (e *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := e.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := e.extClasspath.readClass(className); err != nil {
		return data, entry, err
	}
	return e.userClasspath.readClass(className)
}

func (e *Classpath) String() string {
	return e.userClasspath.String()
}
