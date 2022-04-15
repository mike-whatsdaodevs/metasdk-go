package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	Model "metasdk/models"
)

type AddController struct {
	beego.Controller
}

func (c *AddController) AddWl() {
	tokenId,_ := c.GetInt("tokenId")
	address := c.GetString("address")

	Model.ReadAndCreateOrUpdate(address, tokenId)
	c.Data["json"] = true
    c.ServeJSON()
}


// public(script) fun create_pair<TokenTypeX: store, TokenTypeY: store>(account: signer) {
//       Factory::create_pair<TokenTypeX, TokenTypeY>(&account);
//   }