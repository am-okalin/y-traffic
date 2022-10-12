package biz

func Trip2Trans(trips []Trip) []Trans {
	list := make([]Trans, 0, len(trips)*2)
	for _, trip := range trips {
		list = append(list, trip.In)
		list = append(list, trip.Out)
	}
	return list
}
