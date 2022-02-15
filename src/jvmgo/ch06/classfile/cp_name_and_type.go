package classfile

/*
	CONSTANT_NameAndType_info{
		u1 tag;
		u2 name_index;
		u2 descriptor_index;
	}
*/
// CONSTANT_NameAndType_info 给出字段或者方法的名称和描述符。CONSTANT_Class_info 和 CONSTANT_NameAndType_info 加在一起可以唯一确定一个字段或者方法；
// 字段或方法名由 name_index 给出，字段或方法的描述符由 descriptor_index 给出。name_index 和 descriptor_index 都是常量池索引，指向 CONSTANT_Utf8_info 常量。
// 字段和方法名就是代码中出现的（或者由编译器生成的）字段或者方法的名字。Java虚拟机规范定义了一种简单的语法来描述字段和方法，即描述符可依照以下规则生成：
// 1. 类型描述符
// 1.1 基本类型 byte,short,char,int,long,float,double的描述符是单个字母，分别对应 B,S,C,I,J,F,D	需要注意的是long的描述符为 J
// 1.2 引用类型，是 L + 类的完全限定名 + 分号，比如 Ljava.lang.Object 表示 java.lang.Object
// 1.3 数组类型，是 [+数组元素类型描述符，比如[[D 表示 double[][]; [I 表示 int[]
// 2. 字段描述符就是字段类型的描述符
// 3. 方法描述符是(分号分隔的参数类型描述符) + 返回值类型描述符，其中 void 返回值由字母 V 表示，比如 ()V 表示 void run(),(Ljava.lang.String;)V 表示void main(String[] args)
// 一个很简单的原理是，如果描述符不同，则表示的对象就不是同一个，这就解释了java为什么支持方法的重载，因为参数列表不同时，描述符就不同，指向的就不是同一个对象。
// 所以从虚拟机的角度说，Java也可以支持多个类型不一致的同名字段，即“字段”重写。但java语法不支持就是了。

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (e *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	e.nameIndex = reader.readUint16()
	e.descriptorIndex = reader.readUint16()
}
