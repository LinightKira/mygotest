package basics

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func ChengfabiaoDemo() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d*%d=%2d  ", i, j, i*j)

		}
		fmt.Println()
	}

	fmt.Println()

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j >= 10-i {
				fmt.Printf("%d*%d=%2d  ", j, i, i*j)
			} else {
				fmt.Printf("%8s", "")
			}

		}
		fmt.Println()
	}

}

func SinCosDemo() {
	const size = 300                                   //图像大小
	pic := image.NewGray(image.Rect(0, 0, size, size)) //根据给定大小创建灰度图
	////遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			//填充为白色
			pic.SetGray(x, y, color.Gray{255})

		}
	}

	//从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		//让sin的值范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size
		//sin的幅度为一半的像素，向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		//y := size/2 - math.Cos(s)*size/2
		//用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})

	}
	file, _ := os.Create("sin.png") //写入文件
	png.Encode(file, pic)           //将image信息写入文件
	file.Close()                    //关闭文件
}
