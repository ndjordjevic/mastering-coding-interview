package arrays

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should reverse a string Nenad -> daneN",
			args: args{
				str: "Nenad",
			},
			want: "daneN",
		},
		{
			name: "Should reverse a string Hello world -> dlrow olleH",
			args: args{
				str: "Hello world",
			},
			want: "dlrow olleH",
		},
		{
			name: "Should reverse a string Tijana -> anajiT",
			args: args{
				str: "Tijana",
			},
			want: "anajiT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.str); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
