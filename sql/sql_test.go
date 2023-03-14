package sql

import (
	"github.com/am-okalin/y-traffic/biz"
	"testing"
)

const (
	source = "root:root0987@tcp(127.0.0.1:3306)/test-db?parseTime=True"
)

func Test1(t *testing.T) {
	db, err := NewDb(NewDail(source))
	if err != nil {
		t.Fatal(err)
	}

	err = db.AutoMigrate(&biz.Tran{})
	if err != nil {
		t.Fatal(err)
	}

	trans := biz.Trans()
	t.Log("开始导入数据")

	result := db.CreateInBatches(trans, 1000)

	if result.Error != nil {
		t.Log(result.Error)
	}
	t.Log("done")
}
