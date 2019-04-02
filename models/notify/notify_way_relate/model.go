package notify_way_relate

import "github.com/astaxie/beego/orm"

type Model struct {
	Id        int
	Gid       int
	Rid       int
	CreatedAt string
	UpdatedAt string
}

func init() {
	orm.RegisterModel(new(Model))
}

func (this *Model) TableName() string {
	return "notify_way_relate"
}
