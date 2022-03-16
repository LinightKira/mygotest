package basics

import "fmt"

func TestFmtTypeDemo()  {
	var mystring string
	fmt.Println("请输入需要打印的字符串：")
	fmt.Scanln(&mystring)

	fmt.Printf("刚才输入的内容为：%s，类型为：%T\n",mystring,mystring)

	var myint int
	fmt.Println("请输入需要打印的数字：")
	fmt.Scanln(&myint)
	fmt.Printf("刚才输入的数字为：%v，十进制为：%d，二进制为：%b，八进制为：%o\n", myint,myint,myint,myint)
	fmt.Printf("刚才输入的数字为：%v，十六进制为：%x，十六进制为：%X，地址为：%p\n", myint,myint,myint,&myint)
	fmt.Printf("刚才输入的数字为：%v，字符为：%c\n", myint,myint)
}
