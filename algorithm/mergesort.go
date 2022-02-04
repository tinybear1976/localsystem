/*
  归并排序:
  支持int,float64
  排序方法支持升序与降序
*/
package algorithm

// var (
// 	ErrUnsupportedDataType = errors.New("unsupported data type")
// )

// int切片类型归并排序
type IntSlice_MergeSort []int

// 排序
// 参数 ascending 如果true 升序排列，否则降序
func (objs IntSlice_MergeSort) Sort(ascending bool) []int {
	res := mergeSort(objs, ascending)
	return res
}

func mergeSort(intList []int, isAsc bool) []int {
	length := len(intList)
	// 长度不够，直接退出
	if length <= 1 {
		return intList
	}

	// 先分
	middle := length / 2
	left := mergeSort(intList[:middle], isAsc)
	right := mergeSort(intList[middle:], isAsc)
	// 后治
	return merge(left, right, isAsc)
}

func merge(left, right []int, isAsc bool) []int {
	leftLen := len(left)
	rightLen := len(right)
	// 定义l_index,r _index指针指向left,right的,开始
	l_index, r_index := 0, 0
	// 定义临时切片
	temp := make([]int, 0)
	if isAsc {
		for l_index < leftLen && r_index < rightLen {
			if left[l_index] < right[r_index] {
				// 将左边的值放入temp
				temp = append(temp, left[l_index])
				l_index++
			} else {
				// 将右边的值放入temp
				temp = append(temp, right[r_index])
				r_index++
			}
		}
	} else {
		for l_index < leftLen && r_index < rightLen {
			if left[l_index] > right[r_index] {
				// 将左边的值放入temp
				temp = append(temp, left[l_index])
				l_index++
			} else {
				// 将右边的值放入temp
				temp = append(temp, right[r_index])
				r_index++
			}
		}
	}
	// 结束循环后会多有一个切片有余
	if l_index < leftLen {
		// 如果左边有余
		temp = append(temp, left[l_index:]...)
	} else if r_index < rightLen {
		// 如果右边有余
		temp = append(temp, right[r_index:]...)
	}
	return temp
}

//---------------------------------------------------------------------------------------------

// float64切片类型归并排序
type Float64Slice_MergeSort []float64

// 排序
// 参数 ascending 如果true 升序排列，否则降序
func (objs Float64Slice_MergeSort) Sort(ascending bool) []float64 {
	res := mergeSort_fromFloat64(objs, ascending)
	return res
}

func mergeSort_fromFloat64(lst []float64, isAsc bool) []float64 {
	length := len(lst)
	// 长度不够，直接退出
	if length <= 1 {
		return lst
	}

	// 先分
	middle := length / 2
	left := mergeSort_fromFloat64(lst[:middle], isAsc)
	right := mergeSort_fromFloat64(lst[middle:], isAsc)
	// 后治
	return merge_fromFloat64(left, right, isAsc)
}

func merge_fromFloat64(left, right []float64, isAsc bool) []float64 {
	leftLen := len(left)
	rightLen := len(right)
	// 定义l_index,r _index指针指向left,right的,开始
	l_index, r_index := 0, 0
	// 定义临时切片
	temp := make([]float64, 0)
	if isAsc {
		for l_index < leftLen && r_index < rightLen {
			if left[l_index] < right[r_index] {
				// 将左边的值放入temp
				temp = append(temp, left[l_index])
				l_index++
			} else {
				// 将右边的值放入temp
				temp = append(temp, right[r_index])
				r_index++
			}
		}
	} else {
		for l_index < leftLen && r_index < rightLen {
			if left[l_index] > right[r_index] {
				// 将左边的值放入temp
				temp = append(temp, left[l_index])
				l_index++
			} else {
				// 将右边的值放入temp
				temp = append(temp, right[r_index])
				r_index++
			}
		}
	}
	// 结束循环后会多有一个切片有余
	if l_index < leftLen {
		// 如果左边有余
		temp = append(temp, left[l_index:]...)
	} else if r_index < rightLen {
		// 如果右边有余
		temp = append(temp, right[r_index:]...)
	}
	return temp
}
