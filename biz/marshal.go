package biz

import (
	"time"
)

// Table2Map 讲二维数组转换为 [key=>vals]
func Table2Map(table [][]string) (map[string][]string, int) {
	rowLen := len(table)
	colLen := len(table[0])

	var m = map[string][]string{}
	for _, s := range table[0] {
		m[s] = make([]string, rowLen-1)
	}

	for row := 1; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			key := table[0][col]
			m[key][row-1] = table[row][col]
		}
	}

	return m, rowLen - 1
}

// Trans2Table 将对象转换为二维数组
func Trans2Table(list []Trans) [][]string {
	//todo::用反射处理并抽离公共包
	length := len(list)
	table := make([][]string, 0, length+1)
	table = append(table, []string{
		"TransCode",
		"TicketId",
		"Line",
		"StationId",
		"StationName",
		"TransId",
		"TransTime",
		"TransDate",
	})
	for i := 0; i < length; i++ {
		table = append(table, []string{
			list[i].TransCode,
			list[i].TicketId,
			list[i].Line,
			list[i].StationId,
			list[i].StationName,
			list[i].TransId,
			list[i].TransTime.Format(TransTimeFormat),
			list[i].TransDate,
		})
	}
	return table
}

// Table2Trans 将二维数组转为对象
func Table2Trans(table [][]string) []Trans {
	m, rowLen := Table2Map(table)
	list := make([]Trans, rowLen)
	for i := 0; i < rowLen; i++ {
		list[i].TransCode = m["TransCode"][i]
		list[i].TicketId = m["TicketId"][i]
		list[i].Line = m["Line"][i]
		list[i].StationId = m["StationId"][i]
		list[i].TransId = m["TransId"][i]
		list[i].TransDate = m["TransDate"][i]
		list[i].TransTime, _ = time.Parse(TransTimeFormat, m["TransTime"][i])
	}
	return list
}
