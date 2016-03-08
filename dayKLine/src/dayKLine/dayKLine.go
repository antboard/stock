package main

import (
	"encoding/json"
	"fmt"
	"image/gif"
	"os"
	"price"
)

type aStock struct {
	Name, Code string
}
type stocks struct {
	St []*aStock
}

func deal(strFile string) {
	// 打开 gif 文件
	f, e := os.Open(strFile) // "/Users/jiangyichun/Downloads/code/stock/download/bin/sh601006.gif"
	if e != nil {
		fmt.Println("open err:", e)
		return
	}
	defer f.Close()
	//
	m, e := gif.Decode(f)
	if e != nil {
		fmt.Println("...Decode err:", e)
	}

	bounds := m.Bounds()
	if bounds.Max.X != 545 ||
		bounds.Max.Y != 300 {
		fmt.Println("...gif size is ", bounds, " NOT 545*300")
	}

	// 计算股票代码
	// 横坐标50-100处
	// 纵坐标7-15处
	price.GetPrices(&m)
	// hands.GetHands(&m)
	// date.GetDate(&m)
	// hands.GetData(&m)
	// price.GetPricesData(&m)
}

func main() {
	f, e := os.Open("/Users/jiangyichun/Downloads/code/stock/download/bin/allstock.json")
	if e != nil {
		fmt.Println(e)
	}
	length, e := f.Seek(0, os.SEEK_END)
	f.Seek(0, os.SEEK_SET)

	data := make([]byte, length, length)
	f.Read(data)
	// str := string(data)
	// fmt.Println(str)
	st := &stocks{}
	json.Unmarshal(data, st)
	fmt.Println("json len: ", len(st.St))
	for i := 0; i < len(st.St); i++ {
		strFile := "/Users/jiangyichun/Downloads/code/stock/download/bin/"
		// if st.St[i].Code[0] == '6' {
		// 	strFile += "sh"
		// } else {
		// 	strFile += "sz"
		// }
		strFile = strFile + st.St[i].Code + ".gif"
		fmt.Println(strFile)
		deal(strFile)
	}
}
