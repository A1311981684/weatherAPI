// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"weatherAPI/controllers"
)

func init() {

	//beego.Router("/camera_view/pic.jpg", &controllers.CameraViewCtr{}, "get:GetJpg")

	ns := beego.NewNamespace("/API",
		beego.NSNamespace("/weather_info"),
			beego.NSInclude(
				&controllers.WeatherInfoController{},
		),
		beego.NSNamespace("/weather_img",
			beego.NSInclude(
				&controllers.WeatherImgController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
