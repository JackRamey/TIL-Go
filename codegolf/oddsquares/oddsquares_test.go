package oddsquares

import "testing"

func Test_oddSquares(t *testing.T) {
	tests := []struct {
		target int
		want   int
	}{
		{
			target: 0,
			want:   0,
		},
		{
			target: 1,
			want:   1,
		},
		{
			target: 9,
			want:   1,
		},
		{
			target: 10,
			want:   10,
		},
		{
			target: 44,
			want:   35,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := OddSquares(tt.target); got != tt.want {
				t.Errorf("OddSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
