package set

import (
	"reflect"
	"testing"
)

func Test_String_New(t *testing.T) {
	tests_ok := []struct {
		name  string
		input []string
		want  Set4String
	}{
		{
			name:  "4个正常项",
			input: []string{"abc", "def", "string", "test"},
			want: Set4String{
				"abc":    Empty{},
				"def":    Empty{},
				"string": Empty{},
				"test":   Empty{},
			},
		},
		{
			name:  "3项内容,但有1个重复",
			input: []string{"AAA", "BBB", "AAA"},
			want: Set4String{
				"AAA": Empty{},
				"BBB": Empty{},
			},
		},
		{
			name:  "没有数据",
			input: []string{},
			want:  Set4String{},
		},
		{
			name:  "4项内容,但有2个重复",
			input: []string{"AA", "CC", "AA", "CC"},
			want: Set4String{
				"AA": Empty{},
				"CC": Empty{},
			},
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {

			got := New(item.input)
			if !reflect.DeepEqual(got, item.want) {
				t.Errorf("%v => %v , want %v", item.input, got, item.want)
			}
		})
	}
}
func Test_String_Compare(t *testing.T) {
	tests_ok := []struct {
		name  string
		set1  Set4String
		input Set4String
		want  bool
	}{
		{
			name:  "true1",
			set1:  Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			input: Set4String{"def": Empty{}, "string": Empty{}, "test": Empty{}, "abc": Empty{}},
			want:  true,
		},
		{
			name:  "true2",
			set1:  Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			input: Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			want:  true,
		},
		{
			name:  "false1",
			set1:  Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			input: Set4String{"abc": Empty{}, "string": Empty{}, "test": Empty{}},
			want:  false,
		},
		{
			name:  "false2",
			set1:  Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			input: Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}, "d": Empty{}},
			want:  false,
		},
		{
			name:  "false3",
			set1:  Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			input: Set4String{},
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

func Test_String_Intersection(t *testing.T) {
	tests_ok := []struct {
		name  string
		set1  Set4String
		input []Set4String
		want  Set4String
	}{
		{
			name: "返回1项",
			set1: Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			input: []Set4String{
				{"def": Empty{}, "string": Empty{}, "test": Empty{}},
				{"abc": Empty{}, "string": Empty{}, "test": Empty{}},
				{"abc": Empty{}, "def": Empty{}, "string": Empty{}},
			},
			want: Set4String{
				"string": Empty{},
			},
		},
		{
			name:  "直接返回",
			set1:  Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			input: []Set4String{},
			want: Set4String{
				"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{},
			},
		},
		{
			name: "没有数据",
			set1: Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			input: []Set4String{
				{"def": Empty{}, "string": Empty{}, "test": Empty{}},
				{"abc": Empty{}, "string": Empty{}, "test": Empty{}},
				{"find": Empty{}, "foo": Empty{}, "bar": Empty{}},
				{"abc": Empty{}, "def": Empty{}, "string": Empty{}},
			},
			want: Set4String{},
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

func Test_String_Union(t *testing.T) {
	tests_ok := []struct {
		name  string
		set1  Set4String
		input []Set4String
		want  Set4String
	}{
		{
			name: "多项",
			set1: Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			input: []Set4String{
				{"AAA": Empty{}, "string": Empty{}, "123": Empty{}},
				{"abc": Empty{}, "BBB": Empty{}, "test": Empty{}},
				{"abc": Empty{}, "CCC": Empty{}},
			},
			want: Set4String{
				"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{},
				"123": Empty{}, "AAA": Empty{}, "BBB": Empty{}, "CCC": Empty{},
			},
		},
		{
			name: "基础集合没有数据",
			set1: Set4String{},
			input: []Set4String{
				{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
				{"123": Empty{}, "456": Empty{}, "789": Empty{}},
			},
			want: Set4String{
				"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{},
				"123": Empty{}, "456": Empty{}, "789": Empty{},
			},
		},
		{
			name:  "入参集合没有数据",
			set1:  Set4String{"123": Empty{}, "456": Empty{}, "789": Empty{}},
			input: []Set4String{},
			want: Set4String{
				"123": Empty{}, "456": Empty{}, "789": Empty{},
			},
		},
		{
			name:  "没有数据",
			set1:  Set4String{},
			input: []Set4String{},
			want:  Set4String{},
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

func Test_String_IsExist(t *testing.T) {
	tests_ok := []struct {
		name  string
		set1  Set4String
		input []string
		want  bool
	}{
		{
			name:  "存在1",
			set1:  Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			input: []string{"def", "test"},
			want:  true,
		},
		{
			name:  "存在2",
			set1:  Set4String{"123": Empty{}, "456": Empty{}, "789": Empty{}},
			input: []string{"123", "456", "789"},
			want:  true,
		},
		{
			name:  "入参没有数据",
			set1:  Set4String{"123": Empty{}, "456": Empty{}, "789": Empty{}},
			input: []string{},
			want:  false,
		},
		{
			name:  "没有数据",
			set1:  Set4String{},
			input: []string{"abc", "123"},
			want:  false,
		},
		{
			name:  "完全没有数据",
			set1:  Set4String{},
			input: []string{},
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
func Test_String_Categorizing(t *testing.T) {
	tests_ok := []struct {
		name           string
		set1           Set4String
		input          []string
		want_exist     Set4String
		want_non_exist Set4String
	}{
		{
			name:           "有返回1",
			set1:           Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			input:          []string{"def", "test"},
			want_exist:     Set4String{"def": Empty{}, "test": Empty{}},
			want_non_exist: Set4String{},
		},
		{
			name:           "有返回2",
			set1:           Set4String{"123": Empty{}, "456": Empty{}, "789": Empty{}},
			input:          []string{"123", "AAA", "789"},
			want_exist:     Set4String{"123": Empty{}, "789": Empty{}},
			want_non_exist: Set4String{"AAA": Empty{}},
		},
		{
			name:           "没有入参数据",
			set1:           Set4String{"123": Empty{}, "456": Empty{}, "789": Empty{}},
			input:          []string{},
			want_exist:     Set4String{},
			want_non_exist: Set4String{},
		},
		{
			name:           "没有集合数据",
			set1:           Set4String{},
			input:          []string{"abc", "123"},
			want_exist:     Set4String{},
			want_non_exist: Set4String{"abc": Empty{}, "123": Empty{}},
		},
		{
			name:           "完全没有数据",
			set1:           Set4String{},
			input:          []string{},
			want_exist:     Set4String{},
			want_non_exist: Set4String{},
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
func Test_String_ToStringSlice(t *testing.T) {
	tests_ok := []struct {
		name     string
		set1     Set4String
		want     []string
		want_len int
	}{
		{
			name:     "存在1",
			set1:     Set4String{"abc": Empty{}, "def": Empty{}, "string": Empty{}, "test": Empty{}},
			want:     []string{"abc", "def", "string", "test"},
			want_len: 4,
		},
		{
			name:     "存在2",
			set1:     Set4String{"123": Empty{}, "456": Empty{}, "789": Empty{}},
			want:     []string{"123", "456", "789"},
			want_len: 3,
		},
		{
			name:     "集合没有数据",
			set1:     Set4String{},
			want:     []string{},
			want_len: 0,
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {
			got := item.set1.ToStringSlice()

			if item.want_len != len(got) {
				t.Errorf("(input: len%d) %v => (got len: %d) %v , (want) %v", item.want_len, item.set1, len(got), got, item.want)
			}
		})
	}
}
