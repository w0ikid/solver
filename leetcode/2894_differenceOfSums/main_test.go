package main

import "testing"
func Test_differenceOfSums(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				n: 10,
				m: 3,
			},
			want: 19,
		},
		{
			name: "test2",
			args: args{
				n: 5,
				m: 6,
			},
			want: 15,
		},
		{
			name: "test3",
			args: args{
				n: 5,
				m: 1,
			},
			want: -15,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := differenceOfSums(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("differenceOfSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
