package main

func make2DSlice[T any](x, y int) [][]T {
	buf := make([][]T, x)
	for i := 0; i < x; i++ {
		buf[i] = make([]T, y)
	}
	return buf
}

func copy2DSlice[T any](data [][]T) [][]T {
	toReturn := make([][]T, len(data))
	for i := 0; i < len(data); i++ {
		toReturn[i] = make([]T, len(data[i]))
		for j := 0; j < len(data[i]); j++ {
			toReturn[i][j] = data[i][j]
		}
	}
	return toReturn
}
