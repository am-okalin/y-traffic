package search

import (
	"github.com/am-okalin/kit/tableconv"
	"github.com/am-okalin/y-traffic/filter"
	"testing"
)

// TestAnalysis 生成日期分析数据
func TestAnalysis(t *testing.T) {
	table, err := tableconv.Csv2Table(TransData, ',')
	if err != nil {
		t.Error(err)
	}
	list := filter.Table2Trans(table)
	dateM := filter.TransGroup(list, "Date")
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
	list := filter.Table2Trans(table)
	TransDateM := filter.TransGroup(list, "Date")
	LineM := filter.TransGroup(TransDateM["210816"], "Line")

	for line, trans := range LineM {
		err = tableconv.ToCsv(filter.Trans2Table(trans), PrefixLine+line+".csv")
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
	list := filter.Table2Trans(table)
	TransDateM := filter.TransGroup(list, "Date")
	LineM := filter.TransGroup(TransDateM["210816"], "Line")

	interval := NewMinuteInterval(15)
	for line, trans := range LineM {
		err = tableconv.ToCsv(interval.Interval2Table(trans), PrefixLineDate+line+".csv")
		if err != nil {
			t.Error(err)
		}
	}
	t.Log("done")
}

func Test1226(t *testing.T) {
	//1.周五_0820_观音桥_进站
	//2.周五_0820_观音桥_出站
	//3.周一_0816_光电园_进站
	//4.周一_0816_光电园_出站
	//5.周六_0821_重庆北站南广场_进站
	//6.周六_0821_重庆北站南广场_出站
	//7.周六_0821_04_进站
	//8.周六_0821_04_出站
}
