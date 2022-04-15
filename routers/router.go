package routers

import (
	"metasdk/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/miner/nftlist", &controllers.ListController{},"get:List")
    beego.Router("/miner/nftadd", &controllers.AddController{},"get:AddWl")
}
