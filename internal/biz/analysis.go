package biz

import (
	"fmt"
	"strconv"
	"time"
)

type MinuteInterval struct {
	m        map[int]string
	list     []int
	interval int
	length   int
}

func IntervalInt2Str(val int) string {
	h := strconv.Itoa(val / 60)
	if len(h) == 1 {
		h = "0" + h
	}
	m := strconv.Itoa(val % 60)
	if len(m) == 1 {
		m = "0" + m
	}
	return h + ":" + m
}

func NewMinuteInterval(interval int) *MinuteInterval {
	length := 24 * 60 / interval
	list := make([]int, length)
	m := make(map[int]string, length)
	for i, _ := range list {
		list[i] = i * interval
		m[list[i]] = fmt.Sprintf("%s-%s", IntervalInt2Str(list[i]), IntervalInt2Str(list[i]+interval))
	}
	return &MinuteInterval{
		m:        m,
		list:     list,
		interval: interval,
		length:   length,
	}
}

func (mi MinuteInterval) Len() int {
	return mi.length
}

func (mi MinuteInterval) Interval(t time.Time) int {
	val := t.Hour()*60 + t.Minute()
	for i, v := range mi.list {
		if val < v {
			return mi.list[i-1]
		}
	}
	return mi.list[mi.length-1]
}

//IntervalMap 时间间隔起时分钟=>客流数量
func (mi MinuteInterval) IntervalMap(list []Trans) map[int]int {
	m := make(map[int]int, mi.Len())
	for _, trans := range list {
		key := mi.Interval(trans.TransTime)
		m[key]++
	}
	return m
}

func (mi MinuteInterval) Interval2Table(list []Trans) [][]string {
	var i int
	m := mi.IntervalMap(list)
	table := make([][]string, len(m)+1)
	table[0] = []string{"interval", "number"}
	for k, v := range m {
		i++
		table[i] = []string{mi.m[k], strconv.Itoa(v)}
	}
	return table
}
