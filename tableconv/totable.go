package tableconv

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
)

func InitTable(rl, cl int) [][]string {
	table := make([][]string, rl)
	for i := 0; i < rl; i++ {
		table[i] = make([]string, cl)
	}
	return table
}

// M2Table map转为二维数组
func M2Table(m map[string][]string) [][]string {
	var rl int
	for _, list := range m {
		if rl < len(list) {
			rl = len(list)
		}
	}
	table := InitTable(rl+1, len(m))

	ci := 0
	for key, list := range m {
		table[0][ci] = key
		for ri, val := range list {
			table[ri+1][ci] = val
		}
		ci++
	}
	return table
}

// Objs2Table 对象数组转为二维数组
func Objs2Table(objs any) ([][]string, error) {
	//校验：参数objs必须为[]Type
	vList := reflect.ValueOf(objs)
	rl := vList.Len()
	if !(vList.Kind() == reflect.Slice && rl != 0) {
		return nil, ObjError
	}
	vi := vList.Index(0)
	if vi.Kind() != reflect.Struct {
		return nil, ObjError
	}

	//初始化table
	cl := vi.NumField()
	table := InitTable(rl+1, cl)
	for i := 0; i < cl; i++ {
		table[0][i] = vi.Type().Field(i).Name
	}

	//写入数据
	for i := 0; i < rl; i++ {
		vi = vList.Index(i)
		for j := 0; j < vi.NumField(); j++ {
			vif := vi.Field(j)
			table[i+1][j] = fmt.Sprint(vif)
		}
	}
	return table, nil
}

// Csv2Table 从csv读取为二维数组 comma表示分割符号
func Csv2Table(fname string, comma rune) ([][]string, error) {
	//获取文件句柄
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
