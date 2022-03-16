package basics

import (
	"fmt"
	"time"
)

func InterfaceTest() {
	f := func(f ...interface{}) {
		for _, i2 := range f {
			switch i2.(type) {
			case int, uint:
				fmt.Println("这是一个int类型的数据：", i2)
			case float32, float64:
				fmt.Println("这是一个float类型：", i2)
			case string:
				fmt.Println("这是一个string类型：", i2)
			default:
				fmt.Printf("这个类型是%T，值是%v\n", i2, i2)

			}
		}
	}
	f(1, 1.23, "哈哈哈哈", 'a')
}

//可变参数
//在多个可变参数函数中传递参数
//可变参数变量是一个包含所有参数的切片，如果要将这个含有可变参数的变量传递给下一个可变参数函数，可以在传递时给可变参数变量后面添加...，这样就可以将
//切片中的元素进行传递，而不是传递可变参数变量本身。

func rPrint(i ...int) {
	for _, i3 := range i {
		fmt.Println(i3)
	}
}

func pPrint(i ...int) {
	rPrint(i...) //注意这里的可变参数传参方法
	//如果只传递一部分 可以使用切片的传值方法
	rPrint(i[:2]...)
}


//高阶函数
type MultiPlyFunc2 func(int,int)int //定义一个函数类型
func multiply3(a,b int)int{
	return  a*b
}

//通过高级函数在不侵入原有函数实现的前提下计算执行时间
func execTime3(f MultiPlyFunc2)MultiPlyFunc2{
	return func(a,b int)int{
		start:=time.Now() //起始时间
		c:=f(a,b)  //执行要执行的函数
		end :=time.Since(start)//函数执行完毕耗时
		fmt.Println("-----执行耗时：%v-----\n",end)
		return c //返回计算结果
	}
}

//高级函数
func GaoJiFunc()  {
	a,b:=10,999
	decorator :=execTime3(multiply3)
	//执行修饰器返回函数
	c:=decorator(a,b)
	fmt.Printf("%d*%d=%d\n",a,b,c)

}

//递归函数，实现-斐波那契数列
func  FibonacciTest(){
	fmt.Println("===============第一种写法==================")
	for i := 1; i < 10; i++ {
		fmt.Printf("%d\t",fibonacci(i))
	}
	fmt.Printf("\n")
fmt.Println("===============第二种写法==================")
	for i := 1; i < 10; i++ {
		fmt.Printf("%d\t",fibonacci2(i))
	}

	fmt.Printf("\n")
	fmt.Println("===============第三种写法==================")
	for i := 1; i < 10; i++ {
		fmt.Printf("%d\t",fibonacci3(i))
	}
}

func  fibonacci(n int)  int{
	if n<=2 {
		return 1
	}
	return fibonacci(n-1)+fibonacci(n-2)
}

//斐波那契数列 优化写法，使用缓存的方式提升计算效率
//一方面是因为递归函数调用产生的额外开销，另一方面是因为目前这种实现存在着重复计算，比如我在计算 fibonacci(50)
//时，会转化为计算 fibonacci(49) 与 fibonacci(48) 之和，然后我在计算 fibonacci(49) 时，又会转化为调用 fibonacci(48) 与
//fibonacci(47) 之和，这样一来 fibonacci(48) 就会两次重复计算，这一重复计算就是一次新的递归（从序号 48 递归到序号 1），以
//此类推，大量的重复递归计算堆积，最终导致程序执行缓慢。
//通过内存缓存技术优化递归函数性能
//我们先对后一个原因进行优化，通过缓存中间计算结果来避免重复计算，从而提升递归函数的性能
//这种优化是在内存中保存中间结果，所以称之为内存缓存技术（memoization），这种内存缓存技术在优化计算成本相对昂贵的函数调用时非常有用
const MAX = 50
var fibs [MAX]int
// 缓存中间结果的递归函数优化版
func fibonacci2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	index := n - 1
	if fibs[index] != 0 {
		return fibs[index]
	}
	num := fibonacci2(n-1) + fibonacci2(n-2)
	fibs[index] = num   //算过一次的值就存上，下一次用的时候直接返回
	return num
}

//第三种写法 尾递归
//通过尾递归优化递归函数性能
//我们知道函数调用底层是通过栈来维护的，对于递归函数而言，如果层级太深，同时保存成百上千的调用记录，会导致这个栈越来越大，消耗大量内存空间，严重情况下会导致栈溢出（stack overflow），
//为了优化这个问题，可以引入尾递归优化技术来重用栈，降低对内存空间的消耗，提升递归函数性能。
//尾递归优化是函数式编程的重要特性之一，在了解尾递归优化前，我们先来看看什么是尾递归。
//在计算机科学里，尾调用是指一个函数的最后一个动作是调用一个函数（只能是一个函数调用，不能有其他操作，比如函数相加、乘以）常量等）：
//若这个函数在尾位置调用自身，则称这种情况为尾递归，它是尾调用的一种特殊情形。尾调用的一
//个重要特性是它不是在函数调用栈上添加一个新的堆栈帧 —— 而是更新它，尾递归自然也继承了这一特性，这就使得原来层层递进的
//调用栈变成了线性结构，因而可以极大优化内存占用，提升程序性能，这就是尾递归优化技术
//尾递归的实现需要重构之前的递归函数，确保最后一步只调用自身，要做到这一点，就要把所有用到的内部变量/中间状态变成函数参数
func fibonacci3(n int) int {
	return fibonacciTail(n, 1, 1) // F(1) = 1, F(2) = 1
}
func fibonacciTail(n, first, second int) int {
	if n < 2 {
		return first
	}
	return fibonacciTail(n-1, second, first+second)
}

// MonkeyPeachs 递归函数，求猴子的桃子数
//猴子吃桃子问题 有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个！以后每天猴子都吃其中的一半，然后
//再多吃一个。当到第10天时，想再吃时（还没吃），发现只有 1 个桃子了。问题：最初共多少个桃子？
func MonkeyPeachs(){
	fmt.Println("10天的桃子个数为：",peachs(10))
}

func peachs(n int)int {
	if n ==1{
		return 1
	}
	return (peachs(n-1)+1)*2
}