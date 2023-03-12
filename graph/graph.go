package graph

import (
	"github.com/am-okalin/kit/dijkstra"
	"github.com/am-okalin/y-traffic/station"
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

// 求任意一点到任意一点的路径并存贮: 1.构件图 2.执行最短路径计算 3.记录结果集 sv ev path

func VertexLen(objs []station.Obj) int {
	set := make(map[int]bool)
	for _, obj := range objs {
		set[obj.Vi] = true
	}
	return len(set)
}

func Vertexes(objs []station.Obj, length int) []Vertex {

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
func InitGraph(objs []station.Obj, length int) dijkstra.Graph {
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
	return graph
}

func Navigations() []Navigation {
	objs := station.Objs()
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
