package price

import (
	"charAndNum"
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
	V   float32
}

// CreatePrice ...
func CreatePrice() *Price {
	return &Price{img: image.NewPaletted(image.Rect(0, 0, 50, 7), palette.WebSafe)}
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

// Cover ...
func (p *Price) Cover() {
	n := 0
	p.V = float32(0.0)
	b := float32(1.0) // 倍率
	n = charAndNum.GetNum(p.img, 45, 0)
	p.V += float32(n) * b
	b = b * 10
	n = charAndNum.GetNum(p.img, 39, 0)
	if n == 0xa {
		p.V = p.V * 0.1
		b = 1
	} else {
		p.V += float32(n) * b
		b = b * 10
	}

	n = charAndNum.GetNum(p.img, 33, 0)
	if n == 0xa {
		p.V = p.V * 0.01
		b = 1
	} else {
		p.V += float32(n) * b
		b = b * 10
	}
	n = charAndNum.GetNum(p.img, 27, 0)
	p.V += float32(n) * b
	b = b * 10

	n = charAndNum.GetNum(p.img, 21, 0)
	p.V += float32(n) * b
	b = b * 10
	// fmt.Print(p.V, " ")
	// charAndNum.GetNum(p.img, 20, 0)
}

// Save ...
func (p *Price) Save(idx int) {
	w, _ := os.Create(fmt.Sprintf("stockPrice%d.gif", idx))
	defer w.Close()
	gif.Encode(w, p.img, nil)
}

//CutPrice 切价格--用于测试图是否正确.
func CutPrice(m *image.Image) {
	p := make([]*Price, 10, 10)
	offy := []int{0, 17, 35, 53, 70, 88, 106, 123, 141, 159}
	for i := 0; i < 10; i++ {
		p[i] = CreatePrice()
		p[i].Make(m, 45, 15+offy[i])
		p[i].Save(i)
	}
}

// GetPrices 计算10个股价
func GetPrices(m *image.Image) (ret []float32) {
	ret = make([]float32, 10, 10)
	p := make([]*Price, 10, 10)
	offy := []int{0, 17, 35, 53, 70, 88, 106, 123, 141, 159}
	for i := 0; i < 10; i++ {
		p[i] = CreatePrice()
		p[i].Make(m, 45, 15+offy[i])
		p[i].Cover()

	}
	for i := 9; i >= 0; i-- {
		fmt.Print(p[i].V, ", ")
	}
	fmt.Println("")
	return
}
