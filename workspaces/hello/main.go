// Package main main 文件，go 多模块工作区演示代码
// 实现将传入的字符串反转
package main

import (
	"flag"
	"fmt"
)

var (
	str = ""
)

func init() {
	flag.StringVar(&str, "str", str, "输入字符")
	flag.Parse()
}

func main() {
	if str == "" {
		fmt.Println("示例: go run main.go -str hello")
		fmt.Println("str 参数必填")
		flag.Usage()
		return
	}
	// TODO:: 反转字符串以后再输出
	fmt.Println(str)
	return
}
