package biz

import (
	"github.com/am-okalin/kit/tableconv"
	"time"
)

const (
	ICData            = "../file/source/IC20210816-22.txt"
	YDData            = "../file/source/YD20210816-22.txt"
	TransDataTest     = "../file/test/trans.csv"
	TripsData         = "../file/filter/trips.csv"
	TransData         = "../file/filter/trans.csv"
	Analysis28Weekend = "../file/source/analysis_2_8_weekend.csv"
	Analysis28Working = "../file/source/analysis_2_8_working.csv"
	Analysis29        = "../file/source/analysis_2_9.csv"
	PrefixLine        = "../file/source/line/"
	PrefixLineDate    = "../file/source/linedate/"
	StationsData      = "../file/station/stations.csv"
	OldStations       = "../file/station/stations_old.csv"
	OdmAll            = "../file/odm/all.csv"
	Odm1              = "../file/odm/1.csv"
	Odm2              = "../file/odm/2.csv"
	Odm3              = "../file/odm/3.csv"
	Csn08167090past   = "../file/csn/08167090past.csv"
	Csn08167090       = "../file/csn/08167090.csv"
	Csn08166370past   = "../file/csn/08166370past.csv"
	Csn08166370       = "../file/csn/08166370.csv"
	Csn0816           = "../file/csn/0816.csv"
	Csn0817           = "../file/csn/0817.csv"
	Csn0818           = "../file/csn/0818.csv"
	Csn0819           = "../file/csn/0819.csv"
	Csn0820           = "../file/csn/0820.csv"
)

const (
	In              = "21"
	Out             = "22"
	TransTimeFormat = "20060102150405"
	//TransTimeFormat = time.RFC3339
)

func IC2Trans(fname string) ([]Tran, error) {
	table, err := tableconv.Csv2Table(fname, '\t')
	if err != nil {
		return nil, err
	}
	m, rowLen := tableconv.ToM(table)
	list := make([]Tran, rowLen)
	stationM := IdStationM()
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

func YD2Trip(fname string) ([]Trip, error) {
	table, err := tableconv.Csv2Table(fname, '\t')
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
			In: Tran{
				TransCode:   In,
				TicketId:    m["虚拟卡号"][i],
				Line:        m["IN_STATION_ID"][i][0:2],
				StationId:   m["IN_STATION_ID"][i],
				StationName: m["进站车站"][i],
				TransId:     m["进站行程编号"][i],
				TransTime:   time.Time{},
				Date:        "",
			},
			Out: Tran{
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
