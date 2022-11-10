package tableconv

import (
	"encoding/csv"
	"os"
)

//ToM 将二维数组转换为map
func ToM(table [][]string) (map[string][]string, int) {
	rl := len(table)
	cl := len(table[0])

	var m = make(map[string][]string, cl)
	for _, s := range table[0] {
		m[s] = make([]string, rl-1)
	}

	for row := 1; row < rl; row++ {
		for col := 0; col < cl; col++ {
			key := table[0][col]
			m[key][row-1] = table[row][col]
		}
	}

	return m, rl - 1
}

//ToObj 将二维数组转换为对象数组
func ToObj(table [][]string, v any) error {
	for i := 1; i < len(table); i++ {

	}
	return nil
}

func ToCsv(table [][]string, fname string) error {
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	for _, row := range table {
		err = w.Write(row)
		if err != nil {
			return err
		}
		w.Flush()
	}
	return nil
}
