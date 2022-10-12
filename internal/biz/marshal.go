package biz

import (
	"encoding/csv"
	"os"
)

// Marshal 从二维数组写入到csv文件中
func Marshal(fname string, table [][]string) error {
	//读取文件
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	//写入数据
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

func Map2Table(map[string][]string) [][]string {
	return nil
}

// Unmarshal 从csv文件中读取数据为二维数组
func Unmarshal(fname string, comma rune) ([][]string, error) {
	//读取文件
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//格式化读取为 [][]string
	r := csv.NewReader(file)
	r.Comma = comma
	table, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return table, nil
}

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
