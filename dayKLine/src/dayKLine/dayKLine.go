package main

import (
	"fmt"
	"image/gif"
	"os"
	"price"
)

func main() {
	fmt.Println("股票...")
	// 打开 gif 文件
	f, e := os.Open("/Users/jiangyichun/Downloads/sh601006.gif")
	if e != nil {
		fmt.Println("open err:", e)
	}
	defer f.Close()
	//
	m, e := gif.Decode(f)
	if e != nil {
		fmt.Println("...Decode err:", e)
	}

	bounds := m.Bounds()
	if bounds.Max.X != 545 ||
		bounds.Max.Y != 300 {
		fmt.Println("...gif size is ", bounds, " NOT 545*300")
	}

	// 计算股票代码
	// 横坐标50-100处
	// 纵坐标7-15处
	price.GetPrices(&m)
}
