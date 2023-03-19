package biz

import (
	"github.com/am-okalin/kit/tableconv"
	"testing"
	"time"
)

func TestCsn08167090past(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T09:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T07:00:00Z")
	trips := TripFilter(Trips(), "210816", start, end)
	table := CsnTable(trips, pastNames)
	err := tableconv.ToCsv(table, Csn08167090past)

	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn08167090(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T09:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T07:00:00Z")
	trips := TripFilter(Trips(), "210816", start, end)
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn08167090)

	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn08166370past(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T07:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T06:30:00Z")
	trips := TripFilter(Trips(), "210816", start, end)
	table := CsnTable(trips, pastNames)
	err := tableconv.ToCsv(table, Csn08166370past)

	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn08166370(t *testing.T) {
	var end, _ = time.Parse(time.RFC3339, "2021-08-16T07:00:00Z")
	var start, _ = time.Parse(time.RFC3339, "2021-08-16T06:30:00Z")
	trips := TripFilter(Trips(), "210816", start, end)
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn08166370)

	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn0816(t *testing.T) {
	trips := TripFilterDate(Trips(), "210816")
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn0816)
	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn0817(t *testing.T) {
	trips := TripFilterDate(Trips(), "210817")
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn0817)
	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn0818(t *testing.T) {
	trips := TripFilterDate(Trips(), "210818")
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn0818)
	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn0819(t *testing.T) {
	trips := TripFilterDate(Trips(), "210819")
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn0819)
	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}

func TestCsn0820(t *testing.T) {
	trips := TripFilterDate(Trips(), "210820")
	table := CsnTable(trips, []string{})
	err := tableconv.ToCsv(table, Csn0820)
	if err != nil {
		t.Log(err)
	}
	t.Log("done")
}
