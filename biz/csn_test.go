package biz

import (
	"github.com/am-okalin/kit/tableconv"
	"testing"
	"time"
)

//1、0816日期的6：30-7:00的局部网络断面客流量；
//2、0816日期的7：00-9:00的局部网络断面客流量；
//4、0816日期的7：00-9:00的全网断面客流量；
//end, _ := time.Parse(time.RFC3339, "2021-08-16T24:00:00Z")
//start, _ := time.Parse(time.RFC3339, "2021-08-16T00:00:00Z")

var stations = []string{"朝天门", "小什字", "较场口", "七星岗", "两路口", "鹅岭", "大坪", "石油路", "歇台子", "重庆工商大学", "四公里",
	"南坪", "工贸", "铜元局", "牛角沱", "华新街", "观音桥", "红旗河沟", "嘉州路", "郑家院子", "唐家院子", "狮子坪", "重庆北站南广场",
	"龙头寺", "童家院子", "刘家坪", "大剧院", "江北城", "五里店", "红土地", "黄泥磅", "花卉园", "大龙山", "冉家坝", "光电园", "仁济",
	"上新街", "海棠溪", "罗家坝", "南湖"}

func TestCsn08167090past(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T09:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T07:00:00Z")
	trips := TripFilter(Trips(), "210816", start, end)
	table := CsnTable(trips, stations)
	err := tableconv.ToCsv(table, Csn08167090past)

	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn08167090(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T09:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T07:00:00Z")
	trips := TripFilter(Trips(), "210816", start, end)
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn08167090past)

	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn08166370(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T07:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T06:30:00Z")
	trips := TripFilter(Trips(), "210816", start, end)
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn08167090past)

	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn0816(t *testing.T) {
	trips := TripFilterDate(Trips(), "210816")
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn0816)
	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn0817(t *testing.T) {
	trips := TripFilterDate(Trips(), "210817")
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn0817)
	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn0818(t *testing.T) {
	trips := TripFilterDate(Trips(), "210818")
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn0818)
	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn0819(t *testing.T) {
	trips := TripFilterDate(Trips(), "210819")
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn0819)
	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn0820(t *testing.T) {
	trips := TripFilterDate(Trips(), "210820")
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn0820)
	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}
