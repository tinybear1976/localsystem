package set

import (
	"reflect"
	"testing"
)

func Test_Int_New(t *testing.T) {
	tests_ok := []struct {
		name  string
		input []int
		want  Set4Int
	}{
		{
			name:  "4个正常项",
			input: []int{4, 10, 16, 2},
			want: Set4Int{
				4:  Empty{},
				10: Empty{},
				16: Empty{},
				2:  Empty{},
			},
		},
		{
			name:  "3项内容,但有1个重复",
			input: []int{1, 3, 2, 3},
			want: Set4Int{
				1: Empty{},
				2: Empty{},
				3: Empty{},
			},
		},
		{
			name:  "没有数据",
			input: []int{},
			want:  Set4Int{},
		},
		{
			name:  "4项内容,但有2个重复",
			input: []int{1, 2, 1, 2},
			want: Set4Int{
				1: Empty{},
				2: Empty{},
			},
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {
			got := New_SetInt(item.input)
			if !reflect.DeepEqual(got, item.want) {
				t.Errorf("%v => %v , want %v", item.input, got, item.want)
			}
		})
	}
}
func Test_Int_Compare(t *testing.T) {
	tests_ok := []struct {
		name  string
		set1  Set4Int
		input Set4Int
		want  bool
	}{
		{
			name:  "true1",
			set1:  Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			input: Set4Int{3: Empty{}, 1: Empty{}, 4: Empty{}, 2: Empty{}},
			want:  true,
		},
		{
			name:  "true2",
			set1:  Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			input: Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			want:  true,
		},
		{
			name:  "false1",
			set1:  Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			input: Set4Int{1: Empty{}, 3: Empty{}, 4: Empty{}},
			want:  false,
		},
		{
			name:  "false2",
			set1:  Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			input: Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}, 5: Empty{}},
			want:  false,
		},
		{
			name:  "false3",
			set1:  Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			input: Set4Int{},
			want:  false,
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {
			got := item.set1.Compare(item.input)
			if got != item.want {
				t.Errorf("%v => %v , want %v", item.input, got, item.want)
			}
		})
	}
}

func Test_Int_Intersection(t *testing.T) {
	tests_ok := []struct {
		name  string
		set1  Set4Int
		input []Set4Int
		want  Set4Int
	}{
		{
			name: "返回1项",
			set1: Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			input: []Set4Int{
				{2: Empty{}, 3: Empty{}, 4: Empty{}},
				{1: Empty{}, 3: Empty{}, 4: Empty{}},
				{1: Empty{}, 2: Empty{}, 3: Empty{}},
			},
			want: Set4Int{
				3: Empty{},
			},
		},
		{
			name:  "直接返回",
			set1:  Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			input: []Set4Int{},
			want: Set4Int{
				1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{},
			},
		},
		{
			name: "没有数据",
			set1: Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			input: []Set4Int{
				{2: Empty{}, 3: Empty{}, 4: Empty{}},
				{1: Empty{}, 3: Empty{}, 4: Empty{}},
				{10: Empty{}, 11: Empty{}, 12: Empty{}},
				{1: Empty{}, 2: Empty{}, 3: Empty{}},
			},
			want: Set4Int{},
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {

			got := item.set1.Intersection(item.input...)
			if !reflect.DeepEqual(got, item.want) {
				t.Errorf("%v => %v , want %v", item.input, got, item.want)
			}
		})
	}
}

func Test_Int_Union(t *testing.T) {
	tests_ok := []struct {
		name  string
		set1  Set4Int
		input []Set4Int
		want  Set4Int
	}{
		{
			name: "多项",
			set1: Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			input: []Set4Int{
				{10: Empty{}, 3: Empty{}, 4: Empty{}},
				{1: Empty{}, 11: Empty{}, 2: Empty{}},
				{1: Empty{}, 12: Empty{}},
			},
			want: Set4Int{
				1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{},
				10: Empty{}, 11: Empty{}, 12: Empty{},
			},
		},
		{
			name: "基础集合没有数据",
			set1: Set4Int{},
			input: []Set4Int{
				{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
				{10: Empty{}, 11: Empty{}, 12: Empty{}},
			},
			want: Set4Int{
				1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{},
				10: Empty{}, 11: Empty{}, 12: Empty{},
			},
		},
		{
			name:  "入参集合没有数据",
			set1:  Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}},
			input: []Set4Int{},
			want: Set4Int{
				1: Empty{}, 2: Empty{}, 3: Empty{},
			},
		},
		{
			name:  "没有数据",
			set1:  Set4Int{},
			input: []Set4Int{},
			want:  Set4Int{},
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {
			got, err := item.set1.Union(item.input...)
			if err != nil {
				t.Errorf("!!!!!ERROR:%s (input) %v => (got) %v , (want) %v", err.Error(), item.input, got, item.want)
			}
			if !reflect.DeepEqual(got, item.want) {
				t.Errorf("(input) %v => (got) %v , (want) %v", item.input, got, item.want)
			}
		})
	}
}

func Test_Int_IsExist(t *testing.T) {
	tests_ok := []struct {
		name  string
		set1  Set4Int
		input []int
		want  bool
	}{
		{
			name:  "存在1",
			set1:  Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			input: []int{3, 1},
			want:  true,
		},
		{
			name:  "存在2",
			set1:  Set4Int{1: Empty{}, 3: Empty{}, 4: Empty{}},
			input: []int{1, 3, 4},
			want:  true,
		},
		{
			name:  "入参没有数据",
			set1:  Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}},
			input: []int{},
			want:  false,
		},
		{
			name:  "没有数据",
			set1:  Set4Int{},
			input: []int{1, 3},
			want:  false,
		},
		{
			name:  "完全没有数据",
			set1:  Set4Int{},
			input: []int{},
			want:  false,
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {
			got := item.set1.IsExist(item.input...)
			if got != item.want {
				t.Errorf("(input) %v => (got) %v , (want) %v", item.input, got, item.want)
			}
		})
	}
}
func Test_Int_Categorizing(t *testing.T) {
	tests_ok := []struct {
		name           string
		set1           Set4Int
		input          []int
		want_exist     Set4Int
		want_non_exist Set4Int
	}{
		{
			name:           "有返回1",
			set1:           Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			input:          []int{1, 30},
			want_exist:     Set4Int{1: Empty{}},
			want_non_exist: Set4Int{30: Empty{}},
		},
		{
			name:           "有返回2",
			set1:           Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}},
			input:          []int{1, 2, 3},
			want_exist:     Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}},
			want_non_exist: Set4Int{},
		},
		{
			name:           "没有入参数据",
			set1:           Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}},
			input:          []int{},
			want_exist:     Set4Int{},
			want_non_exist: Set4Int{},
		},
		{
			name:           "没有集合数据",
			set1:           Set4Int{},
			input:          []int{1, 2},
			want_exist:     Set4Int{},
			want_non_exist: Set4Int{1: Empty{}, 2: Empty{}},
		},
		{
			name:           "完全没有数据",
			set1:           Set4Int{},
			input:          []int{},
			want_exist:     Set4Int{},
			want_non_exist: Set4Int{},
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {
			got1, got2 := item.set1.Categorizing(item.input...)
			if !reflect.DeepEqual(got1, item.want_exist) || !reflect.DeepEqual(got2, item.want_non_exist) {
				t.Errorf("(input) %v => (got_exist) %v , (got_non_exist) %v\n\t (want_exist) %v, (want_non_exist) %v\n",
					item.input, got1, got2, item.want_exist, item.want_non_exist)
			}
		})
	}
}
func Test_Int_ToSlice(t *testing.T) {
	tests_ok := []struct {
		name     string
		set1     Set4Int
		want     []int
		want_len int
	}{
		{
			name:     "存在1",
			set1:     Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}, 4: Empty{}},
			want:     []int{1, 2, 3, 4},
			want_len: 4,
		},
		{
			name:     "存在2",
			set1:     Set4Int{1: Empty{}, 2: Empty{}, 3: Empty{}},
			want:     []int{1, 2, 3},
			want_len: 3,
		},
		{
			name:     "集合没有数据",
			set1:     Set4Int{},
			want:     []int{},
			want_len: 0,
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {
			got := item.set1.ToSlice()

			if item.want_len != len(got) {
				t.Errorf("(input: len%d) %v => (got len: %d) %v , (want) %v", item.want_len, item.set1, len(got), got, item.want)
			}
		})
	}
}
