package zeros

import "math/rand"

func removeZeros(in []int) []int {
	headFound := false
	head, lastNonZeroIndex := 0, 0

	for idx, val := range in {
		if val == 0 {
			if !headFound {
				head += 1
			}
		} else {
			if !headFound {
				headFound = true
			}
			lastNonZeroIndex = idx
		}

		// if head is the full length of the array then the array is full of zeros and we can just
		// return an empty slice
		if head >= len(in) {
			return []int{}
		}
	}
	return in[head : lastNonZeroIndex+1]
}

func removeZerosCursors(in []int) []int {
	headFound, tailFound := false, false
	head, tail := 0, len(in)-1

	for {
		if in[head] == 0 {
			head += 1
		} else {
			headFound = true
		}
		if in[tail] == 0 {
			tail -= 1
		} else {
			tailFound = true
		}

		// both cursors found a value, can return
		if headFound && tailFound {
			return in[head : tail+1]
		}

		// head has exceeded tail which means we're all zeros, return empty slice
		if head >= tail {
			return []int{}
		}
	}
}

type data struct {
	Puzzles []puzzle
}

type puzzle struct {
	Problem  []int
	Solution []int
}

func genTestData(num, size int) data {
	puzzles := make([]puzzle, num)
	for i := 0; i < num; i++ {
		prob, sol := genPuzzle(size)
		puzzles[i] = puzzle{
			Problem:  prob,
			Solution: sol,
		}
	}
	return data{
		Puzzles: puzzles,
	}
}

func genPuzzle(size int) (problem []int, solution []int) {
	padFront := rand.Intn(size)
	padRear := rand.Intn(size)
	problem = make([]int, size)
	for i := 0; i < size; i++ {
		problem[i] = rand.Intn(10)
	}

	for i := 0; i < padFront; i++ {
		problem[i] = 0
	}

	for i := 0; i < padRear; i++ {
		j := len(problem) - 1 - i
		problem[j] = 0
	}

	solution = removeZerosCursors(problem)
	return
}
