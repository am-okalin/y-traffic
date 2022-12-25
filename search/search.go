package search

import "github.com/am-okalin/y-traffic/filter"

const (
	TransData         = "../file/source/trans.csv"
	Analysis28Weekend = "../file/source/analysis_2_8_weekend.csv"
	Analysis28Working = "../file/source/analysis_2_8_working.csv"
	Analysis29        = "../file/source/analysis_2_9.csv"
	PrefixLine        = "../file/source/line/"
	PrefixLineDate    = "../file/source/linedate/"
)

func Trip2Trans(trips []filter.Trip) []filter.Trans {
	list := make([]filter.Trans, 0, len(trips)*2)
	for _, trip := range trips {
		list = append(list, trip.In)
		list = append(list, trip.Out)
	}
	return list
}

func Append2Trans(arrs ...[]filter.Trans) []filter.Trans {
	length := 0
	for _, arr := range arrs {
		length += len(arr)
	}
	list := make([]filter.Trans, 0, length)
	for i, _ := range arrs {
		list = append(arrs[i])
	}
	return list
}
