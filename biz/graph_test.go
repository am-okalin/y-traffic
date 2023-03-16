package biz

import (
	"fmt"
	"github.com/am-okalin/kit/tableconv"
	"testing"
)

func Test1(t *testing.T) {
	navigations := Navigations(Stations())
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
	objs := Stations()
	pm := PathMap(objs)
	start := "临江门"
	end := "尖顶坡"
	path := pm[fmt.Sprintf("%s_%s", start, end)]
	t.Log(path)
}
