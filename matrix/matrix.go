package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	n_cols := len(strings.Split(rows[0], " "))

	matrix := make(Matrix, len(rows))

	for i, row := range rows {
		cols := strings.Split(strings.Trim(row, " "), " ")
		if len(cols) != n_cols {
			return nil, fmt.Errorf("invalid number in column, got: %d expected: %d", len(cols), n_cols)
		}

		for _, num_s := range cols {
			num, err := strconv.Atoi(num_s)
			if err != nil {
				return nil, fmt.Errorf("invalid digit: %s", num_s)
			}
			matrix[i] = append(matrix[i], num)
		}
	}

	return matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cols := make(Matrix, len(m[0]))
	for _, row := range m {
		for i := range row { // for every row we add to the column i
			cols[i] = append(cols[i], row[i])
		}
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	rows := make(Matrix, len(m))
	for i := range m { // we cannot copy the entire matrix to the new one, because it will use same mem addresses, need to copy one by one
		new := make([]int, len(m[i])) //we make a empty array in which we will copy the row to
		copy(new, m[i])
		rows[i] = new
	}
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	if row > len(m)-1 || col > len(m[0])-1 || row < 0 || col < 0 {
		return false
	}
	m[row][col] = val
	return true
}

/*
BenchmarkNew-8            846014              1510 ns/op            1040 B/op         23 allocs/op
BenchmarkRows-8          6862144               183.0 ns/op           192 B/op          5 allocs/op
BenchmarkCols-8          1708687               667.2 ns/op           408 B/op         16 allocs/op
PASS
ok      matrix  4.584s
*/
