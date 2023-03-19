package biz

import (
	"testing"
	"time"
)

func TestOdm08160709past(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T09:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T07:00:00Z")
	trips := TripFilter(Trips(), "210816", start, end)
	odm := Odm(pastNames, trips)
	err := OdmCsv(pastNames, odm, Odm08160709past)
	if err != nil {
		t.Error(err)
	}
	t.Log("done")
}

func TestOdm08160(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T09:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T07:00:00Z")
	trips := TripFilter(Trips(), "210816", start, end)
	stationNames := StationNames()
	odm := Odm(stationNames, trips)
	err := OdmCsv(stationNames, odm, Odm0816)
	if err != nil {
		t.Error(err)
	}
	t.Log("done")
}
