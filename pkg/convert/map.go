package convert

import (
	"sort"
)

// 给Map排序并遍历
func MapSoftFor(m map[string]interface{}, callback func(k string, v interface{})) {
	var strs []string
	for k := range m {
		strs = append(strs, k)
	}
	sort.Strings(strs)
	// To perform the opertion you want
	for _, k := range strs {
		// fmt.Printf("%s\t%d\n", k, m[k])
		callback(k, m[k])
	}
}
