package filter

import (
	"github.com/am-okalin/kit/tableconv"
	"github.com/am-okalin/y-traffic/station"
	"time"
)

func IC2Trans(fname string) ([]Trans, error) {
	table, err := tableconv.Csv2Table(fname, Comma)
	if err != nil {
		return nil, err
	}
	m, rowLen := tableconv.ToM(table)
	list := make([]Trans, rowLen)
	stationM := station.IdM()
	for i := 0; i < rowLen; i++ {
		transTime, err := time.Parse(TransTimeFormat, m["TXN_DATE"][i]+m["TXN_TIME"][i])
		if err != nil {
			return nil, err
		}
		list[i].TransCode = m["TRANS_CODE"][i]
		list[i].TicketId = m["TICKET_ID"][i]
		list[i].Line = m["TXN_STATION_ID"][i][0:2]
		list[i].StationId = m["TXN_STATION_ID"][i]
		list[i].StationName = stationM[m["TXN_STATION_ID"][i]].Name
		list[i].TransTime = transTime
		list[i].Date = transTime.Add(-1 * time.Hour).Format("060102")
		list[i].SetTransId()
	}
	return list, nil
}
