package main

import "testing"

func Test_maxScoreWords(t *testing.T) {
	type args struct {
		words   []string
		letters []byte
		score   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				words:   []string{"dog", "cat", "dad", "good"},
				letters: []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'},
				score:   []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: 23,
		},
		{
			name: "test2",
			args: args{
				words:   []string{"xxxz", "ax", "bx", "cx"},
				letters: []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'},
				score:   []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10},
			},
			want: 27,
		},
		{
			name: "test3",
			args: args{
				words:   []string{"leetcode"},
				letters: []byte{'l', 'e', 't', 'c', 'o', 'd'},
				score:   []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := maxScoreWords(tt.args.words, tt.args.letters, tt.args.score); got != tt.want {
				t.Errorf("maxScoreWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
