package zeromatrix

func ZeroMatrix(matrix [][]int32) {
	if len(matrix) == 0 || len(matrix) == 1 && len(matrix[0]) <= 1 {
		return
	}

	numRows := len(matrix)
	numCols := len(matrix[0])

	firstRowHasZero := false
	firstColHasZero := false

	for i := range max(numRows, numCols) {
		if i < numRows && matrix[i][0] == 0 {
			firstColHasZero = true
		}
		if i < numCols && matrix[0][i] == 0 {
			firstRowHasZero = true
		}
	}

	// set the left and top edges to be 0 when there is a 0 in their row/column
	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[row]); col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}

	// zero out any necessary rows
	for row := range numRows {
		if matrix[row][0] == 0 {
			for col := range matrix[row] {
				matrix[row][col] = 0
			}
		}
	}

	// zero out any necessary columns
	for col := range numCols {
		if matrix[0][col] == 0 {
			for row := range numRows {
				matrix[row][col] = 0
			}
		}
	}

	if firstRowHasZero {
		for col := range numCols {
			matrix[0][col] = 0
		}
	}

	if firstColHasZero {
		for row := range numRows {
			matrix[row][0] = 0
		}
	}
}
