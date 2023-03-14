package graph

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
