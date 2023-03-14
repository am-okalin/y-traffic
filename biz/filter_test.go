package biz

import (
	"github.com/am-okalin/kit/tableconv"
	"testing"
)

// TestTransIntegration 获取不同的数据源进行融合
func TestTransIntegration(t *testing.T) {
	//YD数据转为Trip行程
	trips, err := YD2Trip(YDData)
	if err != nil {
		t.Error(err)
	}
	//IC数据转为Trans交通
	list, err := IC2Trans(ICData)
	if err != nil {
		t.Error(err)
	}
	t.Log("数据加载成功")

	//对IC->Trans数据进行数据清洗
	list = FilterByGroup(list)
	t.Log("IC过滤成功")

	//对YD数据进行拆分 -> 清洗 -> 过滤 -> 追加到Trans
	list = append(list, FilterByGroup(Trips2Trans(trips))...)
	t.Log("YD过滤成功, 数据融合成功")

	//将清洗后的数据导入文件
	tab := Trans2Table(list)
	err = tableconv.ToCsv(tab, TransData)
	if err != nil {
		t.Error(err)
	}
	t.Log("done")
}

func TestTripCsv(t *testing.T) {
	// tran list
	trans := Trans()

	// trip list
	trips, err := Trans2Trips(trans)
	if err != nil {
		t.Error(err)
	}

	// set path
	pm := PathMap(Objs())
	trips = SetTripPath(trips, pm)

	//trips csv
	tripT := Trips2Table(trips)
	err = tableconv.ToCsv(tripT, TripsData)
	if err != nil {
		t.Error(err)
	}
	t.Log("done")
}
