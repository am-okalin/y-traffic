package biz

const (
	Comma = '\t'

	ICData    = "../file/source/IC20210816-22.txt"
	YDData    = "../file/source/YD20210816-22.txt"
	Stations  = "../file/source/stations.csv"
	TransData = "../file/source/trans.csv"

	Analysis28Weekend = "../file/source/analysis_2_8_weekend.csv"
	Analysis28Working = "../file/source/analysis_2_8_working.csv"
	Analysis29        = "../file/source/analysis_2_9.csv"
	Analysis210       = "../file/source/analysis_2_10.csv"

	PrefixLine     = "../file/source/line/"
	PrefixLineDate = "../file/source/linedate/"
)

func Trip2Trans(trips []Trip) []Trans {
	list := make([]Trans, 0, len(trips)*2)
	for _, trip := range trips {
		list = append(list, trip.In)
		list = append(list, trip.Out)
	}
	return list
}

func Append2Trans(arrs ...[]Trans) []Trans {
	length := 0
	for _, arr := range arrs {
		length += len(arr)
	}
	list := make([]Trans, 0, length)
	for i, _ := range arrs {
		list = append(arrs[i])
	}
	return list
}
