package biz

import (
	"log"
	"time"
)

//Trip 行程
type Trip struct {
	TripId      string //行程ID
	MatchMarker string //匹配标记[进站抵消 单边进 单边出 补登进 补登出 出站票]
	InTransId   string //进站ID
	OutTransId  string //出站ID
	In          Trans  //进站的 票ID 站台ID 站台名称 时间
	Out         Trans  //出站的 票ID 站台ID 站台名称 时间
}

func YD2Trip(fname string) ([]Trip, error) {
	table, err := Unmarshal(fname, '\t')
	if err != nil {
		return nil, err
	}
	m, rowLen := Table2Map(table)
	list := make([]Trip, 0, rowLen)
	for i := 0; i < rowLen; i++ {
		if m["匹配标记"][i] == "进站抵消" || m["匹配标记"][i] == "单边进" || m["匹配标记"][i] == "单边出" {
			log.Printf("error row[%d]: %s", i, m["匹配标记"][i])
			continue
		}
		tmp := Trip{
			TripId:      m["行程编号"][i],
			MatchMarker: m["匹配标记"][i],
			InTransId:   m["进站行程编号"][i],
			OutTransId:  m["出站行程编号"][i],
			In: Trans{
				TransId:     m["进站行程编号"][i],
				TicketId:    m["虚拟卡号"][i],
				Line:        m["IN_STATION_ID"][i][0:2],
				StationId:   m["IN_STATION_ID"][i],
				StationName: m["进站车站"][i],
				TransCode:   "21",
				TransTime:   time.Time{},
			},
			Out: Trans{
				TransId:     m["出站行程编号"][i],
				TicketId:    m["虚拟卡号"][i],
				Line:        m["OUT_STATION_ID"][i][0:2],
				StationId:   m["OUT_STATION_ID"][i],
				StationName: m["出站车站"][i],
				TransCode:   "22",
				TransTime:   time.Time{},
			},
		}
		tmp.In.TransTime, err = time.Parse("2006/1/2 5:04:05", m["进站时间"][i])
		if err != nil {
			return nil, err
		}
		tmp.Out.TransTime, err = time.Parse("2006/1/2 5:04:05", m["出站时间"][i])
		if err != nil {
			return nil, err
		}
		list = append(list, tmp)
	}
	return list, nil
}
