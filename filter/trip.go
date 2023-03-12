package filter

import "time"

// Trip 行程
type Trip struct {
	TripId      string //行程ID
	MatchMarker string //匹配标记[进站抵消 单边进 单边出 补登进 补登出 出站票]
	InTransId   string //进站ID
	OutTransId  string //出站ID
	In          Trans  //进站的 票ID 站台ID 站台名称 时间
	Out         Trans  //出站的 票ID 站台ID 站台名称 时间
}

func Trips2Trans(trips []Trip) []Trans {
	list := make([]Trans, 0, len(trips)*2)
	for _, trip := range trips {
		list = append(list, trip.In)
		list = append(list, trip.Out)
	}
	return list
}

func Trips2Table(trips []Trip) [][]string {
	length := len(trips)
	table := make([][]string, 0, length+1)
	table = append(table, []string{
		"TripId",
		"TripDate",
		"TripPath",
		"InLine",
		"InStationId",
		"InStationName",
		"InTransId",
		"InTransTime",
		"OutLine",
		"OutStationId",
		"OutStationName",
		"OutTransId",
		"OutTransTime",
	})
	for i := 0; i < length; i++ {
		table = append(table, []string{
			trips[i].TripId,
			trips[i].In.TransDate,
			"",
			trips[i].In.Line,
			trips[i].In.StationId,
			trips[i].In.StationName,
			trips[i].In.TransId,
			trips[i].In.TransTime.Format(time.RFC3339),
			trips[i].Out.Line,
			trips[i].Out.StationId,
			trips[i].Out.StationName,
			trips[i].Out.TransId,
			trips[i].Out.TransTime.Format(time.RFC3339),
		})
	}
	return table
}
