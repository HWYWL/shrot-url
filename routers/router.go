package routers

import (
	"shrot-url/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/getSinaShortUrl", &controllers.UrlController{})
}
