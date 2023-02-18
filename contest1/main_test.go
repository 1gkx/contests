package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {

	tests := []struct {
		source []int
		want int
	}{
		{ source: []int{5,1,4,2,3}, want: 1 },
		{ source: []int{7,5,4,2,3}, want: 2 },
		{ source: []int{5,4,3,2,0}, want: 0 },
	}

	for _, tt := range tests {
		got := min(tt.source)
		assert.Equal(t, got, tt.want, "they should be equal")
	}
}