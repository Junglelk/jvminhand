package main

import (
	"fmt"
	"jvmgo/jvmgo/ch02/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v \n", cp, cmd.class, cmd.args)
	// Replace函数将字符串中的old(此处为".")替换为new(此处为"/")，替换进行n次，当n小于0时，没有次数限制
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s \n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
