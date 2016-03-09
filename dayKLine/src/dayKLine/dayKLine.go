package main

import (
	"date"
	"encoding/json"
	"fmt"
	"hands"
	"image/gif"
	"os"
	"price"
	"runtime"
	"time"
)

type aStock struct {
	Name, Code string
}
type stocks struct {
	St []*aStock
}

var wait chan int
var files chan int

func deal(strFile string) {
	// 打开 gif 文件
	f, e := os.Open(strFile) // "/Users/jiangyichun/Downloads/code/stock/download/bin/sh601006.gif"
	if e != nil {
		fmt.Println("open err:", e)
		files <- 1
		wait <- 1
		return
	}
	// defer f.Close()
	//
	m, e := gif.Decode(f)
	if e != nil {
		fmt.Println("...Decode err:", e)
		files <- 1
		wait <- 1
		f.Close()
		return
	}

	f.Close()

	bounds := m.Bounds()
	if bounds.Max.X != 545 ||
		bounds.Max.Y != 300 {
		fmt.Println("err: ...gif size is ", bounds, " NOT 545*300")
	}

	// 计算股票代码
	// 横坐标50-100处
	// 纵坐标7-15处
	priceNum := price.GetPrices(&m)
	// price.GetPricesData(&m)
	handsNum := hands.GetHands(&m)
	// hands.GetData(&m)
	dateNum := date.GetDate(&m)

	s := fmt.Sprintln(strFile)
	for i := 0; i < len(priceNum); i++ {
		if priceNum[i] < 0.01 {
			s += "err: price err\n"
		}
	}
	s += fmt.Sprintln(priceNum)
	s += fmt.Sprintln(handsNum)
	s += fmt.Sprintln(dateNum)
	fmt.Println(s)
	files <- 1
	wait <- 1

}

func oldmain() {
	wait = make(chan int, 10)
	files = make(chan int, 100)
	for i := 0; i < 100; i++ {
		files <- i
	}
	MULTICORE := runtime.NumCPU() //number of core

	runtime.GOMAXPROCS(MULTICORE) //running in multicore
	fmt.Println(MULTICORE)

	f, e := os.Open("/Users/jiangyichun/Downloads/code/stock/download/bin/allstock.json")
	if e != nil {
		fmt.Println("err", e)
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
	for i := 0; i < len(st.St); i++ { //len(st.St)
		<-files
		strFile := "/Users/jiangyichun/Downloads/code/stock/download/bin/"
		// if st.St[i].Code[0] == '6' {
		// 	strFile += "sh"
		// } else {
		// 	strFile += "sz"
		// }
		strFile = strFile + st.St[i].Code + ".gif"
		// fmt.Println(strFile)
		go deal(strFile)
		// deal(strFile)
	}
	for i := 0; i < len(st.St); i++ {
		<-wait
		// fmt.Println("...", i, "end")
	}
}

func main() {
	wait = make(chan int, 10)
	files = make(chan int, 100)
	deal("/Users/jiangyichun/Downloads/code/stock/download/bin/600156.gif")
	time.Sleep(time.Second * 3)
}
