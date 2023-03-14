package biz

import (
	"fmt"
	"github.com/am-okalin/kit/dijkstra"
	"math"
)

const (
	NavigationsData = "../file/graph/navigations.csv"
)

type Vertex struct {
	Vi   int
	Name string
	Ids  []string
}

type Navigation struct {
	sv     int
	ev     int
	path   []int
	weight float64
}

func NameM(objs []Obj) map[string]Vertex {
	length := VertexLen(objs)
	vertexes := Vertexes(objs, length)
	m := make(map[string]Vertex, length)
	for i, vertex := range vertexes {
		m[vertex.Name] = vertexes[i]
	}
	return m
}

func VertexM(objs []Obj) map[int]Vertex {
	length := VertexLen(objs)
	vertexes := Vertexes(objs, length)
	m := make(map[int]Vertex, length)
	for i, vertex := range vertexes {
		m[vertex.Vi] = vertexes[i]
	}
	return m
}

func VertexLen(objs []Obj) int {
	set := make(map[int]bool)
	for _, obj := range objs {
		set[obj.Vi] = true
	}
	return len(set)
}

func Vertexes(objs []Obj, length int) []Vertex {
	list := make([]Vertex, length)
	for _, obj := range objs {
		if list[obj.Vi].Ids != nil {
			list[obj.Vi].Ids = append(list[obj.Vi].Ids, obj.Id)
		} else {
			list[obj.Vi] = Vertex{
				Vi:   obj.Vi,
				Name: obj.Name,
				Ids:  []string{obj.Id},
			}
		}
	}
	return list
}

// InitGraph 初始化图
func InitGraph(objs []Obj, length int) dijkstra.Graph {
	// 初始化节点
	graph := dijkstra.NewSparseGraph(length)

	// 初始化边
	tmp := objs[0].Line
	for i := 1; i < len(objs); i++ {
		if objs[i].Line == tmp {
			graph.AddEdge(objs[i-1].Vi, objs[i].Vi, 1)
			graph.AddEdge(objs[i].Vi, objs[i-1].Vi, 1)
		} else {
			tmp = objs[i].Line
		}
	}

	//fmt.Println(graph)

	return graph
}

func Navigations(objs []Obj) []Navigation {
	length := VertexLen(objs)
	graph := InitGraph(objs, length)
	vertexes := Vertexes(objs, length)
	navigations := make([]Navigation, 0)
	for _, start := range vertexes {
		prev, dist := dijkstra.Dijkstra2(graph, start.Vi)
		for _, end := range vertexes {
			if start.Vi == end.Vi {
				continue
			}
			tmp := Navigation{sv: start.Vi, ev: end.Vi}
			dijkstra.GetPrev(end.Vi, prev, &tmp.path)
			if dist[end.Vi] == math.MaxFloat64 {
				tmp.weight = -1
			} else {
				tmp.weight = dist[end.Vi]
			}
			navigations = append(navigations, tmp)
		}
	}
	return navigations
}

// PathMap start_station_name + end_station_name => path
func PathMap(objs []Obj) map[string][]Vertex {
	vm := VertexM(objs)
	navigations := Navigations(objs)
	nl := len(navigations)

	m := make(map[string][]Vertex, nl)
	for _, na := range navigations {
		//key := fmt.Sprintf("%d_%d", na.sv, na.ev)
		key := fmt.Sprintf("%s_%s", vm[na.sv].Name, vm[na.ev].Name)
		m[key] = make([]Vertex, 0, len(na.path))
		for _, vi := range na.path {
			m[key] = append(m[key], vm[vi])
		}
	}
	return m
}
