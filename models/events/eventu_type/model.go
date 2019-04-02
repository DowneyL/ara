package eventu_type

import (
	"ara/components/mysql"
	"github.com/astaxie/beego/orm"
)

type Model struct {
	Id                int    `orm:"column(id)"`
	Code              string `orm:"column(code)"`
	IsUse             int    `orm:"column(is_use)"`
	DefaultToTypeId   int    `orm:"column(default_to_type_id)"`
	DefaultToUsername string `orm:"column(default_to_username)"`
	NotifyWayGroupId  int    `orm:"column(notify_way_group_id)"`
	Desc              string `orm:"column(desc)"`
	CreatedAt         string `orm:"column(created_at)"`
	UpdatedAt         string `orm:"column(updated_at)"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Model))
}

func (this *Model) TableName() string {
	return "eventu_type"
}

// 通过code查询事件
func GetTypeByCode(code string) (Model, bool) {
	orm := mysql.GetOrmer()
	eventuType := Model{Code: code}

	err := orm.Read(&eventuType)
	if err != nil {
		return eventuType, false
	}
	return eventuType, true
}

func GetTypeByCode2(code string) ([]Model, error) {
	var types []Model
	qs := mysql.GetQuerySetter(&Model{})
	_, err := qs.Filter("code", code).All(&types, "id", "desc")
	if err != nil {
		return nil, err
	}

	return types, nil
}
