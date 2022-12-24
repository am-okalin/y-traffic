package pkg

import (
	"github.com/am-okalin/kit/tableconv"
	"github.com/am-okalin/y-traffic/filter"
	"time"
)

// Table2Trans 将二维数组转为对象
func Table2Trans(table [][]string) []filter.Trans {
	m, rowLen := tableconv.ToM(table)
	list := make([]filter.Trans, rowLen)
	for i := 0; i < rowLen; i++ {
		list[i].TransCode = m["TransCode"][i]
		list[i].TicketId = m["TicketId"][i]
		list[i].Line = m["Line"][i]
		list[i].StationId = m["StationId"][i]
		list[i].TransId = m["TransId"][i]
		list[i].TransDate = m["TransDate"][i]
		list[i].TransTime, _ = time.Parse(filter.TransTimeFormat, m["TransTime"][i])
	}
	return list
}

// Trans2Table 将对象转换为二维数组
func Trans2Table(list []filter.Trans) [][]string {
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
			list[i].TransTime.Format(filter.TransTimeFormat),
			list[i].TransDate,
		})
	}
	return table
}
