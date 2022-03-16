package basics

import "fmt"

func TestBiBao1(){
	f :=func(num int) func(x int)int{
		sum := num
		return  func(x int) int{
			sum += x
			num ++
			return sum
		}

	}
	a10:=f(10)
	a100:=f(100)

	fmt.Println(a10(1))
	fmt.Println(a10(2))
	fmt.Println(a100(101))
	fmt.Println(a100(1))
	fmt.Println(a100(0))
	fmt.Println(f(10)(1))
}

func TestBiBao2(){
	myFunc :=Counter()
	fmt.Println("myfunc",myFunc)
	fmt.Println(myFunc())
	fmt.Println(myFunc())
	fmt.Println(myFunc())
	myFunc1 :=Counter()
	fmt.Println("myfunc1:",myFunc1)
	fmt.Println(myFunc1())
	fmt.Println(myFunc1())
}

func Counter()func()int{
	i:=0
	res:=func()int{
		i+=1
		return i
	}
	fmt.Println("Counter中的内部函数：",res)
	return res
}