package hashtable

import "testing"

func Test_firstRecurringInt(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should return 3",
			args: args{
				arr: []int{3, 5, 7, 3, 9, 5},
			},
			want: 3,
		}, {
			name: "Should return -1",
			args: args{
				arr: []int{3, 5, 7, 8, 9, 11},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstRecurringInt(tt.args.arr); got != tt.want {
				t.Errorf("firstRecurringInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
