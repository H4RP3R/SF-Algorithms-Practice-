package sortcheck

import "testing"

func Test_checkSliceIsSorted(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty slice", args{a: []int{}}, true},
		{"sorted slice", args{a: []int{1, 2, 3, 4, 5, 6, 7}}, true},
		{"unsorted slice", args{a: []int{1, 2, 3, 666, 4, 5, 0, 6, 7, 9}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSliceIsSorted(tt.args.a); got != tt.want {
				t.Errorf("checkSliceIsSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
