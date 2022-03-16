package pkges

import (
	"flag"
	"fmt"
)
//定义命令行参数 skill，从命令行输入 --skill 可以将=后的字符串传入 skillParam 指针变量
var skillParma = flag.String("skill","","skill 的提升信息")  //name 命令行参数，value 默认值，usage 提示信息
func Flag1(){
	//解析命令行参数，解析完成后，skillParam 指针变量将指向命令行传入的值
	flag.Parse()

	fmt.Println(*skillParma)  //这里就是指令的内容
	//定义一个从字符串映射到 func() 的 map，然后填充这个 map。
	var skill = map[string]func(){
		"fire":func(){
			fmt.Println("chicken fire")
		},
		"run":func(){
			fmt.Println("soldier run")
		},
		"fly":func(){
			fmt.Println("angel fly")
		},
	}
	//初始化 map 的键值对，值为匿名函数
	//如果在 map 定义中存在这个参数就调用，否则打印“技能没有找到”
	if f,ok:=skill[*skillParma];ok{  //这里是map的一个写法，判断map是否存在，存在就返回值和true
		f()  //skillParam 是一个 *string 类型的指针变量，使用 *skillParam 获取到命令行传过来的值，并在 map 中查找对应命令行参数指定的字符串的函数
	}else {
		fmt.Println("skill not found")
	}
}