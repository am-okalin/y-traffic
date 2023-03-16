package biz

import (
	"fmt"
	"strconv"
	"strings"
)

// InitCsnm 断面客流量
func InitCsnm(pm map[string][]Vertex) map[string]int {
	m := make(map[string]int, 0)
	for se, vertexes := range pm {
		if len(vertexes) == 2 {
			m[se] = 0
		}
	}
	return m
}

func CsnMap(trips []Trip, stations []Station) map[string]int {
	pm := PathMap(stations)
	m := InitCsnm(pm)
	for _, trip := range trips {
		for i := 1; i < len(trip.Path); i++ {
			key := fmt.Sprintf("%s_%s", trip.Path[i-1], trip.Path[i])
			m[key]++
		}
	}
	return m
}

func CsnTable(trips []Trip, stationNames []string) [][]string {
	stations := Stations()
	m := CsnMap(trips, stations)
	nvm := NameVertexM(stations)
	set := Arrs2Set(stationNames)
	table := make([][]string, 0)
	table = append(table, []string{"line", "sid", "eid", "direction", "sn", "en", "num"})
	for key, num := range m {
		list := strings.Split(key, "_")
		sv, ev := nvm[list[0]], nvm[list[1]]
		//通过sv, ev交集算出线路
		lines := Intersection(LinesByIds(sv.Ids), LinesByIds(ev.Ids))
		if len(lines) > 1 {
			//fmt.Println(sv, ev) //todo::log
		}
		sid := IdByLine(lines[0], sv.Ids)
		eid := IdByLine(lines[0], ev.Ids)
		direction := strings.Compare(sid, eid)
		if len(stationNames) == 0 {
			table = append(table, []string{lines[0], sid, eid, strconv.Itoa(direction), sv.Name, ev.Name, strconv.Itoa(num)})
		} else if set[sv.Name] && set[ev.Name] {
			table = append(table, []string{lines[0], sid, eid, strconv.Itoa(direction), sv.Name, ev.Name, strconv.Itoa(num)})
		}
	}
	return table
}

func IdByLine(line string, ids []string) string {
	for _, id := range ids {
		if id[0:2] == line {
			return id
		}
	}
	panic(fmt.Sprintf("%s:%v", line, ids))
}

func LinesByIds(ids []string) []string {
	list := make([]string, len(ids))
	for i, id := range ids {
		list[i] = id[0:2]
	}
	return list
}

func Arrs2Set(arrs ...[]string) map[string]bool {
	m := make(map[string]bool, 0)
	for _, arr := range arrs {
		for _, s := range arr {
			m[s] = true
		}
	}
	return m
}

func Intersection(sets ...[]string) []string {
	m := make(map[string]int, 0)
	for _, set := range sets {
		for _, s := range set {
			m[s]++
		}
	}

	list := make([]string, 0, len(m))
	for s, i := range m {
		if i > 1 {
			list = append(list, s)
		}
	}

	return list
}

func Union(sets ...[]string) {

}

func Diff(subtract string, sets ...[]string) {

}
