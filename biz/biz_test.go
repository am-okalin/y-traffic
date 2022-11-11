package biz

import (
	"testing"
	"y-traffic/tableconv"
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
	t.Log("数据融合成功")
	tab := Trans2Table(list)
	err = tableconv.ToCsv(tab, TransData)
	if err != nil {
		t.Error(err)
	}
	t.Log("done")
}

// TestAnalysis 生成日期分析数据
func TestAnalysis(t *testing.T) {
	table, err := tableconv.Csv2Table(TransData, ',')
	if err != nil {
		t.Error(err)
	}
	list := Table2Trans(table)
	dateM := TransGroup(list, "TransDate")
	weekend := Append2Trans(dateM["210818"], dateM["210819"])
	working := Append2Trans(dateM["210816"], dateM["210817"], dateM["210820"], dateM["210821"], dateM["210822"])
	interval := NewMinuteInterval(15)
	err = tableconv.ToCsv(interval.Interval2Table(dateM["210816"]), Analysis29)
	if err != nil {
		t.Error(err)
	}
	err = tableconv.ToCsv(interval.Interval2Table(weekend), Analysis28Weekend)
	if err != nil {
		t.Error(err)
	}
	err = tableconv.ToCsv(interval.Interval2Table(working), Analysis28Working)
	if err != nil {
		t.Error(err)
	}
	t.Log("done")
}

// TestLineData 生成各个线路16号的进出站数据
func TestLine(t *testing.T) {
	table, err := tableconv.Csv2Table(TransData, ',')
	if err != nil {
		t.Error(err)
	}
	list := Table2Trans(table)
	TransDateM := TransGroup(list, "TransDate")
	LineM := TransGroup(TransDateM["210816"], "Line")

	for line, trans := range LineM {
		err = tableconv.ToCsv(Trans2Table(trans), PrefixLine+line+".csv")
		if err != nil {
			t.Error(err)
		}
	}
	t.Log("done")
}

// TestLineDate 生成各个线路16号的进出站数据
func TestLineDate(t *testing.T) {
	table, err := tableconv.Csv2Table(TransData, ',')
	if err != nil {
		t.Error(err)
	}
	list := Table2Trans(table)
	TransDateM := TransGroup(list, "TransDate")
	LineM := TransGroup(TransDateM["210816"], "Line")

	interval := NewMinuteInterval(15)
	for line, trans := range LineM {
		err = tableconv.ToCsv(interval.Interval2Table(trans), PrefixLineDate+line+".csv")
		if err != nil {
			t.Error(err)
		}
	}
	t.Log("done")
}
