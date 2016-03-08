package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type aStock struct {
	Name, Code string
}
type stocks struct {
	St []*aStock
}

// Get ...
func Get(strURL, Code string) {
	out, err := os.Create(Code + ".gif")
	defer out.Close()
	resp, err := http.Get(strURL)
	if err != nil {
		fmt.Println("...Get error ", err)
		return
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	_, err = io.Copy(out, bytes.NewReader(pix))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// 打开文件
	f, e := os.Open("stocklist.txt")
	if e != nil {
		fmt.Println("...open...", e)
	}
	defer f.Close()
	length, e := f.Seek(0, os.SEEK_END)
	f.Seek(0, os.SEEK_SET)
	data := make([]byte, length, length)
	n, err := f.Read(data)
	if err != nil {
		fmt.Println("...read...", err, n)
	}
	s := string(data)
	//fmt.Println(s, data)
	kv := strings.Split(s, "\n")
	fmt.Println(len(kv))

	allstock := stocks{St: make([]*aStock, len(kv), len(kv))}
	for i := 0; i < len(kv); i++ {
		stock := kv[i]
		stockNd := strings.Split(stock, " ")
		if len(stockNd) > 2 {
			for i := 1; i < len(stockNd)-1; i++ {
				stockNd[0] += stockNd[i]
			}
			stockNd[1] = stockNd[len(stockNd)-1]
		}
		// fmt.Println(stockNd[0], ":", stockNd[1])
		allstock.St[i] = &aStock{stockNd[0], stockNd[1]}
	}
	b, e := json.Marshal(allstock)
	if e != nil {
		fmt.Println("...json :", e)
	}
	fmt.Println("json len:", len(b))

	out, err := os.Create("allstock.json")
	defer out.Close()
	out.Write(b)
	for i := 0; i < len(allstock.St); i++ {

		strURL := "http://image.sinajs.cn/newchart/daily/n/"
		if allstock.St[i].Code[0] == '6' {
			strURL += "sh"
			strURL += allstock.St[i].Code
			strURL += ".gif"
		} else {
			strURL += "sz"
			strURL += allstock.St[i].Code
			strURL += ".gif"
		}
		if len(strURL) > 60 {
			return
		}
		fmt.Println("get:", strURL)
		Get(strURL, allstock.St[i].Code)
		time.Sleep(time.Second / 2)
		if i%10 == 0 {
			fmt.Println(i, ":")
		}
	}

}
