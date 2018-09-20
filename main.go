package main

import (
	_ "weatherAPI/routers"

	"github.com/astaxie/beego"
	"weatherAPI/models"
	"fmt"
	)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//初始化图片数据
	models.ImageURIList = models.GetAllLinks()
	if models.ImageNumbers == 12 {
		for _, x := range models.ImageURIList{
			fmt.Println(x)
		}
		fmt.Println("********** Images Fetched **********")
		beego.Run()
	}else {
		fmt.Println("Failed ,Image Number is not correct:", len(models.ImageURIList))
		fmt.Println("Service not Running.")
	}

}
