package Models

import (
	"fmt"
    "time"
    "github.com/beego/beego/v2/client/orm"
     _ "github.com/jinzhu/gorm/dialects/mysql"
)


func (n *NFT) TableName() string {
	return "NFT"
}

// get NFT list
func GetList(take, skip int, address string) (list []NFT) {
    list = []NFT{}

	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("id", "address", "token_id","status", "created_at").
		From("NFT").
		Where("status > 0")


    if address != "" {
        qb.And("address=?")
    }

	qb.OrderBy("token_id ASC, id ASC").
        Limit(take).Offset(skip)

    sql := qb.String()


	o.Raw(sql, address).QueryRows(&list)

	return list
}

func GetListNum(address string) (num int64) {

    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("count(*) as num").
        From("NFT").
        Where("status > 0")

    if address != "" {
        qb.And("address=?")
    }

    sql := qb.String()


    err := o.Raw(sql, address).QueryRow(&num)
    if err == nil {
        fmt.Println("user nums: ", num)
    }


    return num
}


func ReadAndCreateOrUpdate(address string, tokenId int) (bool){
    o := orm.NewOrm()
    
    id := GetTokenId(tokenId);
    st := NFT{Id: id, TokenId: tokenId}

    if id > 0 && o.Read(&st) == nil {
        if address == st.Address {
            return false;
        }
        timestr := time.Now().Format("2006-01-02 15:04:05")
        st.CreatedAt = timestr
        st.Address = address
        if _, err := o.Update(&st); err == nil {
            return true
        }
    } else {
        st.Address = address
        st.Status = 1
        timestr := time.Now().Format("2006-01-02 15:04:05")
        st.CreatedAt = timestr
        o.Insert(&st)
    }
    return true
}


func InsertNFT(address string, tokenId int) {
    o := orm.NewOrm()

    st := new(NFT)
    st.Address = address
    st.TokenId = tokenId
    st.Status = 1
    timestr := time.Now().Format("2006-01-0215:04:05")
    st.CreatedAt = timestr

    o.Insert(st)

}


func GetTokenId(tokenId int) (num int) {
    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("id").
        From("NFT").
        Where("status > 0").
        And("token_id = ?")

    sql := qb.String()


    err := o.Raw(sql, tokenId).QueryRow(&num)
    if err == nil {
        fmt.Println("user nums: ", num)
    }

    return num 
}

func ChangeNFTAmount(address, amount string) {
    o := orm.NewOrm()

    res, err := o.Raw("UPDATE NFT SET amount = ? WHERE address = ?", amount, address).Exec()
    if err == nil {
        num, _ := res.RowsAffected()
        fmt.Println("mysql row affected nums: ", num)
    }

    return

}










