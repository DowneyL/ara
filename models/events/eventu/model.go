package eventu

import "github.com/astaxie/beego/orm"

type Model struct {
	Id               int
	TypeId           int
	Content          string
	FromTypeId       int
	FromUsername     string
	ToTypeId         int
	ToUsername       string
	NotifyWayGroupId int
	CreatedAt        string
	UpdatedAt        string
}

func init() {
	orm.RegisterModel(new(Model))
}

func (this *Model) TableName() string {
	return "eventu"
}
