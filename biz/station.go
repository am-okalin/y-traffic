package biz

import (
	"log"
	"y-traffic/common"
	"y-traffic/table"
)

// StationM stationId=>stationName
var StationM = InitStationM()

// InitStationM 初始化StationM
func InitStationM() map[string]string {
	table, err := table.Csv2Table(common.Stations, ',')
	if err != nil {
		log.Fatal(err)
	}
	m, rowLen := Table2Map(table)
	stationM := make(map[string]string, rowLen)

	for i := 0; i < rowLen; i++ {
		stationM[m["station_id"][i]] = m["station_name"][i]
	}
	return stationM
}

// StationNameById 获取车站名称
func StationNameById(id string) string {
	return StationM[id]
}

type Station struct {
	Id   string
	Name string
}

func StationsByLine() map[string][]Station {
	m := make(map[string][]Station)
	for id, name := range StationM {
		stations, ok := m[id[0:2]]
		if !ok {
			m[id[0:2]] = make([]Station, 0)
		}

		stations = append(stations, Station{id, name})
	}
	return m
}
