package controllers

import (
	"weatherAPI/models"
	"github.com/astaxie/beego"
	"strconv"
	"log"
	"fmt"
	"math/rand"
)

// Operations about weather
type WeatherImgController struct {
	beego.Controller
}

// @Title Test for Getting a IMG
// @Description  Randomly get an image
// @Success 200 {object} models.WeatherImage
// @Failure 403 :object is empty
// @router / [get]
func (w *WeatherImgController) Get() {
	x := rand.Intn(11) //产生随机数
	img := models.GetImage(models.ImageURIList[x])//获取制定图片
	w.Ctx.ResponseWriter.Write(img.Images)
}

// @Title Get One specific IMG
// @Description get IMG by id
// @Param	id	    path 	string 	true	"The key for a IMG"
// @Param	size	path 	string 	true	"The key for Image Size"
// @Success 200 {object} models.WeatherImage
// @Failure 403 :id or size invalid
// @router /:id/:size [get]
func (w *WeatherImgController) GetOneIMG() {
	//从path获取参数
	paramID := w.Ctx.Input.Param(":id")
	paramSIZE := w.Ctx.Input.Param(":size")

	if paramID != "" && paramSIZE != "" {
		id, err := strconv.Atoi(paramID)
		if err != nil {
			fmt.Println(paramID)
			log.Fatal(err)
		}
		IMG, err := models.GetOneImage(id, paramSIZE)
		if err == nil {
			w.Ctx.ResponseWriter.Write(IMG.Images)
			fmt.Println(paramID, paramSIZE)
		} else {
			log.Fatal(err)
		}
	} else {
		w.Ctx.ResponseWriter.Write([]byte("Invalid or nil Parameter!" + paramID + " " + paramSIZE))
	}

}
