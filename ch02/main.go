package main

import (
	"JVM-GO/ch02/classpath"
	"fmt"
	"strings"
)

// F:\GoSpace\bin\ch02.exe -Xjre "D:\Applications\Java\jdk1.8.0_221\jre" java.lang.Object
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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
