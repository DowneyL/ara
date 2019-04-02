package eventu_ext

import "github.com/astaxie/beego/orm"

type Model struct {
	Id        int64
	EventuId  int
	State     int
	CreatedAt string
	UpdatedAt string
}

const (
	STATE_BEGIN = iota
	STATE_INIT
	STATE_BEFORE_CREATE
	STATE_AFTER_CREATE
)

func init() {
	orm.RegisterModel(new(Model))
}

func (this *Model) TableName() string {
	return "eventu_ext"
}
