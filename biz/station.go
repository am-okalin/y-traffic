package biz

import (
	"github.com/am-okalin/kit/tableconv"
	"strconv"
)

type Station struct {
	Line string
	Id   string
	Name string
	Vi   int
}

// Stations 返回站台列表
func Stations() []Station {
	table, err := tableconv.Csv2Table(StationsData, ',')
	if err != nil {
		panic(err)
	}

	stations := make([]Station, len(table)-1)
	for i := 1; i < len(table); i++ {
		vi, err := strconv.Atoi(table[i][3])
		if err != nil {
			panic(err)
		}

		stations[i-1] = Station{
			Line: table[i][0],
			Id:   table[i][1],
			Name: table[i][2],
			Vi:   vi,
		}
	}

	return stations
}

// IdStationM 初始化StationM
func IdStationM() map[string]Station {
	m := make(map[string]Station)
	stations := Stations()
	for i, station := range stations {
		m[station.Id] = stations[i]
	}
	return m
}

func StationNames() []string {
	stations := Stations()
	m := make(map[string]bool)

	for _, station := range stations {
		m[station.Name] = true
	}

	names := make([]string, len(m))

	for _, station := range stations {
		if m[station.Name] {
			names = append(names, station.Name)
		}
	}
	return names
}
