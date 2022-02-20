package heap

import (
	"fmt"
	"jvmgo/jvmgo/ch06/classfile"
	"jvmgo/jvmgo/ch06/classpath"
)

// ClassLoader 依赖Classpath来搜索和读取class文件，cp字段保存Classpath指针。
// classMap字段保存已加载的类数据，key是类的完全限定名。
// 可以将classMap字段视为方法区的具体实现
type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (e *ClassLoader) LoadClass(name string) *Class {
	if class, ok := e.classMap[name]; ok {
		return class
	}
	return e.loadNonArrayClass(name)
}

// loadNonArrayClass 加载类方法。
// 1. 找到class文件并把数据读取到内存；
// 2. 解析class文件，生成虚拟机可用的类数据，并放入方法区
// 3. 链接：分为验证和准备两个必要阶段。
func (e *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := e.readClass(name)
	class := e.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s form %s]", name, entry)
	return class
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

// prepare 给类变量分配空间并给予初始值
// 比如在类定义中，布尔型的定义初始值为false
func prepare(class *Class) {

}

// verify 类文件的验证极其繁琐，故目前假设加载到的类文件均为可信赖的编译器编译出的合理文件
func verify(class *Class) {

}

func (e *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := e.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

//
func (e *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = e
	resolverSuperClass(class)
	resolverInterfaces(class)
	e.classMap[class.name] = class
	return class
}

func resolverInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func resolverSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}
