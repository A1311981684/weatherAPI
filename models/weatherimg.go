package models

import (
	"log"
	"io/ioutil"
	"net/http"
	"golang.org/x/net/html"
	"weatherAPI/const"
	"strings"
	"github.com/pkg/errors"
	"fmt"
)

type WeatherImage struct {
	Images []byte
}

var ImageURIList []string
var ImageNumbers int

//从一个图片网址获取一张图片，转换为[]byte并返回它
func GetImage(url string) (img *WeatherImage){
	//获取网址响应数据
	resp, err:=http.Get(url)
	defer resp.Body.Close()
	if err!=nil{
		log.Fatal(err)
	}
	//读取数据
	pic, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//获取图片
	img = new(WeatherImage)
	img.Images = pic
	return img
}

//获取一张图片并返回
func GetOneImage(id int, toSize string) (img *WeatherImage, err error){
	theURL ,err:= ChangeImgSize(ImageURIList[id], toSize)
	fmt.Println("Getting URL ::::::::::",theURL)
	resp, err:=http.Get(theURL)
	if err !=nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	if err!=nil{
		log.Fatal(err)
		return
	}
	//读取数据
	pic, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	//获取图片
	img = new(WeatherImage)
	img.Images = pic
	return img, nil
}
//TODO:Generate a slice of Images' URL
func getIMGLinks(url string) ([]string, int) {
	//存放所有链接
	var res []string
	var count int
	//请求网页
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//将响应解析成html
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//处理链接所在的节点
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "data-original" {
					//向结果数组添加图像链接
					res = append(res, a.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	count = len(res)
	//返回链接集合，链接数量
	return res, count
}

//获取所有的可用土图片
func GetAllLinks() []string{
	ImageURIList, size := getIMGLinks(_const.MainURL)
	ImageNumbers = size
	return ImageURIList
}

//修改图片尺寸by修改url参数
func ChangeImgSize(url string, changeTo string) (rtURL string, err error){
	splitedURL := strings.Split(url, "/WXCL")
	currentSize := ""
	//URL 不合法
	if len(splitedURL) != 2 {
		return url, errors.New("Split URL failed:"+url)
	}
	//判断当前图片规格
	if strings.Contains(url, _const.SmallSize){
		currentSize = _const.SmallSize
	}else if strings.Contains(url, _const.MediumSize) {
		currentSize = _const.MediumSize
	}else {
		currentSize = _const.LargeSize
	}

	if currentSize == "" {
		return "", errors.New("Invalid Size parameter!")
	}
	//判断转换情况:先判断目标尺寸，再判断当前的尺寸
	switch changeTo{
	case _const.LargeSize://转为Large
		switch currentSize {
			case _const.LargeSize:
				return url, nil
			case _const.MediumSize:
				newURL := splitedURL[0] + "/WXCL" + strings.Replace(splitedURL[1], "/medium","",1)
				return newURL, nil
			case _const.SmallSize:
				newURL := splitedURL[0] + "/WXCL" + strings.Replace(splitedURL[1], "/small","",1)
				return newURL, nil
		}
	case _const.MediumSize://转为Medium
		switch currentSize {
			case _const.MediumSize:
				return url, nil
			case _const.SmallSize:
				newURL := splitedURL[0] + "/WXCL" + strings.Replace(splitedURL[1], _const.SmallSize, _const.MediumSize,1)
				return newURL, nil
			case _const.LargeSize:
				newURL := splitedURL[0] + "/WXCL/medium" + splitedURL[1]
				return newURL, nil
		}
	case _const.SmallSize://转为Small
		switch currentSize {
			case _const.SmallSize:
				return url, nil
			case _const.MediumSize:
				newURL := splitedURL[0] + "/WXCL" + strings.Replace(splitedURL[1], _const.MediumSize, _const.SmallSize,1)
				return newURL, nil
			case _const.LargeSize:
				newURL := splitedURL[0] + "/WXCL/small" + splitedURL[1]
				return newURL, nil
		}
	default:
		return url, errors.New("Invalid Size parameter: " + changeTo)
	}
	return url,errors.New("No case matched...")
}

