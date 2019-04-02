package notify_way

import (
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id        int
	Code      string
	IsUse     int
	Config    string
	Template  string
	Desc      string
	CreatedAt string
	UpdatedAr string
}

func init() {
	orm.RegisterModel(new(Model))
}

func (this *Model) TableName() string {
	return "notify_way"
}

func (this Model) GetWayByCode(code string) (Model, bool) {
	db := orm.NewOrm()
	db.Using("Master")
	notifyWay := Model{Code: code}
	if err := db.Read(&notifyWay); err != nil {
		return notifyWay, false
	}
	return notifyWay, true
}
