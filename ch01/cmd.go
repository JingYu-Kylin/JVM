package main

import (
	"flag"
	"fmt"
	"os"
)

/**
 * Cmd结构体
 */
type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args []string
}

//*是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值
func parseCmd() *Cmd {
	// & 是取地址符号
	cmd := &Cmd{}
	/**
	 * 首先设置flag.Usage变量，把printUsage（）函数赋值给它
	 */
	flag.Usage = printUsage
	/**
	 * 然后调用flag包提供的各种Var（）函数设置需要解析的选项
	 */
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()
	/**
	 * 接着调用Parse（）函数解析选项。
	 */
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

/**
 * 如果Parse（）函数解析失败，它就调用printUsage（）函数把命令的用法打印到控制台
 */
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}
