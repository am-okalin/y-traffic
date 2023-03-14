package filter

import (
	"github.com/am-okalin/kit/tableconv"
	"time"
)

func YD2Trip(fname string) ([]Trip, error) {
	table, err := tableconv.Csv2Table(fname, Comma)
	if err != nil {
		return nil, err
	}
	m, rowLen := tableconv.ToM(table)
	list := make([]Trip, 0, rowLen)
	for i := 0; i < rowLen; i++ {
		if m["匹配标记"][i] == "进站抵消" || m["匹配标记"][i] == "单边进" || m["匹配标记"][i] == "单边出" {
			//log.Printf("error row[%d]: %s", i, m["匹配标记"][i])
			continue
		}
		tmp := Trip{
			TripId:     m["行程编号"][i],
			InTransId:  m["进站行程编号"][i],
			OutTransId: m["出站行程编号"][i],
			In: Trans{
				TransCode:   In,
				TicketId:    m["虚拟卡号"][i],
				Line:        m["IN_STATION_ID"][i][0:2],
				StationId:   m["IN_STATION_ID"][i],
				StationName: m["进站车站"][i],
				TransId:     m["进站行程编号"][i],
				TransTime:   time.Time{},
				Date:        "",
			},
			Out: Trans{
				TransCode:   Out,
				TicketId:    m["虚拟卡号"][i],
				Line:        m["OUT_STATION_ID"][i][0:2],
				StationId:   m["OUT_STATION_ID"][i],
				StationName: m["出站车站"][i],
				TransId:     m["出站行程编号"][i],
				TransTime:   time.Time{},
				Date:        "",
			},
		}
		tmp.In.TransTime, err = time.Parse("2006/1/2 5:04:05", m["进站时间"][i])
		if err != nil {
			return nil, err
		}
		tmp.In.Date = tmp.In.TransTime.Add(-1 * time.Hour).Format("060102")
		tmp.Out.TransTime, err = time.Parse("2006/1/2 5:04:05", m["出站时间"][i])
		if err != nil {
			return nil, err
		}
		tmp.Out.Date = tmp.Out.TransTime.Add(-1 * time.Hour).Format("060102")
		list = append(list, tmp)
	}
	return list, nil
}
