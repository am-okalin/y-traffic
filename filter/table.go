package filter

import (
	"github.com/am-okalin/kit/tableconv"
	"time"
)

// Table2Trans 将二维数组转为对象
func Table2Trans(table [][]string) []Trans {
	m, rowLen := tableconv.ToM(table)
	list := make([]Trans, rowLen)
	for i := 0; i < rowLen; i++ {
		list[i].TransCode = m["TransCode"][i]
		list[i].TicketId = m["TicketId"][i]
		list[i].Line = m["Line"][i]
		list[i].StationId = m["StationId"][i]
		list[i].StationName = m["StationName"][i]
		list[i].TransId = m["TransId"][i]
		list[i].Date = m["TransDate"][i]
		list[i].TransTime, _ = time.Parse(TransTimeFormat, m["TransTime"][i])
	}
	return list
}
