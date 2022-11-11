package biz

const (
	Comma = '\t'

	ICTest    = "../file/IC_test.txt"
	ICData    = "../file/IC20210816-22.txt"
	YDData    = "../file/YD20210816-22.txt"
	Stations  = "../file/stations.csv"
	TransData = "../file/trans.csv"
	TransTest = "../file/trans_test.csv"

	Analysis28Weekend = "../file/analysis_2_8_weekend.csv"
	Analysis28Working = "../file/analysis_2_8_working.csv"
	Analysis29        = "../file/analysis_2_9.csv"
	Analysis210       = "../file/analysis_2_10.csv"

	PrefixLine     = "../file/line/"
	PrefixLineDate = "../file/linedate/"
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
