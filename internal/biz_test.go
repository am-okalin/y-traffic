package internal

import (
	"testing"
	"y-traffic/internal/biz"
)

func TestAnalysis(t *testing.T) {
	trans, err := biz.IC2Trans(biz.ICData)
	if err != nil {
		t.Error(err)
	}
	trip, err := biz.YD2Trip(biz.YDData)
	if err != nil {
		t.Error(err)
	}
	list := biz.Trip2Trans(trip)
	if err != nil {
		t.Error(err)
	}
	list = append(list, trans...)
	mi := biz.NewMinuteInterval(15)
	table := mi.Interval2Table(list)
	err = biz.Marshal("a.csv", table)
	if err != nil {
		t.Error(err)
	}
	t.Log("ok")
}

func TestTrans1(t *testing.T) {
	list1, err := biz.IC2Trans(biz.ICData)
	if err != nil {
		t.Error(err)
	}
	m1 := biz.TransGroup(list1, "TransDate")
	t.Log(len(list1), len(m1))

	for date, list2 := range m1 {
		m2 := biz.TransGroup(list2, "TicketId")
		t.Log(date, len(m2))
	}
	t.Log("done")
}
