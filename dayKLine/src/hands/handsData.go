package hands

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

// Data 手数
type Data struct {
	img *image.Paletted
}

// CreateData ...
func CreateData() *Data {
	return &Data{img: image.NewPaletted(image.Rect(0, 0, w, h), palette.Plan9)}
}

// MakeData 价格的右上角坐标
func (p *Data) MakeData(m *image.Image, left, top int) {
	g := color.RGBA{0x0, 0xa8, 0x0, 0xff}
	r := color.RGBA{0xfc, 0x54, 0x54, 0xff}
	back := color.RGBA{0xff, 0xff, 0xff, 0xff}

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
func (p *Data) Save() {
	f, _ := os.Create(fmt.Sprintf("stockData.gif"))
	defer f.Close()
	gif.Encode(f, p.img, nil)
}

func (p *Data) getValue(x int) (cnt int) {
	back := color.RGBA{0xff, 0xff, 0xff, 0xff}
	// fmt.Println(x, "---", p.img.At(x, 0))
	cnt = 0
	for i := 0; i < h; i++ {
		for j := 0; j < 5; j++ {
			if p.img.At(j+x, i) != back {
				// fmt.Println(j, i, "---", p.img.At(j+x, i))
				cnt = h - i
				return
			}
		}
	}
	return
}

func cutData(m *image.Image) {
	p := CreateData()
	p.MakeData(m, 50, 201)
	p.Save()
}

// GetData ...
func GetData(m *image.Image) {
	p := CreateData()
	p.MakeData(m, 50, 201)
	day := w / 6
	v := make([]int, day, day)
	for i := 0; i < day; i++ {
		v[i] = p.getValue(i * 6)
		fmt.Print(v[i], ", ")
	}
	// 检测数据
	// for i := 0; i < day; i++ {
	// 	p.img.Set(i*6+1, 84-v[i], color.RGBA{0, 0, 0xff, 0xff})
	// }
	// p.Save()
	// 计算 ma5
	// k5 := make([]int, day, day)
	// k5[0] = 0
	// k5[1] = 1
	// for i := 2; i < day-2; i++ {
	// 	k5[i] = (v[i-2] + v[i-1] + v[i-0] + v[i+1] + v[i+2]) / 5
	// 	p.img.Set(i*6+2, 84-k5[i], color.RGBA{0x11, 0x11, 0xff, 0xff})
	// }
	// p.Save()
	fmt.Println("")
}
