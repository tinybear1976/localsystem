package datastructure

type StringList []string

//Len()
func (list StringList) Len() int {
	return len(list)
}

//Less():按照字符串长短排序
func (list StringList) Less(i, j int) bool {
	return len(list[i]) < len(list[j])
}

//Swap()
func (list StringList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
