package graph

import (
	"fmt"
	"github.com/am-okalin/kit/tableconv"
	"github.com/am-okalin/y-traffic/filter"
	"github.com/am-okalin/y-traffic/station"
	"testing"
)

func Test1(t *testing.T) {
	navigations := Navigations(station.Objs())
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
	objs := station.Objs()
	pm := PathMap(objs)
	start := "临江门"
	end := "尖顶坡"
	path := pm[fmt.Sprintf("%s_%s", start, end)]
	t.Log(path)
}

// 全网断面客流量， 指定时间
func Test3(t *testing.T) {
	trips := filter.Trips()
	objs := station.Objs()
	pm := PathMap(objs)
	m := InitCsnm(pm)
	for _, trip := range trips {
		for i := 1; i < len(trip.Path); i++ {
			key := fmt.Sprintf("%s_%s", trip.Path[i-1], trip.Path[i])
			m[key]++
		}
	}
	t.Log("done")
}
