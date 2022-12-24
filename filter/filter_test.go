package filter

import (
	"github.com/am-okalin/kit/tableconv"
	"github.com/am-okalin/y-traffic/pkg"
	"testing"
)

// TestTransIntegration 获取不同的数据源进行融合
func TestTransIntegration(t *testing.T) {
	trips, err := YD2Trip(YDData)
	if err != nil {
		t.Error(err)
	}
	list, err := IC2Trans(ICData)
	if err != nil {
		t.Error(err)
	}
	t.Log("数据加载成功")
	list = FilterByGroup(list)
	t.Log("IC过滤成功")
	list = append(list, FilterByGroup(Trip2Trans(trips))...)
	t.Log("YD过滤成功, 数据融合成功")
	tab := pkg.Trans2Table(list)
	err = tableconv.ToCsv(tab, TransData)
	if err != nil {
		t.Error(err)
	}
	t.Log("done")
}
