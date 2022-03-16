package pkges

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func Md5AndSha1(){
	myString :="加油"
	Md5Inst :=md5.New()
	Md5Inst.Write([]byte(myString))
	res:=Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n",res)
	Sha1Inst:=sha1.New()
	Sha1Inst.Write([]byte(myString))
	res=Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n",res)
}

func Md5File(){
	filePath :="F:/GoProject/src/kaikeba/main.go"
	infile,err:=os.Open(filePath)
	if err!=nil{
		md5f:=md5.New()
		io.Copy(md5f,infile)
		fmt.Printf("%x\n",md5f.Sum([]byte("")))
		sha1f:=sha1.New()
		io.Copy(sha1f,infile)
		fmt.Printf("%x\n",sha1f.Sum([]byte("")))
	}else{
		fmt.Println("打开文件错误：",err)
		os.Exit(1)
	}
}