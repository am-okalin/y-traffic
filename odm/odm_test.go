package odm

import (
	"github.com/am-okalin/y-traffic/station"
	"testing"
	"time"
)

// TestOdmAll 输出od表
func TestOdmAll(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T09:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T06:30:00Z")

	m := make(map[string]bool)
	for _, obj := range station.IdM() {
		m[obj.Name] = true
	}
	var allStations = make([]string, 0)
	for name, _ := range m {
		allStations = append(allStations, name)
	}

	err := Odm(start, end, allStations, OdmAll)
	if err != nil {
		t.Error(err)
	}
}

// TestOdm1 输出od表
func TestOdm1(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T09:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T06:30:00Z")
	stations1 := []string{"朝天门", "小什字", "较场口", "七星岗", "两路口", "鹅岭", "大坪", "石油路", "歇台子", "重庆工商大学", "四公里",
		"南坪", "工贸", "铜元局", "牛角沱", "华新街", "观音桥", "红旗河沟", "嘉州路", "郑家院子", "唐家院子", "狮子坪", "重庆北站南广场",
		"龙头寺", "童家院子", "刘家坪", "大剧院", "江北城", "五里店", "红土地", "黄泥磅", "花卉园", "大龙山", "冉家坝", "光电园", "仁济",
		"上新街", "海棠溪", "罗家坝", "南湖"}

	err := Odm(start, end, stations1, Odm1)
	if err != nil {
		t.Error(err)
	}
}

// TestOdm2 输出od表
func TestOdm2(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T07:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T06:30:00Z")
	stations1 := []string{"朝天门", "小什字", "较场口", "七星岗", "两路口", "鹅岭", "大坪", "石油路", "歇台子", "重庆工商大学", "四公里",
		"南坪", "工贸", "铜元局", "牛角沱", "华新街", "观音桥", "红旗河沟", "嘉州路", "郑家院子", "唐家院子", "狮子坪", "重庆北站南广场",
		"龙头寺", "童家院子", "刘家坪", "大剧院", "江北城", "五里店", "红土地", "黄泥磅", "花卉园", "大龙山", "冉家坝", "光电园", "仁济",
		"上新街", "海棠溪", "罗家坝", "南湖"}

	err := Odm(start, end, stations1, Odm2)
	if err != nil {
		t.Error(err)
	}
}

// TestOdm3 输出od表
func TestOdm3(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T12:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T06:30:00Z")
	stations1 := []string{"朝天门", "小什字", "较场口", "七星岗", "两路口", "鹅岭", "大坪", "石油路", "歇台子", "重庆工商大学", "四公里",
		"南坪", "工贸", "铜元局", "牛角沱", "华新街", "观音桥", "红旗河沟", "嘉州路", "郑家院子", "唐家院子", "狮子坪", "重庆北站南广场",
		"龙头寺", "童家院子", "刘家坪", "大剧院", "江北城", "五里店", "红土地", "黄泥磅", "花卉园", "大龙山", "冉家坝", "光电园", "仁济",
		"上新街", "海棠溪", "罗家坝", "南湖"}

	err := Odm(start, end, stations1, Odm3)
	if err != nil {
		t.Error(err)
	}
}
