package search

import (
	"github.com/am-okalin/kit/tableconv"
	"log"
)

// StationM stationId=>stationName
var StationM = InitStationM()

// InitStationM 初始化StationM
func InitStationM() map[string]string {
	table, err := tableconv.Csv2Table(Stations, ',')
	if err != nil {
		log.Fatal(err)
	}
	m, rowLen := tableconv.ToM(table)
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
