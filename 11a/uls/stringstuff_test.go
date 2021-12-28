package uls

import "testing"

func TestSilentAtoi(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{"0"},
			want: 0,
		},
		{
			name: "10",
			args: args{"10"},
			want: 10,
		},
		{
			name: "-10",
			args: args{"-10"},
			want: -10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SilentAtoi(tt.args.input); got != tt.want {
				t.Errorf("SilentAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
