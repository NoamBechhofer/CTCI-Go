package rotate_matrix

func rotateLayer(matrix [][]int32, layer int) {
	N := len(matrix)
	first := layer
	last := N - 1 - layer
	steps := last - first

	for i := range steps {
		top := &matrix[first][first+i]
		right := &matrix[first+i][last]
		bottom := &matrix[last][last-i]
		left := &matrix[last-i][first]

		*right, *bottom, *left, *top = *top, *right, *bottom, *left
	}
}

// matrix must be square
func RotateMatrix(matrix [][]int32) {
	for i := 0; i < len(matrix)/2; i++ {
		rotateLayer(matrix, i)
	}
}
