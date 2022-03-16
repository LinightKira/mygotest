package basics

import (
	"fmt"
	"math/rand"
	"time"
)

//测试九九乘法表
func TestNineAndNine() {
	var mode int
	for {
		fmt.Println("请请入要查看的乘法表格式:")
		fmt.Scanln(&mode)
		switch mode {
		case 0:
			fmt.Println("程序已退出")
			return
		case 1:
			for i := 1; i < 10; i++ {
				for j := 1; j < 10; j++ {
					if i >= j {
						fmt.Printf("%d*%d=%d\t", i, j, i*j)
					}
				}
				fmt.Printf("\n")
			}
		case 2:
			for i := 1; i < 10; i++ {
				for j := 1; j < 10; j++ {
					if j >= i {
						fmt.Printf("%d*%d=%d\t", i, j, i*j)
					} else {
						fmt.Printf("\t")
					}
				}
				fmt.Printf("\n")
			}
		case 3:
			for i := 1; i < 10; i++ {
				for j := 1; j < 10; j++ {
					if i >= 10-j {
						fmt.Printf("%d*%d=%d\t", i, j, i*j)
					} else {
						fmt.Printf("\t")
					}
				}
				fmt.Printf("\n")
			}
		case 4:
			for i := 1; i < 10; i++ {
				for j := 1; j < 10; j++ {
					if j <= 10-i {
						fmt.Printf("%d*%d=%d\t", i, j, i*j)
					} else {
						fmt.Printf("\t")
					}
				}
				fmt.Printf("\n")
			}
		default:
			for i := 1; i < 10; i++ {
				for j := 1; j < 10; j++ {
					fmt.Printf("%d*%d=%d\t", i, j, i*j)
				}
				fmt.Printf("\n")
			}
		}
	}
}

func MyClassScore() {
	var classNum int
	fmt.Println("请输入班级人数")
	fmt.Scanln(&classNum)
	sum := 0.0
	score := 0.0
	for i := 0; i < classNum; i++ {
		fmt.Printf("请输入第%d位同学的成绩", i+1)
		fmt.Scanf("%f\n", &score)
		sum += score
	}

	fmt.Printf("全班级%d人的总成绩为：%.0f\t", classNum, sum)
	average := sum / float64(classNum)
	fmt.Printf("平均成绩为：%.2f", average)
}

//水仙花数
func NarcissisticNumber() {
	number := 1
	for i := 100; i < 1000; i++ {
		func(num int) {
			bai := num / 100
			shi := num % 100 / 10
			ge := num % 10

			if bai*bai*bai+shi*shi*shi+ge*ge*ge == i {
				fmt.Printf("第%d个水仙花数为：%d\n", number, num)
				number++
			}
		}(i)
	}
}

//生成1-100随机数，看随机多少次可以随到99
func RandomNum() {
	count := 0

	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1 //[0,100)
		fmt.Println("随机数：", n)
		count++
		if n == 99 {
			break
		}
	}

	fmt.Println("一共随机的次数为：", count)
}

//求1-100求和 当和大于20时当前的数
func SumTest() {
	sum := 0
	for i := 1; i < 101; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("和大于20，当前数为：", i)
			break
		}
	}
}

//验证登录次数
func LoginTest() {
	isOK := false
	for i := 0; i < 3; i++ {

		fmt.Println("请输入账号名称：")
		var username string
		fmt.Scanln(&username)
		fmt.Println("请输入密码：")
		var pwd string
		fmt.Scanln(&pwd)
		if username == "Freedom" && pwd == "123" {
			isOK = true
			break
		}

		fmt.Printf("还有%d次登录机会\n", 2-i)

	}
	if isOK {
		fmt.Println("登录成功")
	}else{
		fmt.Println("登录次数用尽，等等吧")
	}

}

//1-100奇数
func Jishu(){
	for i := 1; i <101 ; i++ {
		if i%2==0{
			continue
		}
		fmt.Printf("%d\t",i)

	}
}