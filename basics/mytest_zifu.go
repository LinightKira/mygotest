package basics

import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
Go语言同样支持 Unicode（UTF-8），因此字符同样称为 Unicode 代码点或者 runes，并在内存中使用 int 来表示。在文档中，
一般使用格式 U+hhhh 来表示，其中 h表示一个 16 进制数。
在书写 Unicode 字符时，需要在 16 进制数之前加上前缀\u或者\U。因为 Unicode 至少占用 2 个字节，
所以我们使用 int16 或者 int 类型来表示。如果需要使用到4 字节，则使用\u前缀，如果需要使用到 8 个字节，则使用\U前缀
*/

func TestUnicodeDemo() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	var ch4 = 98
	fmt.Printf("%d - %d - %d - %d\n", ch, ch2, ch3, ch4) // integer
	fmt.Printf("%c - %c - %c - %c\n", ch, ch2, ch3, ch4) // character
	fmt.Printf("%X - %X - %X - %X\n", ch, ch2, ch3, ch4) // UTF-8 bytes
	fmt.Printf("%U - %U - %U - %U", ch, ch2, ch3, ch4)   // UTF-8 code point
}

func TestCharString() {
	var ch byte
	var str string
	//字符 单引号
	ch = 'a'
	fmt.Println("ch=", ch)
	//字符串都是隐藏了一个结束符‘\0’
	str = "Freedom"
	fmt.Println("str内容是：", str)
	fmt.Println("str长度是：", len(str))
	fmt.Printf("str[1]=%c,str[3]=%c\nS", str[1], str[3])

	for k, v := range str {
		fmt.Printf("Str[%v]=%c\t", k, v)
	}

	fmt.Println("")
	//byte &rune &uint8
	//byte和 uint8本质没有任何区别
	//rune，占用4个字节，共32位比特位，所以它和 uint32 本质上也没有区别。它表示的是一个 Unicode字符（Unicode是一个可以表示世界范围内的绝大部分字符的
	//编码规范）

	var b1 byte = 65
	var u1 uint8 = 66
	var u2 uint = 67
	var u3 uint64 = 68
	var u4 int = 69
	var r1 rune = 70
	fmt.Println("b1(byte)=", b1)
	fmt.Println("u1(uint8)=", u1)
	fmt.Println("u2(uint8)=", u2)
	fmt.Println("u3(uint64)=", u3)
	fmt.Println("u4(int)=", u4)

	fmt.Println("r1(rune)=", r1)
	fmt.Printf("b1(byte)=%c\n", b1)
	fmt.Printf("u1(uint8)=%c\n", u1)
	fmt.Printf("u2(uint)=%c\n", u2)
	fmt.Printf("u3(uint64)=%c\n", u3)
	fmt.Printf("u4(int)=%c\n", u4)

	fmt.Printf("r1(rune)=%c\n", r1)
	fmt.Printf("b1(byte).size=%d\t u1(uint8).size=%d\t r1(rune).size=%d\n", unsafe.Sizeof(b1), unsafe.Sizeof(u1), unsafe.Sizeof(r1))
	fmt.Printf("u1(uint8).size=%d\t u2(uint).size=%d\t u3(uint64).size=%d\t u4(int).size=%d\n",
		unsafe.Sizeof(u1), unsafe.Sizeof(u2), unsafe.Sizeof(u3), unsafe.Sizeof(u4))

	//字符串部分
	var str1 string = "hello"
	str2 := []byte{'h', 'e', 'l', 'l', 'o'}
	str3 := []byte{104, 101, 108, 108, 111}

	fmt.Printf("str1:%s\n", str1)
	fmt.Printf("str2:%s\n", str2)
	fmt.Printf("str3:%s\n", str3)

	fmt.Println("str1=", str1)
	fmt.Println("str2=", str2)
	fmt.Println("str3=", str3)

	type myStringDemo string //自定义类型

	var str4 myStringDemo = "freedom"
	fmt.Println("str4=", str4)

	str5 := "你好，世界" //utf-8 每个字占用3个字节
	fmt.Printf("str5=%s,长度为：%d\n", str5, len(str5))

	var emptyByte []byte

	fmt.Println("空字符数组的值为：", emptyByte)

	//Unicode示范
	fmt.Printf("测试\\u:\u0023和\\U:\U00000023\n")
	fmt.Printf("\\u:\u0023")

	fmt.Printf("\n")

	str = "1234567890"
	for i, i2 := range str {
		fmt.Printf("str[%d]:%c\t", i, i2)
	}
	fmt.Println("")
	str1 = str[2:]
	fmt.Printf("str1[2:]长度为：%d,内容为:", len(str1))
	for i := 0; i < len(str1); i++ {
		fmt.Printf("str1[%d]:=%c\t", i, str1[i])
	}
	fmt.Println("")
	str1 = str[:4]
	fmt.Printf("str1[:4]长度为：%d,内容为:", len(str1))
	for i := 0; i < len(str1); i++ {
		fmt.Printf("str1[%d]:=%c\t", i, str1[i])
	}

}

// StringChange 字符串类型转换
func StringChange() {
	//基本类型转 string 类型 第一种写法
	var intNum int = 100
	var flotNum float32 = 1.23456
	var boolKey bool = true
	var chars []byte = []byte{'f', 'r', 'e', 'd', 'o', 'm'}

	str := fmt.Sprintf("%d", intNum)
	fmt.Println("intNum=", str)
	str = fmt.Sprintf("%f", flotNum)
	fmt.Println("flotNum=", str)
	fmt.Printf("floaNum2=%q\n", str)
	str = fmt.Sprintf("%t", boolKey)
	fmt.Println("boolKey=", str)
	str = fmt.Sprintf("%s", chars)
	fmt.Println("chars=", str)

	//第二种写法 strconv 要比第一种效率高
	str = strconv.FormatInt(int64(intNum), 10)
	fmt.Printf("str type %T,str=%q\n", str, str)

	str = strconv.FormatFloat(float64(flotNum), 'f', 10, 64)
	fmt.Printf("str type %T,str=%q\n", str, str)

	str = strconv.FormatBool(boolKey)
	fmt.Printf("str type %T,str=%q\n", str, str)

	//str 与 int之间进行转换
	//strconv 还有一个包 itoa  整型转字符串
	str = strconv.Itoa(intNum) //int 转str
	fmt.Printf("str type %T,str=%q\n", str, str)
	str = "-1024"
	intNum2, _ := strconv.ParseInt(str, 10, 32) //字符串转整型,区分正负数
	fmt.Printf("str type %T,str=%v\n", intNum2, intNum2)

	//Atoi  字符串转整型
	str = "1560"
	intNum, _ = strconv.Atoi(str)
	fmt.Printf("str type %T,str=%v\n", intNum, intNum)

	str = "3090"
	intNum3, _ := strconv.ParseUint(str, 10, 0)
	fmt.Printf("str type %T,str=%v\n", intNum3, intNum3)

	//bool类型转换
	//ParseBool() 函数用于将字符串转换为 bool 类型的值，它只能接受 1、0、t、f、T、F、true、false、True、False、TRUE、FALSE，其它的值均返回错误，函数签名如
	//下。

	parseKey, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println("ParseBool转化错误:", err)
	} else {
		fmt.Printf("parsekey type %T,parsekey=%v\n", parseKey, parseKey)
	}

	str = "true"
	parseKey, err = strconv.ParseBool(str)
	if err != nil {
		fmt.Println("ParseBool转化错误:", err)
	} else {
		fmt.Printf("parsekey type %T,parsekey=%v\n", parseKey, parseKey)
	}
	/*
		FormatFloat() 函数用于将浮点数转换为字符串类型，函数签名如下：
		func FormatFloat(f float64, fmt byte, prec, bitSize int) string
		参数说明：
		• bitSize 表示参数 f 的来源类型（32 表示 float32、64 表示 float64），会据此进行舍入。
		fmt 表示格式，可以设置为“f”表示 -ddd.dddd、“b”表示 -ddddp±ddd，指数为二进制、“e”表示 -d.dddde±dd 十进制指数、“E”表示 -d.ddddE±dd 十进制指
		数、“g”表示指数很大时用“e”格式，否则“f”格式、“G”表示指数很大时用“E”格式，否则“f”格式。
		•
		prec 控制精度（排除指数部分）：当参数 fmt 为“f”、“e”、“E”时，它表示小数点后的数字个数；当参数 fmt 为“g”、“G”时，它控制总的数字个数。如果
		prec 为 -1，则代表使用最少数量的、但又必需的数字来表示 f。
	*/

	str = strconv.FormatFloat(3.1415926, 'g', 4, 32)
	fmt.Printf("str type %T,str=%v\n", str, str)

	//append Demo
	b10 := []byte("int (base 10):")
	fmt.Printf("b10的类型是：%T,\t值是：%s\n", b10, b10)
	b10 = strconv.AppendInt(b10, 123, 10)
	fmt.Printf("b10的类型是：%T,\t值是：%s\n", b10, b10)

	//int类型的相关转换
	var v1 uint = 100
	v2 := int16(v1)
	v3 := int64(v2)

	//浮点数转换
	v4 := float32(v1)
	v3 = int64(v4)

	if v3 == 0 {
	} //为了编译过 搞了一个临时判断用

	sv1 := []byte("生活不止眼前的苟且")
	sv1 = strconv.AppendQuote(sv1, ",还有诗和远方的田野")
	fmt.Println(string(sv1))

	st1, _ := strconv.ParseInt("-87654321", 10, 16)
	fmt.Println(st1)
	st2, _ := strconv.Atoi("-1234")
	fmt.Println(st2)
	st3 := strconv.Itoa(3090)
	fmt.Println(st3)
	st4, _ := strconv.ParseBool("true") //string 转布尔
	fmt.Println(st4)
	st5 := strconv.FormatBool(st4) //布尔转string
	fmt.Println(st5)

	st6 := strconv.FormatFloat(123.14159, 'f', 10, 32) //10表示小数点后10位
	fmt.Println(st6)
	st7, _ := strconv.ParseFloat(st6, 32)
	fmt.Println(st7)

	q1 := strconv.Quote("Hello, 世界") // 为字符串加引号
	fmt.Println(q1)
	q2 := strconv.QuoteToASCII("Hello, 世界") // 将字符串转化为 ASCII 编码
	fmt.Println(q2)

	v1 = 234
	fmt.Printf("v1:=%d\t", v1)
	fmt.Printf("v1:=%5d\t", v1)
	fmt.Printf("v1:=%05d\t\n", v1)

	v4 = 123.456
	fmt.Printf("f4=%f\t", v4)
	fmt.Printf("f4=%.7f\t", v4)
	fmt.Printf("f4=%.8f\t", v4)
}

//2022-03-14 字符串常用的系统函数
func StrTest() {
	str := "生活不只眼前的苟且，gogogo"
	fmt.Println("str的长度为：", len(str))

	//字符串遍历，同时处理有中文的问题
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\t", r[i])
	}
	fmt.Printf("\n")

	//字符串转整型
	str = "1024"
	num, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println("字符串转整型：", num)
	} else {
		fmt.Println("转化失败：:", err)
	}

	//整型转字符串
	num = 3090
	str = strconv.Itoa(num)

	str="宠辱不惊，闲看庭前花开花落，2020"
	fmt.Println("整型转字符串：", str)
	//字符串转[]byte
	bStr :=[]byte(str)
	fmt.Printf("%v",bStr)


}
