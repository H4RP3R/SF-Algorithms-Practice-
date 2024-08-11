package main

import (
	"reflect"
	"testing"
)

func Test_qSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty array", args{arr: []int{}}, []int{}},
		{"sorted array", args{arr: []int{0, 1, 2, 3, 4, 5}}, []int{0, 1, 2, 3, 4, 5}},
		{"one value array", args{arr: []int{1}}, []int{1}},
		{"reversed array", args{arr: []int{3, 2, 1}}, []int{1, 2, 3}},
		{"unsorted array", args{arr: []int{5, 15, 2, 13, 7, 16, 10, 2}}, []int{2, 2, 5, 7, 10, 13, 15, 16}},
		{"unsorted array with negatives", args{arr: []int{1, 9, 7, 4, 6, 2, 1, 13, 22, -3, 12, 76}},
			[]int{-3, 1, 1, 2, 4, 6, 7, 9, 12, 13, 22, 76}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := qSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("qSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
