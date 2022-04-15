package Models

import (
    "github.com/beego/beego/v2/client/orm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)

type NFT struct {
    Id          int    `orm:"column(id)" json:"id"`
    Address     string `orm:"column(address);description(用户地址)" json:"address"`
    TokenId     int `orm:"column(token_id);description(数量)" json:"token_id"`
    Status      uint8  `orm:"column(status);description(状态)" json:"status"`
    CreatedAt   string `orm:"description(创建时间);column(created_at)" json:"created_at"`
}


func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    // set default database
    orm.RegisterDataBase("default", "mysql", "root:@(127.0.0.1)/metabtc?charset=utf8")
    orm.Debug = true
    // // // register model
    // orm.RegisterModel(new(Stu))
    orm.RegisterModel(new(NFT))

    orm.RunSyncdb("default", false, true)


    // // // create table
    // orm.RunSyncdb("mermaidnft", false, true)
}