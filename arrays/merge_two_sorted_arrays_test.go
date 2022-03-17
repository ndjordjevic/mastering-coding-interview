package arrays

import (
	"reflect"
	"testing"
)

func Test_mergeTwoSortedArrays(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should merge two sorted arrays",
			args: args{
				arr1: []int{0, 3, 4, 5, 18},
				arr2: []int{3, 4, 6, 30},
			},
			want: []int{0, 3, 3, 4, 4, 5, 6, 18, 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoSortedArrays(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
