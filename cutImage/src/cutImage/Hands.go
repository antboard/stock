package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"os"
)

// Hands 价格
type Hands struct {
	img *image.Paletted
}

// CreateHands ...
func CreateHands() *Hands {
	return &Hands{img: image.NewPaletted(image.Rect(0, 0, 50, 7), palette.Plan9)}
}

//Make 价格的右上角坐标
func (p *Hands) Make(m *image.Image, right, top int) {
	c := color.RGBA{0x9b, 0x9b, 0x9b, 0xff}
	back := color.RGBA{0xff, 0xff, 0xff, 0}

	offx := right - 50 + 1
	for x := 0; x < 50; x++ {
		for y := 0; y < 7; y++ {
			r, g, b, a := (*m).At(x+offx, y+top).RGBA()

			if uint8(r) == 0x9b &&
				uint8(g) == 0x9b &&
				uint8(b) == 0x9b &&
				uint8(a) == 0xff {
				p.img.Set(x, y, c)
			} else {
				p.img.Set(x, y, back)
			}
		}
	}
}

// Save ...
func (p *Hands) Save(idx int) {
	w, _ := os.Create(fmt.Sprintf("stockHands%d.gif", idx))
	defer w.Close()
	gif.Encode(w, p.img, nil)
}

func cutHands(m *image.Image) {
	p := make([]*Hands, 0, 4)
	back := color.RGBA{0, 0, 0, 0}
	// offy := []int{0, 21, 42, 63}
	last := 197
	// for i := 0; i < 45; i++ {
	// 	fmt.Println((*m).At(i, last))
	// }

	for i := 0; ; i++ {
	Find:
		for ; last < 270; last++ {
			for k := 0; k < 45; k++ {
				if (*m).At(k, last) != back {
					break Find
				}
			}
		}
		if last >= 270 {
			break
		}
		p = append(p, CreateHands())
		fmt.Println("...", last)
		p[i].Make(m, 45, last)
		last += 8
		p[i].Save(i)
	}
}
