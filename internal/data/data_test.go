package data

import (
	"testing"
)

func Test1(t *testing.T) {
	db, err := NewDb(NewDail(Source))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(db)
}
