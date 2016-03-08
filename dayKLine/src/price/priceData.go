package price

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

type rc struct {
	BlockHigh, BlockLow int
	LineHigh, LineLow   int
}

// Data 价格
type Data struct {
	img *image.Paletted
	v   []*rc
}

// CreateData ...
func CreateData() *Data {
	return &Data{img: image.NewPaletted(image.Rect(0, 0, wPriceData, hPriceData), palette.Plan9),
		v: make([]*rc, 80, 80)}
}

//Make 价格的右上角坐标
func (p *Data) Make(m *image.Image, left, top int) {
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

func (p *Data) getLineHigh(x int) (cnt int) {
	back := color.RGBA{0xff, 0xff, 0xff, 0xff}
	// fmt.Println(x, "---", p.img.At(x, 0))
	cnt = 0
	j := 2
	for i := 0; i < hPriceData; i++ {
		if p.img.At(j+x, i) != back {
			// fmt.Println(j, i, "---", p.img.At(j+x, i))
			cnt = hPriceData - i
			return
		}
	}
	return
}

func (p *Data) getLineLow(x int) (cnt int) {
	back := color.RGBA{0xff, 0xff, 0xff, 0xff}
	cnt = 0
	j := 2
	for i := hPriceData - 1; i >= 0; i-- {
		if p.img.At(j+x, i) != back {
			// fmt.Println(j, i, "---", p.img.At(j+x, i))
			cnt = hPriceData - i
			return
		}
	}
	return
}

func (p *Data) getBlockHigh(x int) (cnt int) {
	back := color.RGBA{0xff, 0xff, 0xff, 0xff}
	// fmt.Println(x, "---", p.img.At(x, 0))
	cnt = 0
	for i := 0; i < hPriceData; i++ {
		if p.img.At(0+x, i) != back ||
			p.img.At(1+x, i) != back ||
			p.img.At(3+x, i) != back ||
			p.img.At(4+x, i) != back {
			// fmt.Println(j, i, "---", p.img.At(j+x, i))
			cnt = hPriceData - i
			return
		}
	}
	return
}

func (p *Data) getBlockLow(x int) (cnt int) {
	back := color.RGBA{0xff, 0xff, 0xff, 0xff}
	// fmt.Println(x, "---", p.img.At(x, 0))
	cnt = 0
	for i := hPriceData - 1; i >= 0; i-- {
		if p.img.At(0+x, i) != back ||
			p.img.At(1+x, i) != back ||
			p.img.At(3+x, i) != back ||
			p.img.At(4+x, i) != back {
			// fmt.Println(j, i, "---", p.img.At(j+x, i))
			cnt = hPriceData - i
			return
		}
	}
	return
}

func (p *Data) getValue() {
	for i := 0; i < 80; i++ {
		x := i * 6
		BlockHigh := p.getBlockHigh(x)
		BlockLow := p.getBlockLow(x)
		LineHigh := p.getLineHigh(x)
		LineLow := p.getLineLow(x)
		// 块被盖掉
		if BlockHigh == 0 {
			BlockHigh = LineHigh
		}
		if BlockLow == 0 {
			BlockLow = LineLow
		}
		// 线被盖掉
		if LineLow == 0 {
			LineLow = BlockLow
		}
		if LineHigh == 0 {
			LineHigh = BlockHigh
		}

		if BlockHigh > LineHigh {
			LineHigh = BlockHigh
		}
		if BlockLow < LineLow {
			LineLow = BlockLow
		}
		p.v[i] = &rc{BlockHigh, BlockLow, LineHigh, LineLow}
		fmt.Print(p.v[i], ", ")
	}
	fmt.Println("")
}

func (p *Data) drawValue() {
	c := color.RGBA{0, 0, 0, 0xff}
	for i := 0; i < 80; i++ {
		r := p.v[i]
		// 画框 竖
		for j := hPriceData - r.BlockHigh; j < hPriceData-r.BlockLow; j++ {
			p.img.Set(i*6, j, c)
			p.img.Set(i*6+5, j, c)
		}
		// 画框 横
		for k := 0; k < 5; k++ {
			p.img.Set(k+i*6, hPriceData-r.BlockLow, c)
			p.img.Set(k+i*6, hPriceData-r.BlockHigh, c)
		}
		// 画中线
		for k := hPriceData - r.LineHigh; k <= hPriceData-r.LineLow; k++ {
			p.img.Set(i*6+3, k, c)
		}
	}
}

// Save ...
func (p *Data) Save() {
	f, _ := os.Create(fmt.Sprintf("stockData.gif"))
	defer f.Close()
	gif.Encode(f, p.img, nil)
}

func cutData(m *image.Image) {
	p := CreateData()
	p.Make(m, 50, 18)
	p.Save()
}

// GetPricesData ...
func GetPricesData(m *image.Image) {
	p := CreateData()
	p.Make(m, 50, 18)
	p.getValue()
	p.drawValue()
	p.Save()
}
