package str

// ContainsString 判断一个字符串是否在在于字符串数组元素中
// 如果存在那么会返回 true并且返回该元素在数组集合中的索引
// 如果不存在则会返回 false并且此时index则为-1
func ContainsString(arr []string, target string) (bool, int) {
	for index, item := range arr {
		if item == target {
			return true, index
		}
	}
	return false, -1
}
