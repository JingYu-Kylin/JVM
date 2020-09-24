package main

import (
	"JVM-GO/ch07/classpath"
	"JVM-GO/ch07/rtda/heap"
	"fmt"
	"strings"
)
// rtda是run-time data area的首字母缩写

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 1.8.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	// 先创建ClassLoader实例，然后用它来加载主类
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	// 最后执行主类的main（）方法。
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		// 解释执行main（）方法
		interpret(mainMethod, cmd.verboseInstFlag)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}



