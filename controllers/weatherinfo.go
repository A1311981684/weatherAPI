package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"weatherAPI/const"
	"weatherAPI/models"
)

//Operation about Weather Information
type WeatherInfoController struct{
	beego.Controller
}

// @Title GetAll
// @Description Get All The Data Can Be Retrieved From The URL
// @Param city path string true "The City Needed To Be Checked."
// @Success 200 {object} models.PrimaryData
// @Failure 403 :city invalid
// @router /:city [get]
func (w *WeatherInfoController) GetAllInfo() {
	city := w.Ctx.Input.Param(":city")
	pd ,err := models.GetTotalJSON(_const.WeatherURL + city)
	if err != nil {
		fmt.Println(err)
		w.Ctx.WriteString(err.Error())
	}else{
		w.Data["json"] = &pd
		w.ServeJSON()
	}
}

//@Title GetToday
//@Description Get Today's Weather information
//@Param   city path string true "Get specific city's current day weather"
//@Success 200 {object} models.Today
//@Failure 403 :city invalid
//@router  /weather_info/today/:city [get]
func (w *WeatherInfoController) GetToday(){
	city := w.Ctx.Input.Param(":city")
	var td models.Today

	pd, err := models.GetTotalJSON(_const.WeatherURL + city)
	if err != nil {
		log.Fatal(err)
	}

	td.Humidity = pd.Data.Humidity
	td.Pm10 = pd.Data.Pm10
	td.Pm25 = pd.Data.Pm25
	td.Quality = pd.Data.Quality
	td.Temperature = pd.Data.Temperature
	td.CatchCold = pd.Data.CatchCold
	td.Someday = *pd.Data.Forecast[0]

	w.Data["json"] = td
	w.ServeJSON()
}

//@Title GetAllForecast
//@Description Get all forecast info except for TODAY
//@Param city path string true "Get Specific city's forecast"
//@Success 200 {object} []models.SomeDay
//@Failure 403 :city invalid
//@router /weather_info/forecast/:city [get]
func (w *WeatherInfoController) GetAllForecast(){
	city := w.Ctx.Input.Param(":city")
	pd, err := models.GetTotalJSON(_const.WeatherURL + city)
	if err != nil {
		log.Fatal(err)
	}
	var fc []*models.SomeDay
	fc = pd.Data.Forecast[1:]
	w.Data["json"] = fc
	w.ServeJSON()
}