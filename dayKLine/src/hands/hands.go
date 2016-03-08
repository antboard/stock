package hands

import (
	"charAndNum"
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
	// p.img.Set(0, 0, c)
	// fmt.Println(p.img.At(0, 0))
	// {153 153 153 255}

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

func (p *Hands) cover() {
	n := charAndNum.GetNum(p.img, 45, 0)
	n += charAndNum.GetNum(p.img, 39, 0) * 10
	n += charAndNum.GetNum(p.img, 33, 0) * 100
	n += charAndNum.GetNum(p.img, 27, 0) * 100
	n += charAndNum.GetNum(p.img, 21, 0) * 1000
	fmt.Print(n, ", ")
}

func cutHands(m *image.Image) {
	p := make([]*Hands, 4, 4)
	offy := []int{0, 21, 42, 63}
	for i := 0; i < 4; i++ {
		p[i] = CreateHands()
		p[i].Make(m, 45, 197+offy[i])
		p[i].Save(i)
	}
}

// GetHands ...
func GetHands(m *image.Image) {
	p := make([]*Hands, 4, 4)
	offy := []int{0, 21, 42, 63}
	for i := 0; i < 4; i++ {
		p[i] = CreateHands()
		p[i].Make(m, 45, 197+offy[i])
		// p[i].Save(i)
		p[i].cover()
	}
	fmt.Println("")
}
