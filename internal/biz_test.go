package internal

import (
	"testing"
	"y-traffic/internal/biz"
)

func TestSaveTrans(t *testing.T) {
	list, err := biz.IC2Trans(biz.ICTest)
	if err != nil {
		t.Error(err)
	}
	err = biz.SaveTrans(list)
	if err != nil {
		t.Error(err)
	}
	t.Log("ok")
}

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
