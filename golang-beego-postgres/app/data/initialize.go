package data

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Author))
	orm.RegisterModel(new(Book))
}
