package biz

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"strings"
	"time"
	"y-traffic/tableconv"
)

const (
	In              = "21"
	Out             = "22"
	TransTimeFormat = "20060102150405"
)

var transCodes = []string{In, Out}

// Trans 进出站
type Trans struct {
	TransCode   string    //交通类型码[21进站 22出站]
	TicketId    string    //票ID
	Line        string    //地铁线路
	StationId   string    //站台ID
	StationName string    //站台名称
	TransId     string    //进出站ID
	TransTime   time.Time //进出站时间
	TransDate   string    //进出站日期(凌晨1点前属于前一天)
}

func hash(list ...string) string {
	h := sha256.New()
	for _, str := range list {
		h.Write([]byte(str))
	}
	b := h.Sum(nil)
	return hex.EncodeToString(b)
}

// SetTransId 进出类型+车站+时间+票=生成行程唯一ID
func (t *Trans) SetTransId() {
	list := []string{t.TransCode, t.StationId, t.TransTime.Format(TransTimeFormat), t.TicketId}
	//也可使用hash编码生成唯一id
	//return hash(list...)
	t.TransId = strings.Join(list, "_")
}

func (t Trans) StrByField(groupBy string) string {
	return reflect.ValueOf(t).FieldByName(groupBy).String()
}

func IC2Trans(fname string) ([]Trans, error) {
	table, err := tableconv.Csv2Table(fname, Comma)
	if err != nil {
		return nil, err
	}
	m, rowLen := tableconv.ToM(table)
	list := make([]Trans, rowLen)
	for i := 0; i < rowLen; i++ {
		transTime, err := time.Parse(TransTimeFormat, m["TXN_DATE"][i]+m["TXN_TIME"][i])
		if err != nil {
			return nil, err
		}
		list[i].TransCode = m["TRANS_CODE"][i]
		list[i].TicketId = m["TICKET_ID"][i]
		list[i].Line = m["TXN_STATION_ID"][i][0:2]
		list[i].StationId = m["TXN_STATION_ID"][i]
		list[i].StationName = StationNameById(m["TXN_STATION_ID"][i])
		list[i].TransTime = transTime
		list[i].TransDate = transTime.Add(-1 * time.Hour).Format("060102")
		list[i].SetTransId()
	}
	return list, nil
}

func TransGroup(list []Trans, groupBy string) map[string][]Trans {
	numM := make(map[string]int)
	for _, trans := range list {
		key := trans.StrByField(groupBy)
		numM[key]++
	}
	m := make(map[string][]Trans, len(numM))
	for i, trans := range list {
		key := trans.StrByField(groupBy)
		if m[key] == nil {
			m[key] = make([]Trans, 0, numM[key])
		}
		m[key] = append(m[key], list[i])
	}
	return m
}
