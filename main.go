package main

import (
	"fmt"
	"znusr/internal/znusr"
)

func main() {
	tests := [][]int{
		[]int{1134472500, 2, 2},
	}

	for _, test := range tests {
		fmt.Println(znusr.Sublist(test))
	}
}
