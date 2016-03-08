package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
)

// Price 价格
type Price struct {
	img *image.Paletted
}

// CreatePrice ...
func CreatePrice() *Price {
	return &Price{img: image.NewPaletted(image.Rect(0, 0, 50, 7), palette.Plan9)}
}

//Make 价格的右上角坐标
func (p *Price) Make(m *image.Image, right, top int) {
	c := color.RGBA{0xfc, 4, 4, 0xff}
	back := color.RGBA{0xff, 0xff, 0xff, 0}

	offx := right - 50 + 1
	for x := 0; x < 50; x++ {
		for y := 0; y < 7; y++ {
			r, g, b, a := (*m).At(x+offx, y+top).RGBA()

			if uint8(r) == 0xfc &&
				uint8(g) == 0x04 &&
				uint8(b) == 0x04 &&
				uint8(a) == 0xff {
				p.img.Set(x, y, c)
			} else {
				p.img.Set(x, y, back)
			}
		}
	}
}

// Save ...
func (p *Price) Save(idx int) {
	w, _ := os.Create(fmt.Sprintf("stockPrice%d.gif", idx))
	defer w.Close()
	gif.Encode(w, p.img, nil)
}

func cutPrice(m *image.Image) {
	back := color.RGBA{0, 0, 0, 0}
	p := make([]*Price, 10, 10)
	// offy := []int{0, 17, 35, 53, 70, 88, 106, 123, 141, 159}
	last := 15

	for i := 0; i < 45; i++ {
		fmt.Println((*m).At(i, 15))
	}
	for i := 0; i < 8; i++ {
		p[i] = CreatePrice()
	All:
		for ; last < 180; last++ {
			for l := 0; l < 45; l++ {
				clast := (*m).At(l, last)
				if clast != back {
					fmt.Println(clast, "price ", i, " = ", last)
					break All
				}
			}
		}
		p[i].Make(m, 45, last)
		last += 8
		p[i].Save(i)
	}
}
