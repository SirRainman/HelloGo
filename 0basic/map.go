package main

import (
	"fmt"
	"reflect"
)

func main() {
	//定义map
	var countryCapitalMap map[string]string /*创建集合 */
	//初始化map，map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。
	countryCapitalMap = make(map[string]string)

	// countryCapitalMaps := map[string]string {
	//     {"Italy", "罗马"},
	//     {"Japan", "东京"},
	//     {"India", "新德里"},
	// }

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	fmt.Println(countryCapitalMap)

	// 删除城市
	delete(countryCapitalMap, "France")

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	fmt.Println(reflect.TypeOf(countryCapitalMap["xxx"]))

}
