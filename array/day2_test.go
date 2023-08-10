package array

import (
	"reflect"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{nums: []int{-1, 2, 1, -4}, target: 1}, 2},
		{"case2", args{nums: []int{0, 0, 0}, target: 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"第一条 case 2 5 7 11 15", args{nums: []int{2, 7, 11, 15}, target: 9}, []int{0, 1}},
		{"第二条 case 3,2,4", args{nums: []int{3, 2, 4}, target: 6}, []int{1, 2}},
		{"第三条 case 3,3", args{nums: []int{3, 3}, target: 6}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum2(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
