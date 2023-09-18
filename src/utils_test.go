package main

import "testing"

func TestCopy2dSlice(t *testing.T) {
	src := make2DSlice[rune](10, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			src[i][j] = rune(i + j)
		}
	}
	result := copy2DSlice(src)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if result[i][j] != src[i][j] {
				t.Fatalf("copy2DSlice() operation failed: %v", result)
			}
		}
	}
}
