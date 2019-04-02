package notify_way_group

import (
	"ara/models/notify/notify_way"
	"github.com/astaxie/beego/orm"
	"strings"
)

type Model struct {
	Id        int
	Code      string
	Desc      string
	CreatedAt string
	UpdatedAt string
}

type GroupInfo struct {
	Model
	ways []notify_way.Model
}

func init() {
	orm.RegisterModel(new(Model))
}

func (this *Model) TableName() string {
	return "notify_way_group"
}

func (this *Model) GetGroupInfosByCode(code string) map[string]interface{} {
	db := orm.NewOrm()
	db.Using("Master")
	wayIdInGroupSql := "select nwr.rid as rid,nwr.gid as gid from notify_way_relate as nwr LEFT JOIN notify_way_group as nwg on nwg.id=nwr.gid where nwg.code= ?"
	var wayRaw []orm.Params
	waysNum, err := db.Raw(wayIdInGroupSql, code).Values(&wayRaw)
	if err != nil {

	}

	// 获取通知方式的id
	var wayId = make([]string, waysNum)
	for idx, wayIdItem := range wayRaw {
		rid := wayIdItem["rid"]
		if rid, ok := rid.(string); ok {
			wayId[idx] = rid
		}
	}

	// 获取通知方式的详情
	var wayList []orm.Params
	wayStr := strings.Join(wayId, ",")
	usingNum, err := db.Raw("select * from notify_way where id in (?) and is_use=1", wayStr).Values(&wayList)
	if err != nil {

	}

	wayInfos := make([]map[string]interface{}, usingNum)
	for idx, wayInfoItem := range wayList {
		wayInfos[idx] = map[string]interface{}{
			"id":       wayInfoItem["id"],
			"code":     wayInfoItem["code"],
			"config":   wayInfoItem["config"],
			"template": wayInfoItem["template"],
			"desc":     wayInfoItem["desc"],
		}
	}

	gid := wayRaw[0]["gid"]
	data := map[string]interface{}{
		"gid":      gid,
		"wayInfos": wayInfos,
	}

	return data

}
