package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
)

type myColor struct {
	r, g, b, a, cnt uint32
	Img             *image.Paletted
}

func main() {
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
	fmt.Println("...gif size is ", bounds)

	// 统计颜色.
	clrs := make([]*myColor, 0, 100)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
	LP:
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			// 查找重复颜色
			for i := 0; i < len(clrs); i++ {
				if r == clrs[i].r &&
					g == clrs[i].g &&
					b == clrs[i].b &&
					a == clrs[i].a {
					clrs[i].cnt++
					continue LP
				}
			}
			clrs = append(clrs, &myColor{r, g, b, a, 1, nil})
		}
	}
	all := bounds.Max.Y * bounds.Max.X
	fmt.Println("...All colors is ", all)
	fmt.Println("...colors is: ", len(clrs))
	for i := 0; i < len(clrs); i++ {
		fmt.Printf("{%x, %x, %x, %x, %d, %d}\r\n",
			clrs[i].r, clrs[i].g, clrs[i].b, clrs[i].a, clrs[i].cnt, clrs[i].cnt*uint32(10000)/uint32(all))
	}

	// 删除无用数据
	clrsNew := make([]*myColor, 0, 100)
	for i := 0; i < len(clrs); i++ {
		if clrs[i].cnt > 50 {
			clrsNew = append(clrsNew, clrs[i])
		}
	}
	clrs = clrsNew
	colors := len(clrs)

	for i := 0; i < colors; i++ {
		clrs[i].Img = image.NewPaletted(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y), palette.Plan9)
	}
	for i := 0; i < colors; i++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				clrs[i].Img.Set(x, y, color.RGBA{0xff, 0xff, 0xff, 0})
			}
		}
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			// 查找重复颜色
			for i := 0; i < colors; i++ {
				if r == clrs[i].r &&
					g == clrs[i].g &&
					b == clrs[i].b &&
					a == clrs[i].a {
					clrs[i].Img.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
					break
				}
			}
			// 打标
			if x%10 == 0 &&
				y%10 == 0 {
				for i := 0; i < colors; i++ {
					clrs[i].Img.Set(x, y, color.RGBA{uint8(0), uint8(0), uint8(0xff), uint8(0xff)})
				}
			}
		}
	}
	for i := 0; i < colors; i++ {
		w, _ := os.Create(fmt.Sprintf("r%xg%xb%x.gif", uint8(clrs[i].r), uint8(clrs[i].g), uint8(clrs[i].b)))
		defer w.Close()
		gif.Encode(w, clrs[i].Img, nil)
	}

}
