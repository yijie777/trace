package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"trace/controllers"
)

func init() {

	beego.Router("/produce", &controllers.Trace{}, "POST:NewFood")
	beego.Router("/userinfo", &controllers.User{}, "*:Userinof")
	beego.Router("/adddistribution", &controllers.Trace{}, "POST:Adddistribution")
	beego.Router("/addretail", &controllers.Trace{}, "POST:Addretail")
	beego.Router("/trace", &controllers.GetFoodInfo{}, "*:Trace")
	beego.Router("/food", &controllers.GetFoodInfo{}, "*:GetFood")
	beego.Router("/producing", &controllers.GetFoodInfo{}, "GET:Producing")
	beego.Router("/distributing", &controllers.GetFoodInfo{}, "GET:Distributing")
	beego.Router("/retailing", &controllers.GetFoodInfo{}, "GET:Retailing")
}
