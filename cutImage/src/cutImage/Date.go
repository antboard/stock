package main

import (
	"charAndNum"
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

// MakeDate 价格的右左上角坐标
func (p *Date) MakeDate(m *image.Image, left, top int) int {
	c := color.RGBA{0x0, 0x86, 0xd2, 0xff}
	back := color.RGBA{0xff, 0xff, 0xff, 0}

	length := 59
	for x := 0; x < length; x++ {
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

	bClear := false
	if charAndNum.GetChar(p.img, 27, 0, false) == '-' {
		fmt.Println("...---...")
		bClear = true
	}

	if bClear {
		for x := 0; x < 15; x++ {
			for y := 0; y < 7; y++ {
				p.img.Set(x, y, back)
			}
		}
		for x := 45; x < 59; x++ {
			for y := 0; y < 7; y++ {
				p.img.Set(x, y, back)
			}
		}
	}
	return length
}

// Save ...
func (p *Date) Save(idx int) {
	w, _ := os.Create(fmt.Sprintf("stockDate%d.gif", idx))
	defer w.Close()
	gif.Encode(w, p.img, nil)
}

func cutDate(m *image.Image) {
	p := make([]*Date, 0, 16)
	back := color.RGBA{0xb4, 0xb4, 0xb4, 0xff}

	// offx := []int{0, 45, 105, 165, 234, 279, 339, 390}
	left := 55
	for i := 0; ; i++ {
	Find:
		for ; left < 540; left++ {
			for j := 20; j < 30; j++ {
				pt := (*m).At(left, j)
				// fmt.Println(pt)
				if pt == back {
					fmt.Println("...", left)
					break Find
				}
			}
		}

		if left >= 540 {
			break
		}

		p = append(p, CreateDate())
		p[i].MakeDate(m, left-30, 183)
		left++
		p[i].Save(i)
	}
}

// func cutDate(m *image.Image) {
// 	p := make([]*Date, 0, 16)
// 	back := color.RGBA{0, 0, 0, 0}

// 	// offx := []int{0, 45, 105, 165, 234, 279, 339, 390}
// 	last := 40
// 	for i := 0; ; i++ {
// 	Find:
// 		for ; last < 500; last++ {
// 			for i := 0; i < 7; i++ {
// 				if (*m).At(last, i+183) != back {
// 					// fmt.Println("...", last)
// 					break Find
// 				}
// 			}
// 		}
// 		if last > 500 {
// 			break
// 		}

// 		p = append(p, CreateDate())
// 		last += p[i].MakeDate(m, last, 183)
// 		p[i].Save(i)
// 	}
// }
