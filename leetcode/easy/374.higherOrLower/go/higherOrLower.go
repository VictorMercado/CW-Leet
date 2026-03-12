package higherorlower

var Target int
func Guess(num int) int {
	if num > Target {
		return -1
	} else if num < Target {
		return 1
	}
	return 0
}

func GuessNumber(n int) int {
  if n == 1 {
		return 1
	}
	lower := 0
	higher := n
	guessNum := higher / 2
	guessResult := Guess(guessNum)

	for guessResult != 0 {
		switch guessResult {
			// too high
			case -1: {
				higher = guessNum
				guessNum = higher - ((higher - lower) / 2)
			}
			// too low
			case 1:
				lower = guessNum
				guessNum = (lower + higher) / 2 + 1
		}
		guessResult = Guess(guessNum)
	}

	return guessNum
}