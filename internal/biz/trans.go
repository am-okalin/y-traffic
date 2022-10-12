package biz

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
	"y-traffic/internal/data"
)

//Trans 进出站表
type Trans struct {
	TransId     string    //进出站ID
	TicketId    string    //票ID
	Line        string    //地铁线路
	StationName string    //站台名称
	StationId   string    //站台ID
	TransCode   string    //交通类型码[21进站 22出站]
	TransTime   time.Time //进出站时间
}

func hash(list ...string) string {
	h := sha256.New()
	for _, str := range list {
		h.Write([]byte(str))
	}
	b := h.Sum(nil)
	return hex.EncodeToString(b)
}

//TransId 进出类型+车站+时间+票=生成行程唯一ID
func TransId(trans Trans) string {
	list := []string{trans.TransCode, trans.StationId, trans.TransTime.Format("20060102150405"), trans.TicketId}
	return strings.Join(list, "_")
	//todo::使用hash编码生成唯一id
	//return hash(list...)
}

func IC2Trans(fname string) ([]Trans, error) {
	table, err := Unmarshal(fname, '\t')
	if err != nil {
		return nil, err
	}
	m, rowLen := Table2Map(table)
	list := make([]Trans, rowLen)
	for i := 0; i < rowLen; i++ {
		list[i].TicketId = m["TICKET_ID"][i]
		list[i].TransCode = m["TRANS_CODE"][i]
		list[i].Line = m["TXN_STATION_ID"][i][0:2]
		list[i].StationId = m["TXN_STATION_ID"][i]
		list[i].StationName = StationNameById(m["TXN_STATION_ID"][i])
		list[i].TransTime, err = time.Parse("20060102150405", m["TXN_DATE"][i]+m["TXN_TIME"][i])
		if err != nil {
			return nil, err
		}
		list[i].TransId = TransId(list[i])
	}
	return list, nil
}

func SaveTrans(list []Trans) error {
	db, err := data.NewDb(data.NewDail(data.Source))
	if err != nil {
		return err
	}
	res := db.Create(&list)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
