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
	wPriceData = 480
	hPriceData = 160
)

// PriceData 价格
type PriceData struct {
	img *image.Paletted
}

// CreatePriceData ...
func CreatePriceData() *PriceData {
	return &PriceData{img: image.NewPaletted(image.Rect(0, 0, wPriceData, hPriceData), palette.Plan9)}
}

//Make 价格的右上角坐标
func (p *PriceData) Make(m *image.Image, left, top int) {
	g := color.RGBA{0x0, 0xa8, 0x0, 0xff}
	r := color.RGBA{0xfc, 0x54, 0x54, 0xff}
	back := color.RGBA{0xff, 0xff, 0xff, 0}

	for x := 0; x < wPriceData; x++ {
		for y := 0; y < hPriceData; y++ {
			c := (*m).At(x+left, y+top)
			//p.img.Set(x, y, c)
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
func (p *PriceData) Save() {
	f, _ := os.Create(fmt.Sprintf("stockPriceData.gif"))
	defer f.Close()
	gif.Encode(f, p.img, nil)
}

func cutPriceData(m *image.Image) {

	p := CreatePriceData()
	p.Make(m, 50, 18)
	p.Save()
}
