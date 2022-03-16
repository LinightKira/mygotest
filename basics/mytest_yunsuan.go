package basics

import (
	"fmt"
	"math"
)

func YunSuan(){

	//计算圆的面积
	area:=func(r int)float32{
		return  float32(r)*float32(r)*math.Pi
	}(5)

	fmt.Printf("半径为5的圆，面积为：%.6f\n",area)

	//计算X天是几周零几天

	func(num int){
		weeks :=num/7
		days:=num%7

		fmt.Printf("%d天是%d周零%d天\n",num,weeks,days)
	}(46)

	fmt.Println("请输入一个年份的数字，程序会返回是否是闰年")
	var year int

	fmt.Scanln(&year)

	func (year int)(){
		str:="是闰年"
		if year %400 ==0 ||(year %4==0 && year %100!=0) {

		}else{
			str ="不是闰年"
		}

		fmt.Printf("%d年%s",year,str)
	}(year)

}



