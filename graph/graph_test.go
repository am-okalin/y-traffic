package graph

import (
	"github.com/am-okalin/kit/tableconv"
	"testing"
)

func Test1(t *testing.T) {
	navigations := Navigations()
	table, err := tableconv.Objs2Table(navigations)
	if err != nil {
		t.Error(err)
	}

	err = tableconv.ToCsv(table, NavigationsData)
	if err != nil {
		t.Error(err)
	}
}
