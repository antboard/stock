package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
)

const (
	w = 480
	h = 84
)

// HandsData 手数
type HandsData struct {
	img *image.Paletted
}

// CreateHandsData ...
func CreateHandsData() *HandsData {
	return &HandsData{img: image.NewPaletted(image.Rect(0, 0, w, h), palette.Plan9)}
}

// MakeHandsData 价格的右上角坐标
func (p *HandsData) MakeHandsData(m *image.Image, left, top int) {
	g := color.RGBA{0x0, 0xa8, 0x0, 0xff}
	r := color.RGBA{0xfc, 0x54, 0x54, 0xff}
	back := color.RGBA{0xff, 0xff, 0xff, 0}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := (*m).At(x+left, y+top)
			if c == g {
				p.img.Set(x, y, g)
			} else if c == r {
				p.img.Set(x, y, r)
			} else {
				p.img.Set(x, y, back)
			}
		}
	}
}

// Save ...
func (p *HandsData) Save() {
	f, _ := os.Create(fmt.Sprintf("stockHandsData.gif"))
	defer f.Close()
	gif.Encode(f, p.img, nil)
}

func cutHandsData(m *image.Image) {
	p := CreateHandsData()
	p.MakeHandsData(m, 50, 201)
	p.Save()
}
