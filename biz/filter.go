package biz

import "sort"

func FilterByGroup(oldList []Trans) []Trans {
	dateM := TransGroup(oldList, "TransDate")
	newList := make([]Trans, 0, len(oldList)/10)
	for _, dateList := range dateM {
		TicketIdM := TransGroup(dateList, "TicketId")
		for s, trans := range TicketIdM {
			TicketIdM[s] = Filter(trans)
			newList = append(newList, TicketIdM[s]...)
		}
	}
	return newList
}

func Filter(origin []Trans) []Trans {
	//对进出站表按时间排序
	sort.Slice(origin, func(i, j int) bool { return origin[i].TransTime.Before(origin[j].TransTime) })

	//过滤部分 transCode错误 进出站时间错误 的数据
	all := make([]Trans, 0, len(origin))
	for i, trans := range origin {
		if trans.TransCode != In && trans.TransCode != Out {
			continue
		}
		if trans.TransTime.Hour() >= 0 && trans.TransTime.Hour() < 6 {
			continue
		}
		all = append(all, origin[i])
	}

	//过滤进出站配对失败的数据
	var tc int
	right := make([]Trans, 0, len(all))
	for i, trans := range all {
		if trans.TransCode == transCodes[tc] {
			if trans.TransCode == Out && len(right) > 0 && trans.StationId == right[len(right)-1].StationId {
				//进出站抵消
				right = right[:len(right)-1]
			} else {
				//正常数据
				right = append(right, all[i])
			}
			tc = (tc + 1) % 2
		} else if trans.TransCode == In {
			//多次进站取最后一个(多次出站取第一个)
			right[len(right)-1] = all[i]
		}
	}

	//匹配为奇数则去除最后一个
	if len(right)%2 == 1 {
		right = right[:len(right)-1]
	}
	if len(right) == 0 {
		return nil
	}
	return right
}
