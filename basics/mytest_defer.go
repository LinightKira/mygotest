package basics

import "fmt"

// TestDefer DeferDemo
//defer是当前函数结束时执行，仔细理解一下，很重要
//在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行
func TestDefer() {
	fb()

	fmt.Println("=================================")

	i := func() (r int) {
		defer func()  {
			r++
			fmt.Println("defer r=", r) //defer输出的值，就是定义时的值。而不是defer真正执行时的变量值
		}()
		r = 10
		return
	}()//注意defer的嵌套关系，匿名函数结束时，defer就开始执行了
	fmt.Println("i=",i)
	j:=fd()
	fmt.Println("j=",j)


}
func fd()int{
	r:=0
	defer func(){
		r++
		fmt.Println("defer r2:",r)
	}()
	return  r
}
func fb(){
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

//证明，只要声明函数的返回值变量名称，就会在函数初始化时候为之赋值为0，而且在函数体作用域可见
//反回值 声明时函数会初始化赋值为0，即t=0
//t=0，t=i,defer t =i+3
//先return，再defer，所以在执行完return之后，还要再执行defer里的语句，依然可以修改本应该返回的结果
func deferFunc1(i int)(t int){
	t =i
	defer func(){
		t+=3
	}()
	return t
}
func deferFunc2(i int)int{
	t:=i
	defer func(){
		t+=3
	}()
	return t
}
//t=0,t=2,defer t=3
func deferFunc3(i int)(t int){
	defer func() {
		t+=i
	}()
	return 2
}
func deferFunc4()(t int){
	defer func(i int) {
		fmt.Println("4-1:",i)
		fmt.Println("4-2:",t)
	}(t)
	t=1
	return 2
}

func TestDefer2(){
	fmt.Println("1:",deferFunc1(1))  //4
	fmt.Println("2:",deferFunc2(1)) //1
	fmt.Println("3:",deferFunc3(1)) //3
	deferFunc4() //0,2
	fmt.Println("4:",deferFunc4())//2
}

// TestDeferPanic
//遇到panic时，遍历本协程的defer链表，并执行defer。在执行defer过程中:遇到recover则停止panic，返回recover处继续往下执行。
//如果没有遇到recover，遍历完本协程的defer链表后，向stderr抛出panic信息
func TestDeferPanic(){
	deferCall()
	fmt.Println("正常结束")
}
func TestDeferPanic2(){
	deferCall2()
	fmt.Println("正常结束")
}

func deferCall(){
	defer func() {fmt.Println("defer:panic之前1")}()
	defer func() {fmt.Println("defer:panic之前2")}()
	panic("异常内容") //触发Defer出栈
	defer func() {fmt.Println("defer:panic之后，永远不会执行")}()
}



func deferCall2(){
	defer func() {
		fmt.Println("defer:panic之前1")
		if err :=recover();err!=nil{
			fmt.Println("Recover err:",err)
		}
	}()

	defer func() {fmt.Println("defer:panic之前2")}()
	panic("异常内容") //触发Defer出栈
	defer func() {fmt.Println("defer:panic之后，永远不会执行")}()
}