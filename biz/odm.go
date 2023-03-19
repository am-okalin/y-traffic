package biz

import (
	"github.com/am-okalin/kit/tableconv"
	"strconv"
)

func Odm(StationNames []string, trips []Trip) map[string]map[string]int {
	//初始化odm
	length := len(StationNames)
	odm := make(map[string]map[string]int, length)
	for _, si := range StationNames {
		odm[si] = make(map[string]int, length)
		for _, sj := range StationNames {
			odm[si][sj] = 0
		}
	}
	for _, trip := range trips {
		_, ok := odm[trip.In.StationName][trip.Out.StationName]
		if ok {
			odm[trip.In.StationName][trip.Out.StationName]++
		}
	}
	return odm
}

func OdmCsv(StationNames []string, odm map[string]map[string]int, fname string) error {
	//输出odm至table
	length := len(odm)
	table := make([][]string, length+1)
	table[0] = append([]string{"O\\D"}, StationNames...)
	for row := 1; row <= length; row++ {
		table[row] = make([]string, length+1)
		table[row][0] = StationNames[row-1]
		for col := 1; col <= length; col++ {
			table[row][col] = strconv.Itoa(odm[StationNames[row-1]][StationNames[col-1]])
		}
	}

	//导出table至csv
	return tableconv.ToCsv(table, fname)
}
