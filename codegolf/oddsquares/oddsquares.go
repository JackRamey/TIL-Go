package oddsquares

func OddSquares(target int) int {
	switch target {
	case 0:
		return 0
	case 1:
		return 1
	}
	sum := 1
	for i := 3; i*i < target; i += 2 {
		sum += i * i
	}

	return sum
}
