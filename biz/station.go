package biz

import (
	"github.com/am-okalin/kit/tableconv"
	"strconv"
)

type Obj struct {
	Line string
	Id   string
	Name string
	Vi   int
}

// Objs 返回站台列表
func Objs() []Obj {
	table, err := tableconv.Csv2Table(Stations, ',')
	if err != nil {
		panic(err)
	}

	stations := make([]Obj, len(table)-1)
	for i := 1; i < len(table); i++ {
		vi, err := strconv.Atoi(table[i][3])
		if err != nil {
			panic(err)
		}

		stations[i-1] = Obj{
			Line: table[i][0],
			Id:   table[i][1],
			Name: table[i][2],
			Vi:   vi,
		}
	}

	return stations
}

// IdM 初始化StationM
func IdM() map[string]Obj {
	m := make(map[string]Obj)
	stations := Objs()
	for i, station := range stations {
		m[station.Id] = stations[i]
	}
	return m
}
