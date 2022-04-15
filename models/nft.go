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
func GetList(take, skip int) (list []NFT) {
    list = []NFT{}

	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("id", "address", "token_id","status", "created_at").
		From("NFT").
		Where("status > 0")

	qb.OrderBy("token_id ASC, id ASC").
        Limit(take).Offset(skip)

    sql := qb.String()


	o.Raw(sql).QueryRows(&list)

	return list
}

func GetListNum() (num int64) {

    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("count(*) as num").
        From("NFT").
        Where("status > 0")

    sql := qb.String()


    err := o.Raw(sql).QueryRow(&num)
    if err == nil {
        fmt.Println("user nums: ", num)
    }


    return num
}


func ReadAndCreateOrUpdate(address string, tokenId uint64) (bool){
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


func InsertNFTr(address string, amount uint64) {
    o := orm.NewOrm()

    st := new(NFT)
    st.Address = address
    st.Amount = amount
    st.Status = 1
    timestr := time.Now().Format("2006-01-0215:04:05")
    st.CreatedAt = timestr

    o.Insert(st)

}

func GetAddressRand(address string, amount uint64) (num int) {
    o := orm.NewOrm()

    id := GetAddressId(address);
    st := NFT{Id: id, Address: address}

    if id == 0 || o.Read(&st) != nil {
        return 0;
    }


    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("count(*) as num").
        From("NFT").
        Where("status > 0").
        And("(amount > ? OR (amount = ? AND id < ?))")

    sql := qb.String()

    err := o.Raw(sql, amount, amount, id).QueryRow(&num)
    if err == nil {
        fmt.Println("user nums: ", num)
    }

    return num + 1
}



func GetTokenId(tokenId uint) (num int) {
    o := orm.NewOrm()

    qb, _ := orm.NewQueryBuilder("mysql")

    qb.Select("id").
        From("NFT").
        Where("status > 0").
        And("token_id = ?")

    sql := qb.String()


    err := o.Raw(sql, address).QueryRow(&num)
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










