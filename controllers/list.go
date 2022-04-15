package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	Model "metasdk/models"
)

type ListController struct {
	beego.Controller
}

type Res struct{
	List []Model.NFT `json:"list"`
	Total int `json:"total"`
	CurrentPage int `json:"current_page"`
}

const PageNum = 50

func (c *ListController) List() {
	page, _ := c.GetInt("page")

	skip := 0;

	if page <= 0 {
		page = 1
	}

	skip = (page - 1) * PageNum;

	total := int(Model.GetListNum())

	list := Model.GetList(PageNum, skip)
	
	res := Res{
		CurrentPage : page,
		List: list,
		Total: total,
	}

	c.Data["json"] = &res
    c.ServeJSON()
}



// public(script) fun create_pair<TokenTypeX: store, TokenTypeY: store>(account: signer) {
//       Factory::create_pair<TokenTypeX, TokenTypeY>(&account);
//   }