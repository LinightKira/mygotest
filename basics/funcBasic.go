package basics

import (
	"fmt"
	"strings"
	"time"
)

const (
	SecondsPreMinute = 60                    //一分钟有60s
	SecondsPreHour   = 60 * SecondsPreMinute //一小时的秒数
	SecondsPreDay    = 24 * SecondsPreHour   //一天的秒数
)

func GetTime(seconds int) (day, hour, minute int) {
	day = seconds / SecondsPreDay
	hour = seconds / SecondsPreHour
	minute = seconds / SecondsPreMinute
	return
}

//可变参数，类型一致
func Add(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func Mul(args ...int) int {
	var sum = 0
	for _, v := range args {
		sum *= v
		//	fmt.Println(k)
	}
	return sum
}

//通过语法糖和接口，传可变参数
func ShowDatasType(list ...interface{}) {
	for _, v := range list {
		switch v.(type) {
		case int:
			fmt.Println("int类型")
		case string:
			fmt.Println("string类型")
		case byte, rune:
			fmt.Println("字符类型")
		case float32, float64:
			fmt.Println("浮点数类型")
		default:
			fmt.Println("其他类型")

		}
	}
}

type MyFuncType func(int, int) int //自定义函数类型

func Addx(a int, b int) int {
	return a + b
}

func Subx(a int, b int) int {
	return a - b
}

func MyFuncTypeDemo() {
	var ff MyFuncType

	ff = Addx

	fmt.Println("ADD:", ff(1, 2))
	ff = Subx
	fmt.Println("SUB:", ff(9, 1))
}

//匿名函数调用
func NimingFuncDemo() {
	useTimes(func(arg int) (sum int) {
		for i := 0; i < arg; i++ {
			sum += i
		}
		return sum
	}, 100000000)
}

func useTimes(myfunc func(int) int, arg int) {
	startTime := time.Now()            //开始时间
	myfunc(arg)                        //函数调用
	fmt.Println(time.Since(startTime)) //结束时间
}

//函数递归 闭包调用
func squares() func() int {
	var x int
	fmt.Println("x:", x)
	fmt.Println("&x:", &x)
	return func() int {
		x++
		return x * x
	}
}

/*
squares() 是一个函数，返回的数据类型是 func()int
返回的是一个匿名函数，但是这个匿名函数引用到函数外的x,因此这个函数和x组成一个整体，构成闭包
可以理解为：闭包是类，函数是操作，x是字段，函数和它使用到的x构成闭包
当我们反复调用f函数时，因为x时初始化1次，因此每一次调用就进行累计
我们要搞清楚闭包的关键，就是要分析出返回的函数它使用（引用）到哪些变量，因为函数和它引用的变量共同构成闭包
*/

//闭包：闭包就是一个函数与其相关的引用环境组合的一个整体（实体）
/*
函数值不仅仅是一串代码，还记录了状态
在squares中定义的匿名内部函数可以访问和更新squares中的局部变量
这意味着匿名函数和squares中，存在变量引用，这就是函数值属于引用类型和函数值不可比较的原因
Go使用闭包(closures)技术实现函数值，Go程序员也把函数值叫做闭包
通过这个例子，我们看到变量的生命周期不由它的作用域决定，squares返回后，变量x仍然隐式的存在于f中
闭包不关心捕获了的变量和常量是否已经超出了作用域，所以只要闭包还在使用它，这些变量就会还存在
*/
func BiBaoFuncDemo() {
	f := squares()
	for i := 0; i < 6; i++ {
		fmt.Println(f())
	}
}

//闭包的高级使用案例
//编写一个函数 makeSuffix(suffix string)可以接受一个文件后缀名(比如.jpg),并返回一个闭包
//调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg)，则返回文件名.jpg
//要求使用闭包的方式完成
//string.HasSuffix ,该函数可以判断某个字符串是否有指定后缀

func makeSuffix(suffix string) func(string) string {
	//如果name 没有指定的后缀名，就加上，否则就返回原来的名字
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func BiBaoFileNameDemo() {
	f := makeSuffix("jpg")
	fmt.Println("文件名处理后：", f("原文件名是PNG.png"))
	fmt.Println("文件名处理后：", f("源文件是JPG.jpg"))
}

//闭包计数器，非常好的实际用途
func BiBaoCountDemo() {
	fcount := count()
	for i := 0; i < 50; i++ {
		fmt.Println("计数器次数为：", fcount())
	}

}

func count() func() int {
	sum := 0
	return func() int {
		sum++

		return sum
	}
}

//高阶函数
/*
 高阶函数：
 根据go语言的数据类型的特点，可以将一个函数作为另一个函数的参数。
fun1(),fun2()
 将fun1函数作为了fun2这个函数的参数。
fun2函数：就叫高阶函数
 接收了一个函数作为参数的函数，高阶函数
fun1函数：回调函数
 作为另一个函数的参数的函数，叫做回调函数。
*/

func multiply(a, b int) int {
	return a * b
}

func multiply2(a, b int) int {
	return a << b
}

//装饰器模式实现
//为函数类型设置别名，提高代码可读性
type MultiPlyFunc func(int, int) int

//通过高阶函数在不侵入原有函数实现的前提下计算乘法函数执行时间
func execTime(f MultiPlyFunc) MultiPlyFunc {
	return func(a, b int) int {
		start := time.Now()      //起始时间
		c := f(a, b)             //执行乘法运算函数
		end := time.Since(start) //函数执行完毕耗时
		fmt.Printf("执行耗时：%v\n", end)
		return c //返回计算结果
	}
}

//写法2
func execTime2(myfunc func(a, b int) int, a, b int) int {
	start := time.Now() //开始时间
	x := myfunc(a, b)
	fmt.Printf("执行耗时：%v\n", time.Since(start))
	return x
}

//高阶函数Demo，使用高阶函数，在不侵入函数的前提下计算函数的耗时
func ProFuncDemo() {
	//高阶函数测试1
	a, b := 5, 10

	f := execTime(multiply)

	fmt.Println("乘法函数方法1的结果：", f(a, b))

	g := execTime2(multiply2, a, b)

	fmt.Println("乘法函数方法2的结果：", g)

}

//递归函数
//用递归的方式求N的阶乘

func factorial(num int) int {
	if num == 0 {
		return 0

	} else {
		return factorial(num-1) + num
	}
}

//递归函数，求100的阶乘
func FactorialDemo() {
	fmt.Println("100的阶乘为：", factorial(100))
}
