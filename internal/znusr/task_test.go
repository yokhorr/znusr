package znusr

import "testing"

func TestMult(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Multiply positive numbers", 3, 5, 15},
		{"Multiply by zero", 5, 0, 0},
		{"Multiply zero", 0, 5, 0},
		{"Multiply by one", 5, 1, 5},
		{"Multiply one", 1, 5, 5},
		{"Multiply large numbers", 10, 10, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Mult(tt.a, tt.b)

			if got != tt.expected {
				t.Errorf("Mult(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
func TestSub(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Subtract smaller from larger", 10, 3, 7},
		{"Subtract larger from smaller", 3, 10, 0}, // Truncated to 0
		{"Subtract equal numbers", 5, 5, 0},
		{"Subtract zero", 10, 0, 10},
		{"Subtract from zero", 0, 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sub(tt.a, tt.b)

			if got != tt.expected {
				t.Errorf("Sub(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
