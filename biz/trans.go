package biz

import (
	"fmt"
	"github.com/am-okalin/kit/tableconv"
	"reflect"
	"sort"
	"strings"
	"time"
)

var transCodes = []string{In, Out}

// Tran 进出站
type Tran struct {
	TransCode   string    `gorm:"index"`      //交通类型码[21进站 22出站]
	TicketId    string    `gorm:"index"`      //票ID
	Line        string    `gorm:"index"`      //地铁线路
	StationId   string    `gorm:"index"`      //站台ID
	StationName string    `gorm:"index"`      //站台名称
	TransId     string    `gorm:"primaryKey"` //进出站ID
	TransTime   time.Time `gorm:"index"`      //进出站时间
	Date        string    `gorm:"index"`      //进出站日期(凌晨1点前属于前一天)
}

// SetTransId 进出类型+车站+时间+票=生成行程唯一ID
func (t *Tran) SetTransId() {
	//return hash(list...) //也可使用hash编码生成唯一id
	list := []string{t.TransCode, t.StationId, t.TransTime.Format(TransTimeFormat), t.TicketId}
	t.TransId = strings.Join(list, "_")
}

func (t *Tran) StrByField(groupBy string) string {
	return reflect.ValueOf(t).Elem().FieldByName(groupBy).String()
}

func TransGroup(list []Tran, groupBy string) map[string][]Tran {
	numM := make(map[string]int)
	for _, trans := range list {
		key := trans.StrByField(groupBy)
		numM[key]++
	}
	m := make(map[string][]Tran, len(numM))
	for i, trans := range list {
		key := trans.StrByField(groupBy)
		if m[key] == nil {
			m[key] = make([]Tran, 0, numM[key])
		}
		m[key] = append(m[key], list[i])
	}
	return m
}

// Trans2Table 将对象转换为二维数组
func Trans2Table(trans []Tran) [][]string {
	//todo::用反射处理并抽离公共包
	length := len(trans)
	table := make([][]string, 0, length+1)
	table = append(table, []string{
		"TransCode",
		"TicketId",
		"Line",
		"StationId",
		"StationName",
		"TransId",
		"TransTime",
		"Date",
		"CreateAt",
	})
	for i := 0; i < length; i++ {
		table = append(table, []string{
			trans[i].TransCode,
			trans[i].TicketId,
			trans[i].Line,
			trans[i].StationId,
			trans[i].StationName,
			trans[i].TransId,
			trans[i].TransTime.Format(TransTimeFormat),
			trans[i].Date,
			trans[i].TransTime.Format(time.RFC3339),
		})
	}
	return table
}

func Append2Trans(arrs ...[]Tran) []Tran {
	length := 0
	for _, arr := range arrs {
		length += len(arr)
	}
	list := make([]Tran, 0, length)
	for i, _ := range arrs {
		list = append(arrs[i])
	}
	return list
}

// Trans 获取trips
func Trans() []Tran {
	table, err := tableconv.Csv2Table(TransData, ',')
	if err != nil {
		panic(err)
	}
	m, rowLen := tableconv.ToM(table)
	list := make([]Tran, rowLen)
	for i := 0; i < rowLen; i++ {
		list[i].TransCode = m["TransCode"][i]
		list[i].TicketId = m["TicketId"][i]
		list[i].Line = m["Line"][i]
		list[i].StationId = m["StationId"][i]
		list[i].StationName = m["StationName"][i]
		list[i].TransId = m["TransId"][i]
		list[i].Date = m["Date"][i]
		list[i].TransTime, _ = time.Parse(TransTimeFormat, m["TransTime"][i])
	}
	return list
}

// TripsByTicket 有序trans数组转换为trip列表
func TripsByTicket(ticketId string, list []Tran) []Trip {
	if len(list) == 0 {
		return nil
	}

	l := len(list) / 2
	trips := make([]Trip, l)

	for i := 0; i < l; i++ {
		trips[i] = Trip{
			Date:       list[i*2].Date,
			TripId:     fmt.Sprintf("%s_%d", ticketId, i),
			TicketId:   list[i*2].TicketId,
			InTransId:  list[i*2].TransId,
			OutTransId: list[i*2+1].TransId,
			In:         list[i*2],
			Out:        list[i*2+1],
			Path:       nil,
		}
	}

	return trips
}

// Trans2Trips 将所有trans都转换为trip
func Trans2Trips(origin []Tran) ([]Trip, error) {
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
