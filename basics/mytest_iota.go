package basics

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	FlagNone = 1 << iota  //1
	FlagRed //2
	FlagGreen //4
	FlagBlue //8
)

// 声明芯片类型
type ChipType int
const (
	None ChipType = iota
	CPU // 中央处理器
	GPU // 图形处理器
)
func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func TestIoatWeekday()  {
	fmt.Println("Sunday:",Sunday)
	fmt.Println("Monday:",Monday)
	fmt.Println("Tuesday:",Tuesday)
	fmt.Println("Wednesday:",Wednesday)
	fmt.Println("Thursday:",Thursday)
	fmt.Println("Friday:",Friday)
	fmt.Println("Saturday:",Saturday)
	fmt.Println("====================================")
	fmt.Println("FlagNone:",FlagNone)
	fmt.Println("FlagRed:",FlagRed)
	fmt.Println("FlagGreen:",FlagGreen)
	fmt.Println("FlagBlue:",FlagBlue)


	// 输出CPU的值并以整型格式显示
	fmt.Printf("%s %d", CPU, CPU)  //因为重新写了string()方法，所以可以使用字符串输出%s
}


