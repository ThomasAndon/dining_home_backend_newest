package logic

import (
	models "dining_home_backend_newest/service/main/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"testing"
	"xorm.io/core"
)

func TestGetXorm(t *testing.T) {

	engine, err := xorm.NewEngine("mysql", "root:1234@(127.0.0.1:3306)/eat?charset=utf8mb4")
	engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, "d_"))
	if err != nil {
		panic(err)
	}
	msgContent := MsgContent{
		ToUsername:   "wwfc9c5ccc385b29d0",
		FromUsername: "TangCaitong",
		CreateTime:   111111,
		MsgType:      "text",
		Content:      "1",
		Msgid:        "12345678",
		Agentid:      1000002,
	}

	id := &models.Identity{CpwxId: msgContent.FromUsername}

	engine.Get(id)
	t.Log(id)

}
