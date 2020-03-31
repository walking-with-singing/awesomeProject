package main

import (
	"testing"
)

func TestGetSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				a: 2,
				b: 3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}