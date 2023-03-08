package main

import (
	"fmt"
	"github.com/MeguruForever/calculator/src"
)
func main(){
	fmt.Println("输入算式")
	var str string
	fmt.Scan(&str)
	str = src.Match(str)
	fmt.Println(str)
	fmt.Println(src.Calculate(str))
	fmt.Println("输入任意键退出")
	var s string
	fmt.Scan(&s)
}