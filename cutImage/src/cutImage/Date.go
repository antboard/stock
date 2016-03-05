package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
)

// Date 价格
type Date struct {
	img *image.Paletted
}

// CreateDate ...
func CreateDate() *Date {
	return &Date{img: image.NewPaletted(image.Rect(0, 0, 59, 7), palette.Plan9)}
}

// MakeYear 价格的左上角坐标
func (p *Date) MakeYear(m *image.Image, left, top int) {
	c := color.RGBA{0x0, 0x86, 0xd2, 0xff}
	back := color.RGBA{0xff, 0xff, 0xff, 0}

	for x := 0; x < 30; x++ {
		for y := 0; y < 7; y++ {
			r, g, b, a := (*m).At(x+left, y+top).RGBA()

			if uint8(r) == 0 &&
				uint8(g) == 0x86 &&
				uint8(b) == 0xd2 &&
				uint8(a) == 0xff {
				p.img.Set(x, y, c)
			} else {
				p.img.Set(x, y, back)
			}
		}
	}
}

// MakeDate 价格的右上角坐标
func (p *Date) MakeDate(m *image.Image, left, top int) {
	c := color.RGBA{0x0, 0x86, 0xd2, 0xff}
	back := color.RGBA{0xff, 0xff, 0xff, 0}

	for x := 0; x < 29; x++ {
		for y := 0; y < 7; y++ {
			r, g, b, a := (*m).At(x+left, y+top).RGBA()

			if uint8(r) == 0 &&
				uint8(g) == 0x86 &&
				uint8(b) == 0xd2 &&
				uint8(a) == 0xff {
				p.img.Set(x+30, y, c)
			} else {
				p.img.Set(x+30, y, back)
			}
		}
	}
}

// Save ...
func (p *Date) Save(idx int) {
	w, _ := os.Create(fmt.Sprintf("stockDate%d.gif", idx))
	defer w.Close()
	gif.Encode(w, p.img, nil)
}

func cutDate(m *image.Image) {
	p := make([]*Date, 8, 8)
	offx := []int{0, 45, 105, 165, 234, 279, 339, 390}
	for i := 0; i < 8; i++ {
		p[i] = CreateDate()
		p[i].MakeDate(m, 104+offx[i], 183)
	}
	offyear := []int{0, 0, 0, 0, 234, 234, 234, 234}
	for i := 0; i < 8; i++ {
		p[i].MakeYear(m, 74+offyear[i], 183)
		p[i].Save(i)
	}
}