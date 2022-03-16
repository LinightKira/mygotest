package pkges

import (
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"
)

// QrcodeTest 生成二维码 2022-03-16 20：54
func QrcodeTest(){
	qrcode.WriteFile("http://www.baidu.com/",qrcode.Medium,256,"golang_qrcode.png")
}

func QrcodeTest2(){
	qr,err :=qrcode.New("http://www.mi.com/",qrcode.Medium)
	if err !=nil{
		log.Fatal(err)
	}else{
		qr.BackgroundColor = color.RGBA{
			R: 50,
			G: 205,
			B: 50,
			A: 255,
		}
		qr.ForegroundColor=color.White
		qr.WriteFile(1024,"二维码高级.png")
	}
}