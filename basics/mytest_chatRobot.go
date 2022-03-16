package basics

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChatRobot() {
	//准备从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字：")
	input, err := inputReader.ReadString('\n') //读取数据直到碰到\n为止
	if err != nil {
		fmt.Printf("出错了：%s\n", err)
		//异常退出
		os.Exit(1)
	} else {
		//用切片操作 删除最后的\n
		name := input[:len(input)-2]
		fmt.Printf("你好%s,请问有什么可以帮到您？\n", name)

	}
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("出错了：%s\n", err)
			continue
		}
		input = input[:len(input)-2]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "没事", "再见", "886":
			fmt.Println("再见")
			os.Exit(0)
		default:
			fmt.Println("你在说神马？")
		}

	}
}

func ChatRobot2() {
	//准备从标准输入读取数据
	var name string
	fmt.Println("请输入你的名字：")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Printf("出错了：%s\n", err)
		//异常退出
		os.Exit(1)
	} else {
		//用切片操作 删除最后的\n
		fmt.Printf("你好%s,请问有什么可以帮到您？\n", name)

	}
	var input string
	for {
		_, err = fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("出错了：%s\n", err)
			continue
		}
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "没事", "再见", "886":
			fmt.Println("再见")
			os.Exit(0)
		default:
			fmt.Println("你在说神马？")
		}
	}
}
