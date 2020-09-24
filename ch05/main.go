package main

import (
	"JVM-GO/ch05/classfile"
	"JVM-GO/ch05/classpath"
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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		// 解释执行main（）方法
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}

}

/**
读取并解析class文件
 */
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}

/**
查找类的main（）方法
 */
func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}



