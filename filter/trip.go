package filter

import (
	"fmt"
	"github.com/am-okalin/kit/tableconv"
	"github.com/am-okalin/y-traffic/graph"
	"strings"
	"time"
)

// Trip 行程
type Trip struct {
	TripId     string //行程ID
	Date       string
	Path       []string
	InTransId  string //进站ID
	OutTransId string //出站ID
	In         Trans  //进站的 票ID 站台ID 站台名称 时间
	Out        Trans  //出站的 票ID 站台ID 站台名称 时间
}

func Trips2Trans(trips []Trip) []Trans {
	list := make([]Trans, 0, len(trips)*2)
	for _, trip := range trips {
		list = append(list, trip.In)
		list = append(list, trip.Out)
	}
	return list
}

func SetTripPath(trips []Trip, pm map[string][]graph.Vertex) []Trip {
	for i, trip := range trips {
		key := fmt.Sprintf("%s_%s", trip.In.StationName, trip.Out.StationName)
		trips[i].Path = make([]string, len(pm[key]))
		for j, vertex := range pm[key] {
			trips[i].Path[j] = vertex.Name
		}
	}
	return trips
}

func Trips2Table(trips []Trip) [][]string {
	length := len(trips)
	table := make([][]string, 0, length+1)
	table = append(table, []string{
		"TripId",
		"Date",
		"InTransId",
		"OutTransId",
		"InTransTime",
		"OutTransTime",
		"InLine",
		"OutLine",
		"InStationId",
		"OutStationId",
		"InStationName",
		"OutStationName",
		"Path",
	})
	for i := 0; i < length; i++ {
		table = append(table, []string{
			trips[i].TripId,
			trips[i].In.Date,
			trips[i].In.TransId,
			trips[i].Out.TransId,
			trips[i].In.TransTime.Format(time.RFC3339),
			trips[i].Out.TransTime.Format(time.RFC3339),
			trips[i].In.Line,
			trips[i].Out.Line,
			trips[i].In.StationId,
			trips[i].Out.StationId,
			trips[i].In.StationName,
			trips[i].Out.StationName,
			strings.Join(trips[i].Path, "-"),
		})
	}
	return table
}

// Trips 获取trips
func Trips() []Trip {
	table, err := tableconv.Csv2Table(TripsData, ',')
	if err != nil {
		panic(err)
	}
	m, rowLen := tableconv.ToM(table)
	list := make([]Trip, rowLen)
	for i := 0; i < rowLen; i++ {
		list[i].TripId = m["TripId"][i]
		list[i].Date = m["Data"][i]
		list[i].In.Date = m["Data"][i]
		list[i].Out.Date = m["Data"][i]
		list[i].In.TransId = m["InTransId"][i]
		list[i].Out.TransId = m["OutTransId"][i]
		list[i].In.TransTime, _ = time.Parse(time.RFC3339, m["InTransTime"][i])
		list[i].Out.TransTime, _ = time.Parse(time.RFC3339, m["OutTransTime"][i])
		list[i].In.Line = m["InLine"][i]
		list[i].Out.Line = m["OutLine"][i]
		list[i].In.StationId = m["InStationId"][i]
		list[i].Out.StationId = m["OutStationId"][i]
		list[i].In.StationName = m["InStationName"][i]
		list[i].Out.StationName = m["OutStationName"][i]
		list[i].Path = strings.Split(m["OutPath"][i], "-")
	}
	return list
}
