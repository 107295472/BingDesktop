package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"encoding/json"
	"time"
	"github.com/tidwall/gjson"
	"strings"
	"net/http"
	"io/ioutil"
)
type BingUI struct {
	Img       string             `json:"img"`
	Idx string `json:"idx"`
	Ti string `json:"ti"`
}
func bingHandleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "bing":
		// Unmarshal payload
		var path string
		if len(m.Payload) > 0 {
		// Unmarshal payload
		if err = json.Unmarshal(m.Payload, &path); err != nil {
			payload = err.Error()
			return
		}
	}
		// Explore
		if payload, err = bingExplore(path); err != nil {
			payload = err.Error()
			return
		}

	case "Save":
	}
	return
}
func bingExplore(path string) (e BingUI, err error) {
	//astilog.Info(path)
	a,b:=Next(path)
	e.Ti=a
	e.Img=b
	e.Idx=path
	return
}
func Next(idx string)(string,string){
	var htt="http://cn.bing.com"
	var url="http://cn.bing.com/HPImageArchive.aspx?format=js&idx=" + idx + "&n=1&nc=" + string(time.Now().Unix()) + "&video=1";
	json:=httpGet(url)
	img := gjson.Get(json, "images.#.url")
	title := gjson.Get(json, "images.#.copyright")
	var title2 string
	var img1 string
	for _, name := range title.Array() {
		title2=name.String()
	}
	for _, name := range img.Array() {
		img1=name.String()
	}
	//截取字符串
	str1:=strings.Index(title2, "©")
	var s= string([]byte(title2)[:str1-1])
	return s,htt+img1
}
func httpGet(url string) string{
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body)
}