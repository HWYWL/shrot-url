package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

//新浪接口地址和source(需要开发者账号申请)
const api = "https://api.weibo.com/2/short_url/shorten.json?"
const source = "1950792609"

type UrlController struct {
	beego.Controller
}

// 获取前端数据传回来的长连接
func (this *UrlController) Get() {
	url := this.GetString("url")

	//请求新浪接口
	req := httplib.Post(api)
	req.Param("source",source)
	req.Param("url_long",url)

	s, err := req.String()

	this.Data["json"] = map[string]interface{}{"code": 0,"data": s, "message": err}
	this.ServeJSON()
	return
}
