package znusr

import "testing"

func TestSublist(t *testing.T) {
	// Gödel encoding of the list [2, 3, 4, 5]:
	// 2^2 * 3^3 * 5^4 * 7^5 = 1_134_472_500.
	const L = 1_134_472_500

	tests := []struct {
		name string
		L    int
		I    int
		d    int
		want int
	}{
		{
			name: "entire list",
			L:    L, I: 0, d: 4,
			// [2, 3, 4, 5]: 2^2 * 3^3 * 5^4 * 7^5 = 1_134_472_500.
			want: 1_134_472_500,
		},
		{
			name: "first element",
			L:    L, I: 0, d: 1,
			// [2]: 2^2 = 4.
			want: 4,
		},
		{
			name: "second element",
			L:    L, I: 1, d: 1,
			// [3]: 2^3 = 8.
			want: 8,
		},
		{
			name: "last element",
			L:    L, I: 3, d: 1,
			// [5]: 2^5 = 32.
			want: 32,
		},
		{
			name: "first two elements",
			L:    L, I: 0, d: 2,
			// [2, 3]: 2^2 * 3^3 = 108.
			want: 108,
		},
		{
			name: "two middle elements",
			L:    L, I: 1, d: 2,
			// [3, 4]: 2^3 * 3^4 = 648.
			want: 648,
		},
		{
			name: "last two elements",
			L:    L, I: 2, d: 2,
			// [4, 5]: 2^4 * 3^5 = 3_888.
			want: 3_888,
		},
		{
			name: "three elements starting from the second",
			L:    L, I: 1, d: 3,
			// [3, 4, 5]: 2^3 * 3^4 * 5^5 = 2_025_000.
			want: 2_025_000,
		},
		{
			name: "empty sublist at the beginning",
			L:    L, I: 0, d: 0,
			// The empty list is encoded as 1.
			want: 1,
		},
		{
			name: "empty sublist at the end",
			L:    L, I: 4, d: 0,
			// The empty list is encoded as 1.
			want: 1,
		},
		{
			name: "empty list",
			L:    1, I: 0, d: 0,
			// The empty list is encoded as 1.
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sublist([]int{tt.L, tt.I, tt.d})

			if got != tt.want {
				t.Errorf(
					"Sublist(%d, %d, %d) = %d, want %d",
					tt.L, tt.I, tt.d, got, tt.want,
				)
			}
		})
	}
}
