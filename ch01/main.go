package main

import "fmt"

// ch01.exe -help
// ch01.exe -version
func main() {
	/**
	 * 先调用ParseCommand（）函数解析命令行参数
	 */
	cmd := parseCmd()
	if cmd.versionFlag {
		/**
		 * 。如果用户输入了-version选项，则输出版本信息
		 */
		fmt.Println("version 1.8.0")
	} else if cmd.helpFlag || cmd.class == "" {
		/**
		 * 如果解析出现错误，或者用户输入了-help选项，则调用PrintUsage（）函数打印出帮助信息
		 */
		printUsage()
	} else {
		/**
		 * 如果一切正常，则调用startJVM（）函数启动Java虚拟机
		 * 因为我们还没有真正开始编写Java虚拟机，所以startJVM（）函数暂时只是打印一些信息而已，
		 */
		startJVM(cmd)
	}
}
