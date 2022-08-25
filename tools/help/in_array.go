package help

// InArray 判断某一个值是否含在切片之中
func InArray(need interface{}, haystack interface{}) bool {
	switch key := need.(type) {
	case int:
		for _, item := range haystack.([]int) {
			if item == key {
				return true
			}
		}
	case string:
		for _, item := range haystack.([]string) {
			if item == key {
				return true
			}
		}
	case int64:
		for _, item := range haystack.([]int64) {
			if item == key {
				return true
			}
		}
	case float64:
		for _, item := range haystack.([]float64) {
			if item == key {
				return true
			}
		}

	default:
		return false
	}
	return false
}

// DelStringFromSlice 从字符串切片中删除指定元素
func DelStringFromSlice(source []string, item string) []string {
	var result []string
	for _, v := range source {
		if v != item {
			result = append(result, v)
		}
	}
	return result
}
