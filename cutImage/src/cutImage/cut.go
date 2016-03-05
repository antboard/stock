package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
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
	// 横坐标50-100处, 50pix
	// 纵坐标7-15处 8pix
	// 颜色000
	cutCode(&m)
	// 股票价格,从y15开始切10个价格
	// 字加空行共18像素
	// 字高7像素,行距10像素
	cutPrice(&m)
	// 价格数据
	cutPriceData(&m)
	//手数
	cutHands(&m)
	// 日期
	cutDate(&m)
	// 手数数据
	cutHandsData(&m)
}

func cutCode(m *image.Image) {
	c := color.RGBA{0, 0, 0, 0xff}
	back := color.RGBA{0xff, 0xff, 0xff, 0}
	img := image.NewPaletted(image.Rect(0, 0, 50, 8), palette.Plan9)

	for x := 0; x < 50; x++ {
		for y := 0; y < 8; y++ {
			r, g, b, a := (*m).At(x+50, y+7).RGBA()
			if r == 0 &&
				g == 0 &&
				b == 0 &&
				uint8(a) == 0xff {
				img.Set(x, y, c)
			} else {
				img.Set(x, y, back)
			}
		}
	}
	w, _ := os.Create(fmt.Sprintf("stockCode.gif"))
	defer w.Close()
	gif.Encode(w, img, nil)
}
