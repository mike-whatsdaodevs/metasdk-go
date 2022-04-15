package main

import (
	_ "metasdk/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

