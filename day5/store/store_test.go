package store

import "testing"

func TestRangeOverlaps(t *testing.T) {
	tests := []struct {
		name     string
		r1       Range
		r2       Range
		expected bool
	}{
		{name: "simple overlap", r1: Range{From: 1, To: 5}, r2: Range{From: 4, To: 8}, expected: true},
		{name: "no overlap", r1: Range{From: 1, To: 3}, r2: Range{From: 4, To: 6}, expected: false},
		{name: "touching at boundary", r1: Range{From: 1, To: 3}, r2: Range{From: 3, To: 5}, expected: true},
		{name: "contained within", r1: Range{From: 2, To: 10}, r2: Range{From: 4, To: 6}, expected: true},
		{name: "identical ranges", r1: Range{From: 5, To: 9}, r2: Range{From: 5, To: 9}, expected: true},
		{name: "reverse overlap", r1: Range{From: 7, To: 12}, r2: Range{From: 1, To: 8}, expected: true},
		{name: "single point overlap", r1: Range{From: 5, To: 5}, r2: Range{From: 5, To: 7}, expected: true},
		{name: "disjoint single points", r1: Range{From: 5, To: 5}, r2: Range{From: 6, To: 6}, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r1.Overlaps(tt.r2)
			if got != tt.expected {
				t.Fatalf("Overlaps(%v, %v) = %v, expected %v", tt.r1, tt.r2, got, tt.expected)
			}
		})
	}
}
