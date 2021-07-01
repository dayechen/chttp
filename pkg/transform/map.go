package transform

import (
	"sort"
)

// 给Map排序并遍历
func MapSoftFor(m map[string]interface{}, callback func(k string)) {
	var strs []string
	for k := range m {
		strs = append(strs, k)
	}
	sort.Strings(strs)
	for _, k := range strs {
		callback(k)
	}
}
