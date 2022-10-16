package internal

import (
	"testing"
	"y-traffic/internal/biz"
)

func append2Trans(arrs ...[]biz.Trans) []biz.Trans {
	length := 0
	for _, arr := range arrs {
		length += len(arr)
	}
	list := make([]biz.Trans, 0, length)
	for i, _ := range arrs {
		list = append(arrs[i])
	}
	return list
}

func TestAnalysis(t *testing.T) {
	table, err := biz.Unmarshal(biz.TransFile, ',')
	if err != nil {
		t.Error(err)
	}
	list := biz.Table2Trans(table)
	dateM := biz.TransGroup(list, "TransDate")
	weekend := append2Trans(dateM["210818"], dateM["210819"])
	working := append2Trans(dateM["210816"], dateM["210817"], dateM["210820"], dateM["210821"], dateM["210822"])
	interval := biz.NewMinuteInterval(15)
	err = biz.Marshal(biz.Analysis29, interval.Interval2Table(dateM["210816"]))
	if err != nil {
		t.Error(err)
	}
	err = biz.Marshal(biz.Analysis28Weekend, interval.Interval2Table(weekend))
	if err != nil {
		t.Error(err)
	}
	err = biz.Marshal(biz.Analysis28Working, interval.Interval2Table(working))
	if err != nil {
		t.Error(err)
	}
	t.Log("done")
}

func TestTransIntegration(t *testing.T) {
	trips, err := biz.YD2Trip(biz.YDData)
	if err != nil {
		t.Error(err)
	}
	list, err := biz.IC2Trans(biz.ICData)
	if err != nil {
		t.Error(err)
	}
	t.Log("数据加载成功")
	list = biz.FilterByGroup(list)
	t.Log("IC过滤成功")
	list = append(list, biz.Trip2Trans(trips)...)
	t.Log("数据融合成功")
	table := biz.Trans2Table(list)
	err = biz.Marshal(biz.TransFile, table)
	if err != nil {
		t.Error(err)
	}
	t.Log("done")
}
