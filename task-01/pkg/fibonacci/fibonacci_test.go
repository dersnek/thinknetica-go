package fibonacci

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Fib(0)",
			n:    0,
			want: 0,
		},
		{
			name: "Fib(1)",
			n:    1,
			want: 1,
		},
		{
			name: "Fib(5)",
			n:    5,
			want: 5,
		},
		{
			name: "Fib(11)",
			n:    11,
			want: 89,
		},
		{
			name: "Fib(20)",
			n:    20,
			want: 6765,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.n); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
