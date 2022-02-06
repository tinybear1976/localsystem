package set

import (
	"bytes"
	"encoding/gob"
)

type Set4Int map[int]Empty

// 通过字符串切片创建字符串集合
func New_SetInt(intSlice []int) Set4Int {
	m := Set4Int{}
	for _, i := range intSlice {
		m[i] = Empty{}
	}
	return m
}

// 比较两个集合是否元素一致
func (set1 Set4Int) Compare(set2 Set4Int) bool {
	// 元素数量不同
	set1Len := len(set1)
	if set1Len != len(set2) {
		return false
	}
	for k := range set1 {
		if _, ok := set2[k]; !ok {
			return false
		}
	}
	return true
}

// 两个或多个集合交集
func (set1 Set4Int) Intersection(sets ...Set4Int) Set4Int {
	res := Set4Int{}
	for k := range set1 {
		res[k] = Empty{}
	}
	if len(sets) == 0 {
		return res
	}
	for _, set := range sets {
		if len(res) == 0 {
			break
		}
		for k := range res {
			if _, ok := set[k]; !ok {
				delete(res, k)
			}
		}
	}
	return res
}

// 两个或多个集合并集
func (set1 Set4Int) Union(sets ...Set4Int) (Set4Int, error) {
	res := Set4Int{}
	// 暂时忽略错误
	if len(set1) > 0 {
		err := deepCopy_int(&res, set1)
		if err != nil {
			return res, err
		}
	}
	for _, set := range sets {
		for k := range set {
			res[k] = Empty{}
		}
	}
	return res, nil
}

// 检查1个或多个字符串在集合中是否存在，只有有一个不符合就返回false，全部存在返回true。
// 如果入参没有传入任何内容，返回false
// 如果指定的集合没有任何内容返回false
func (set1 Set4Int) IsExist(keys ...int) bool {
	if len(set1) == 0 {
		return false
	}
	if len(keys) == 0 {
		return false
	}
	for _, s := range keys {
		if _, ok := set1[s]; !ok {
			return false
		}
	}
	return true
}

// 将给定的字符串按照在集合中出现、未出现进行分检返回
// 如果被检索集合不包含任何元素，则传入检索项全部返回non_exist
// 如果没有传入参数，则返回的 exist, non_exist分别都为空集合
func (set1 Set4Int) Categorizing(keys ...int) (exist, non_exist Set4Int) {
	exist = Set4Int{}
	non_exist = Set4Int{}
	if len(keys) == 0 {
		return
	}
	if len(set1) == 0 {
		non_exist = New_SetInt(keys)
		return
	}
	for _, k := range keys {
		if _, ok := set1[k]; ok {
			exist[k] = Empty{}
		} else {
			non_exist[k] = Empty{}
		}
	}
	return
}

// 将集合转为字符串切片
func (set1 Set4Int) ToSlice() []int {
	res := make([]int, 0)
	for k := range set1 {
		res = append(res, k)
	}
	return res
}

// 将字符串加入集合
func (set1 Set4Int) Add(keys ...int) {
	for _, s := range keys {
		set1[s] = Empty{}
	}
}

// 从集合中移除指定key
func (set1 Set4Int) Remove(keys ...int) {
	for _, s := range keys {
		delete(set1, s)
	}
}

func deepCopy_int(pdst *Set4Int, src Set4Int) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(pdst)
}
