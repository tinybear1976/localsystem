package algorithm

import (
	"reflect"
	"testing"
)

func TestInt_Sort_Asc(t *testing.T) {
	tests_ok := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "11个数升序,S: 0, E: 23",
			input: []int{2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10},
			want:  []int{0, 2, 3, 3, 4, 5, 8, 9, 9, 10, 23},
		},
		{
			name:  "3个数升序",
			input: []int{4, 23, 8},
			want:  []int{4, 8, 23},
		},
		{
			name:  "没有数据",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "1个数",
			input: []int{4},
			want:  []int{4},
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {
			data := IntSlice_MergeSort(item.input)
			got := data.Sort(true)
			if !reflect.DeepEqual(got, item.want) {
				t.Errorf("%v Sort(true)=%v , want %v", item.input, got, item.want)
			}
		})
	}
}

func TestInt_Sort_Desc(t *testing.T) {
	tests_ok := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "11个数降序",
			input: []int{70, 90, 60, 10, 74, 100, 20, 80, 40, 30, 50},
			want:  []int{100, 90, 80, 74, 70, 60, 50, 40, 30, 20, 10},
		},
		{
			name:  "3个数降序",
			input: []int{4, 23, 8},
			want:  []int{23, 8, 4},
		},
		{
			name:  "没有数据",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "1个数",
			input: []int{4},
			want:  []int{4},
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {
			data := IntSlice_MergeSort(item.input)
			got := data.Sort(false)
			if !reflect.DeepEqual(got, item.want) {
				t.Errorf("%v Sort(false)=%v , want %v", item.input, got, item.want)
			}
		})
	}
}

func BenchmarkSequentialFrequency(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	input := IntSlice_MergeSort{
		2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10,
		2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10,
		2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10,
		2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10,
		2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10,
		2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10,
		3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0, 10, 2, 9, 4, 23, 3, 8, 3, 0, 10, 2, 9, 4, 23, 3, 8, 5, 9, 3, 0,
	}
	for i := 0; i < b.N; i++ {
		input.Sort(true)
	}
}

func TestFloat64_Sort_Asc(t *testing.T) {
	tests_ok := []struct {
		name  string
		input []float64
		want  []float64
	}{
		{
			name:  "11个数降升序",
			input: []float64{2.2, 1.9, 44, 2.3, 30.3, 8.0, 53, 91, 23, 0.2, 0},
			want:  []float64{0, 0.2, 1.9, 2.2, 2.3, 8.0, 23, 30.3, 44, 53, 91},
		},
		{
			name:  "3个数升序",
			input: []float64{4, 23, 8},
			want:  []float64{4, 8, 23},
		},
		{
			name:  "没有数据",
			input: []float64{},
			want:  []float64{},
		},
		{
			name:  "1个数",
			input: []float64{4},
			want:  []float64{4},
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {
			data := Float64Slice_MergeSort(item.input)
			got := data.Sort(true)
			if !reflect.DeepEqual(got, item.want) {
				t.Errorf("%v Sort(true)=%v , want %v", item.input, got, item.want)
			}
		})
	}
}

func TestFloat64_Sort_Desc(t *testing.T) {
	tests_ok := []struct {
		name  string
		input []float64
		want  []float64
	}{
		{
			name:  "11个数降升序",
			input: []float64{2.2, 1.9, 44, 2.3, 30.3, 8.0, 53, 91, 23, 0.2, 0},
			want:  []float64{91, 53, 44, 30.3, 23, 8.0, 2.3, 2.2, 1.9, 0.2, 0},
		},
		{
			name:  "3个数升序",
			input: []float64{4, 23, 8},
			want:  []float64{23, 8, 4},
		},
		{
			name:  "没有数据",
			input: []float64{},
			want:  []float64{},
		},
		{
			name:  "1个数",
			input: []float64{4},
			want:  []float64{4},
		},
	}
	for _, item := range tests_ok {
		t.Run(item.name, func(t *testing.T) {
			data := Float64Slice_MergeSort(item.input)
			got := data.Sort(false)
			if !reflect.DeepEqual(got, item.want) {
				t.Errorf("%v Sort(true)=%v , want %v", item.input, got, item.want)
			}
		})
	}
}
