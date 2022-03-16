package basics

import (
	"bytes"
	"fmt"
)

func SwitchTypeDemo() {
	var x interface{}
	var y = 10.1
	x = y
	switch i := x.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("未知类型")
		fmt.Printf("类型为：%T", i)
	}
}

func SwitchYearDays() {
	fmt.Println("请输入年份和月份：")
	var year, month int
	fmt.Scanln(&year)
	fmt.Scanln(&month)
	var days int
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if  (year%4 == 0 && year%100 != 0) ||year%400 == 0  {
			days = 29
		} else {
			days = 28
		}
	default:
		days = 0
	}
	fmt.Printf("%d年%d月的天数为%d天", year, month, days)
}

// 定义一个函数, 参数数量为0~n, 类型约束为字符串
func joinString(aList ...string)string{
	// 定义一个字节缓冲, 快速地连接字符串
	var b bytes.Buffer
	// 遍历可变参数列表aList, 类型为[]string
	for _, s := range aList {
		// 将遍历出的字符串连续写入字节数组
		b.WriteString(s)
	}
	// 将连接好的字节数组转换为字符串并输出
	return b.String()

}
func JoinStringTest(){
	str:=joinString("Freedom","Justice","紫夜加油！","2022-03-11 22:08")
	fmt.Println(str)
}