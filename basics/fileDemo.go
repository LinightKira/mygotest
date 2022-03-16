package basics

import (
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

//文件哈希值校验Demo 注意对比文件修改前后的内容
func FileHashCheckDemo() {
	path := "F:/GoProject/src/kaikeba/goBasics/basics/fileHashDemo.go"
	file, _ := os.Open(path)
	md5run := sha512.New()
	io.Copy(md5run, file)
	fmt.Printf("%x\n", md5run.Sum(nil))

}

func FileDeferCloseDemo() {
	path := "F:/紫夜世离/MyGolandAccount.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打开错误：", err)
		return
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return
	}

	fmt.Println("文件大小为：", info.Size())
}
