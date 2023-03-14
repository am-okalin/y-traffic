package biz

import (
	"fmt"
	"github.com/am-okalin/kit/tableconv"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	navigations := Navigations(Objs())
	table, err := tableconv.Objs2Table(navigations)
	if err != nil {
		t.Error(err)
	}

	err = tableconv.ToCsv(table, NavigationsData)
	if err != nil {
		t.Error(err)
	}
}

func Test2(t *testing.T) {
	objs := Objs()
	pm := PathMap(objs)
	start := "临江门"
	end := "尖顶坡"
	path := pm[fmt.Sprintf("%s_%s", start, end)]
	t.Log(path)
}

// 全网断面客流量， 指定时间
func Test3(t *testing.T) {
	fname := Csn1
	date := "210816"
	end, _ := time.Parse(time.RFC3339, "2021-08-16T09:00:00Z")
	start, _ := time.Parse(time.RFC3339, "2021-08-16T06:30:00Z")
	stations := []string{"朝天门", "小什字", "较场口", "七星岗", "两路口", "鹅岭", "大坪", "石油路", "歇台子", "重庆工商大学", "四公里",
		"南坪", "工贸", "铜元局", "牛角沱", "华新街", "观音桥", "红旗河沟", "嘉州路", "郑家院子", "唐家院子", "狮子坪", "重庆北站南广场",
		"龙头寺", "童家院子", "刘家坪", "大剧院", "江北城", "五里店", "红土地", "黄泥磅", "花卉园", "大龙山", "冉家坝", "光电园", "仁济",
		"上新街", "海棠溪", "罗家坝", "南湖"}

	trips := TripFilter(Trips(), date, start, end)
	table := CsnTable(trips, stations)
	err := tableconv.ToCsv(table, fname)
	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}
