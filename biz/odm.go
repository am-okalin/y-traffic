package biz

import (
	"github.com/am-okalin/kit/tableconv"
	"sort"
	"strconv"
	"time"
)

func Odm(start, end time.Time, stations []string, fname string) error {
	//过滤站点时间
	origin := Screening(Trans(), stations, start, end)

	//[]Tran -> []Trip
	trips, err := GetTrips(origin)
	if err != nil {
		return err
	}

	//初始化odm
	length := len(stations)
	odm := make(map[string]map[string]int, length)
	for _, si := range stations {
		odm[si] = make(map[string]int, length)
		for _, sj := range stations {
			odm[si][sj] = 0
		}
	}

	//遍历trips得出odm
	for _, trip := range trips {
		odm[trip.In.StationName][trip.Out.StationName]++
	}

	//输出odm至table
	table := make([][]string, length+1)
	table[0] = append([]string{"O\\D"}, stations...)
	for row := 1; row <= length; row++ {
		table[row] = make([]string, length+1)
		table[row][0] = stations[row-1]
		for col := 1; col <= length; col++ {
			table[row][col] = strconv.Itoa(odm[stations[row-1]][stations[col-1]])
		}
	}

	//导出table至csv
	return tableconv.ToCsv(table, fname)
}

func GetTrips(origin []Tran) ([]Trip, error) {
	// 按ticker_id分组匹配
	trips := make([]Trip, 0)
	m := TransGroup(origin, "TicketId")
	for ticketId, list := range m {
		if len(list) <= 1 {
			m[ticketId] = nil
			continue
		}
		//进出站匹配
		sort.Slice(list, func(i, j int) bool { return list[i].TransTime.Before(list[j].TransTime) })
		m[ticketId] = InOutMatch(list)
		trips = append(trips, TripsByTicket(ticketId, m[ticketId])...)
	}

	return trips, nil
}

// Screening 按时间过滤
func Screening(old []Tran, stations []string, start, end time.Time) []Tran {
	list := make([]Tran, 0)

	m := make(map[string]bool, len(stations))
	for _, station := range stations {
		m[station] = true
	}

	for i, trans := range old {
		if !m[trans.StationName] {
			continue
		}
		if trans.TransTime.Before(start) {
			continue
		}
		if trans.TransTime.After(end) {
			continue
		}
		list = append(list, old[i])
	}
	return list
}