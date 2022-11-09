package table

import (
	"strconv"
	"testing"
	"time"
)

const (
	to = "./to.csv"
)

type Param struct {
	P1, P2 string
}

type Obj struct {
	Id      int
	Name    string
	Arr     []int
	M       map[string]int
	CrateAt time.Time
	Param   *Param
}

func MockObjs() []Obj {
	length := 10
	objs := make([]Obj, length)
	for i := 0; i < length; i++ {
		objs[i] = Obj{
			Id:      i,
			Name:    "name" + strconv.Itoa(i),
			Arr:     []int{1, 2, 3},
			M:       map[string]int{"m1": 1, "m2": 2, "m3": 3},
			CrateAt: time.Now(),
			Param:   &Param{P1: "p1" + strconv.Itoa(i), P2: "p2" + strconv.Itoa(i)},
		}
	}
	return objs
}

func Test1(t *testing.T) {
	objs := MockObjs()
	table, err := Objs2Table(objs)
	t.Log(err)
	err = ToCsv(table, to)
	t.Log(err)
}

func Test2(t *testing.T) {
	objs := MockObjs()
	table, err := Objs2Table(objs)
	t.Log(err)

	v := make([]Obj, 10)
	err = ToObj(table, &v)
	t.Log(v, err)
}
