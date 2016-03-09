package charAndNum

import (
	"fmt"
	"image"
	"image/color"
)

// 本文件用于文字识别
// 1. 假定背景为白色
// 3. 字模的比较:字体像素与白色比较.如果是就是1,不是就0.

// Dump ...
func Dump(x [][]int8) {
	for j := 0; j < 7; j++ {
		for i := 0; i < 5; i++ {
			if x[j][i] == 0 {
				fmt.Print("0")
			} else {
				// fmt.Print(" ")
				fmt.Print("1")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

// CMP ...
func CMP(f, src [][]int8) bool {
	for j := 0; j < 7; j++ {
		for i := 0; i < 5; i++ {
			if src[j][i] == 0 && f[j][i] != src[j][i] {
				return false
			}
		}
	}
	return true
}

// CMPA ...
func CMPA(f, src [][]int8) bool {
	for j := 0; j < 7; j++ {
		for i := 0; i < 5; i++ {
			if f[j][i] != src[j][i] {
				return false
			}
		}
	}
	return true
}

// 字库
var (
	NUM0 = [][]int8{
		{1, 0, 0, 0, 1}, // 1
		{0, 1, 1, 1, 0}, // 2
		{0, 1, 1, 0, 0}, // 3
		{0, 1, 0, 1, 0}, // 4
		{0, 0, 1, 1, 0}, // 5
		{0, 1, 1, 1, 0}, // 6
		{1, 0, 0, 0, 1}, // 7
	}
	NUM1 = [][]int8{
		{1, 1, 0, 1, 1}, // 1
		{1, 0, 0, 1, 1}, // 2
		{1, 1, 0, 1, 1}, // 3
		{1, 1, 0, 1, 1}, // 4
		{1, 1, 0, 1, 1}, // 5
		{1, 1, 0, 1, 1}, // 6
		{1, 0, 0, 0, 1}, // 7
	}
	NUM2 = [][]int8{
		{1, 0, 0, 0, 1}, // 1
		{0, 1, 1, 1, 0}, // 2
		{1, 1, 1, 1, 0}, // 3
		{1, 1, 1, 0, 1}, // 4
		{1, 1, 0, 1, 1}, // 5
		{1, 0, 1, 1, 1}, // 6
		{0, 0, 0, 0, 0}, // 7
	}
	NUM3 = [][]int8{
		{1, 0, 0, 0, 1}, // 1
		{0, 1, 1, 1, 0}, // 2
		{1, 1, 1, 1, 0}, // 3
		{1, 1, 0, 0, 1}, // 4
		{1, 1, 1, 1, 0}, // 5
		{0, 1, 1, 1, 0}, // 6
		{1, 0, 0, 0, 1}, // 7
	}
	NUM4 = [][]int8{
		{1, 1, 1, 0, 1}, // 1
		{1, 1, 0, 0, 1}, // 2
		{1, 0, 1, 0, 1}, // 3
		{0, 1, 1, 0, 1}, // 4
		{0, 0, 0, 0, 0}, // 5
		{1, 1, 1, 0, 1}, // 6
		{1, 1, 1, 0, 1}, // 7
	}
	NUM5 = [][]int8{
		{0, 0, 0, 0, 0}, // 1
		{0, 1, 1, 1, 1}, // 2
		{0, 0, 0, 0, 1}, // 3
		{1, 1, 1, 1, 0}, // 4
		{1, 1, 1, 1, 0}, // 5
		{0, 1, 1, 1, 0}, // 6
		{1, 0, 0, 0, 1}, // 7
	}
	NUM6 = [][]int8{
		{1, 1, 0, 0, 1}, // 1
		{1, 0, 1, 1, 1}, // 2
		{0, 1, 1, 1, 1}, // 3
		{0, 0, 0, 0, 1}, // 4
		{0, 1, 1, 1, 0}, // 5
		{0, 1, 1, 1, 0}, // 6
		{1, 0, 0, 0, 1}, // 7
	}
	NUM7 = [][]int8{
		{0, 0, 0, 0, 0}, // 1
		{1, 1, 1, 1, 0}, // 2
		{1, 1, 1, 0, 1}, // 3
		{1, 1, 1, 0, 1}, // 4
		{1, 1, 0, 1, 1}, // 5
		{1, 1, 0, 1, 1}, // 6
		{1, 1, 0, 1, 1}, // 7
	}
	NUM8 = [][]int8{
		{1, 0, 0, 0, 1}, // 1
		{0, 1, 1, 1, 0}, // 2
		{0, 1, 1, 1, 0}, // 3
		{1, 0, 0, 0, 1}, // 4
		{0, 1, 1, 1, 0}, // 5
		{0, 1, 1, 1, 0}, // 6
		{1, 0, 0, 0, 1}, // 7
	}
	NUM9 = [][]int8{
		{1, 0, 0, 0, 1}, // 1
		{0, 1, 1, 1, 0}, // 2
		{0, 1, 1, 1, 0}, // 3
		{1, 0, 0, 0, 0}, // 4
		{1, 1, 1, 1, 0}, // 5
		{1, 1, 1, 0, 1}, // 6
		{1, 0, 0, 1, 1}, // 7
	}
	NUMdot = [][]int8{
		{1, 1, 1, 1, 1}, // 1
		{1, 1, 1, 1, 1}, // 2
		{1, 1, 1, 1, 1}, // 3
		{1, 1, 1, 1, 1}, // 4
		{1, 1, 1, 1, 1}, // 5
		{1, 0, 0, 1, 1}, // 6
		{1, 0, 0, 1, 1}, // 7
	}
	CHARs = [][]int8{
		{1, 0, 0, 0, 1}, // 1
		{0, 1, 1, 1, 0}, // 2
		{1, 0, 0, 1, 1}, // 3
		{1, 1, 1, 0, 1}, // 4
		{1, 1, 1, 1, 0}, // 5
		{1, 1, 1, 1, 1}, // 6
		{1, 1, 1, 1, 1}, // 7
	}
	CHARdiv = [][]int8{
		{1, 1, 1, 1, 1}, // 1
		{1, 1, 1, 1, 1}, // 2
		{1, 1, 1, 1, 1}, // 3
		{0, 0, 0, 0, 0}, // 4
		{1, 1, 1, 1, 1}, // 5
		{1, 1, 1, 1, 1}, // 6
		{1, 1, 1, 1, 1}, // 7
	}
	NUM = [][]int8{
		{1, 1, 1, 1, 1}, // 1
		{1, 1, 1, 1, 1}, // 2
		{1, 1, 1, 1, 1}, // 3
		{1, 1, 1, 1, 1}, // 4
		{1, 1, 1, 1, 1}, // 5
		{1, 1, 1, 1, 1}, // 6
		{1, 1, 1, 1, 1}, // 7
	}
	CHARz = [][]int8{
		{1, 1, 1, 1, 1}, // 1
		{1, 1, 1, 1, 1}, // 2
		{1, 1, 1, 1, 1}, // 3
		{1, 1, 1, 1, 1}, // 4
		{1, 1, 1, 1, 1}, // 5
		{1, 1, 1, 1, 1}, // 6
		{1, 1, 1, 1, 1}, // 7
	}
)

// GetNum 从 x,y 位置开始计算一个字符
func GetNum(m *image.Paletted, x, y int) int {
	c := GetChar(m, x, y, true)
	if c == ' ' {
		return 0
	}
	if c == 'p' {
		return -1
	}
	n := c - '0'
	if n > 9 || n < 0 {
		fmt.Println("...GetNum error...", string(c), n)
	}
	return int(n)
}

// GetChar 从 x,y 位置开始计算一个字符
func GetChar(m *image.Paletted, x, y int, b bool) (n byte) {
	// 计算字模
	// 背景色
	num := [][]int8{
		{1, 1, 1, 1, 1}, // 1
		{1, 1, 1, 1, 1}, // 2
		{1, 1, 1, 1, 1}, // 3
		{1, 1, 1, 1, 1}, // 4
		{1, 1, 1, 1, 1}, // 5
		{1, 1, 1, 1, 1}, // 6
		{1, 1, 1, 1, 1}, // 7
	}
	back := color.RGBA{0xff, 0xff, 0xff, 0xff}
	for i := 0; i < 7; i++ {
		for j := 0; j < 5; j++ {
			if m.At(j+x, i+y) != back {
				num[i][j] = 0
			}
		}
	}

	n = '0'
	if CMP(num, NUM0) {
		n = '0'
	} else if CMP(num, NUM2) {
		n = '2'
	} else if CMP(num, NUM1) {
		n = '1'
	} else if CMP(num, NUM4) {
		n = '4'
	} else if CMP(num, NUM5) {
		n = '5'
	} else if CMP(num, NUM6) {
		n = '6'
	} else if CMP(num, NUM7) {
		n = '7'
	} else if CMP(num, NUM8) {
		n = '8'
	} else if CMP(num, NUM9) {
		n = '9'
	} else if CMP(num, NUM3) {
		n = '3'
	} else if CMP(num, NUMdot) {
		n = '.'
	} else if CMP(num, CHARdiv) {
		n = '-'
	} else if CMPA(num, NUM) {
		n = ' '
	} else {
		if b {
			n = 'p'
			fmt.Println("err: x = ", x, "y = ", y)
			Dump(num)
		}
	}
	// fmt.Println("bit", string(n))
	return n
}

// GetCharA 从 x,y 位置开始计算一个字符
func GetCharA(m *image.Paletted, x, y int, b bool) (n byte) {
	// 计算字模
	// 背景色
	num := [][]int8{
		{1, 1, 1, 1, 1}, // 1
		{1, 1, 1, 1, 1}, // 2
		{1, 1, 1, 1, 1}, // 3
		{1, 1, 1, 1, 1}, // 4
		{1, 1, 1, 1, 1}, // 5
		{1, 1, 1, 1, 1}, // 6
		{1, 1, 1, 1, 1}, // 7
	}
	back := color.RGBA{0xff, 0xff, 0xff, 0xff}
	for i := 0; i < 7; i++ {
		for j := 0; j < 5; j++ {
			if m.At(j+x, i+y) != back {
				num[i][j] = 0
			}
		}
	}

	n = '0'
	if CMPA(num, NUM0) {
		n = '0'
	} else if CMPA(num, NUM2) {
		n = '2'
	} else if CMPA(num, NUM1) {
		n = '1'
	} else if CMPA(num, NUM4) {
		n = '4'
	} else if CMPA(num, NUM5) {
		n = '5'
	} else if CMPA(num, NUM6) {
		n = '6'
	} else if CMPA(num, NUM7) {
		n = '7'
	} else if CMPA(num, NUM8) {
		n = '8'
	} else if CMPA(num, NUM9) {
		n = '9'
	} else if CMPA(num, NUM3) {
		n = '3'
	} else if CMPA(num, NUMdot) {
		n = '.'
	} else if CMPA(num, CHARdiv) {
		n = '-'
	} else if CMPA(num, NUM) {
		n = ' '
	} else {
		if b {
			n = 'p'
			fmt.Println("err: x = ", x, "y = ", y)
			Dump(num)
		}
	}
	// fmt.Println("bit", string(n))
	return n
}
