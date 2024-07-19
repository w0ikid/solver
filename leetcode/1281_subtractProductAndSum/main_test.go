package main

import "testing"

func Test_subtractProductAndSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := subtractProductAndSum(tt.args.n); got != tt.want {
				t.Errorf("subtractProductAndSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
