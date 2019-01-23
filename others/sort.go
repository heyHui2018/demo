package main

import (
	"fmt"
	"sort"
	"strconv"
)

type SortStruct struct {
	list map[string]interface{}
}

type AscSort []SortStruct

func (s AscSort) Len() int {
	return len(s)
}
func (s AscSort) Less(i, j int) bool { //从大到小排序
	return s[i].list["updateTime"].(int) < s[j].list["updateTime"].(int)
}
func (s AscSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type DescSort []SortStruct

func (s DescSort) Len() int {
	return len(s)
}
func (s DescSort) Less(i, j int) bool { //从小到大排序
	return s[i].list["updateTime"].(int) > s[j].list["updateTime"].(int)
}
func (s DescSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	var result_list []SortStruct
	for ii := 0; ii < 5; ii++ {
		result_map := make(map[string]interface{})
		result_map["id"] = strconv.Itoa(ii)
		result_map["updateTime"] = ii * ii
		result := new(SortStruct)
		result.list = result_map
		result_list = append(result_list, *result)
	}
	fmt.Println("result_list原 = ", result_list)
	sort.Sort(DescSort(result_list))
	fmt.Println("refer Desc = ", result_list)
	sort.Sort(AscSort(result_list))
	fmt.Println("refer Asc = ", result_list)
}
