package sql

import (
	"github.com/am-okalin/kit/tableconv"
	"github.com/am-okalin/y-traffic/filter"
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

	err = db.AutoMigrate(&filter.Trans{})
	if err != nil {
		t.Fatal(err)
	}

	table, err := tableconv.Csv2Table(filter.TransData, ',')
	if err != nil {
		t.Error(err)
	}
	trans := filter.Table2Trans(table)

	t.Log("开始导入数据")

	result := db.CreateInBatches(trans, 1000)

	if result.Error != nil {
		t.Log(result.Error)
	}
	t.Log("done")
}
