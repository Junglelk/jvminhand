package heap

// SymRef 前面提到的常量池中的引用类型符号有些许共性，所以抽出一个结构体以减少重复代码
// cp字段存放符号引用所在的运行时常量池指针，这样可以通过符号引用访问到运行时常量池，
// 进一步又可以访问到类数据（运行时常量池中有类数据）。className字段存放类的完全限定名。
// class字段缓存解析后的类结构体指针。
// 对类符号引用，只要有类名就可以解析符号引用，对于字段，首先要解析类符号引用得到类数据，然后用字段名和描述符查找字段数据
// 方法符号引用的解析和字段符号引用类似
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (e *SymRef) ResolvedClass() *Class {
	if e.class == nil {
		e.resolvedClassRef()
	}
	return e.class
}

func (e *SymRef) resolvedClassRef() {
	d := e.cp.class
	c := d.loader.LoadClass(e.className)
	if c.IsAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	e.class = c
}
