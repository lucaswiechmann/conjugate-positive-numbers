package algorithm

import (
	"strconv"
)

func ConjugateNumbers(arrNumbers []string) []int {
	var firstElem = 0
	firstElem, _ = strconv.Atoi(arrNumbers[0])

	matrixFinal := make([][]int, firstElem)
	for i := range matrixFinal {
		matrixFinal[i] = make([]int, firstElem)
	}

	for i := 0; i < len(arrNumbers); i++ {
		curr, _ := strconv.Atoi(arrNumbers[i])
		for j := 0; j < curr; j++ {
			matrixFinal[j][i] = 1
		}
	}

	conjugateArrayNumbers := make([]int, firstElem)

	for i := 0; i < firstElem; i++ {
		count := 0
		for j := 0; j < firstElem; j++ {
			if matrixFinal[i][j] == 1 {
				count++
			}
		}
		conjugateArrayNumbers[i] = count
	}

	// fmt.Println(matrixFinal)

	return conjugateArrayNumbers
}
