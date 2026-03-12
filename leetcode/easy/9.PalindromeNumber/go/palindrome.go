package palindrome

import "log"

func IsPalindromeSlow(x int) bool {
	if x < 0 {return false}
	if x == 0 {return true}
	var arr []int
	for x > 0 {
		arr = append(arr, x % 10)
		x = x /10
	}

	for i := 0; i < len(arr)/2; i++ {
		endIndex := (len(arr) - 1) - i
		value := arr[i]

		if value != arr[endIndex] {
			log.Println("no palidrome")
			return false
		}
	}

	return true
}

//x 1 2 3 4 5
//y 1 2 3 4

// x: 1 2 3 4 5
// y: 1 2 3 4 5 -> loop ends and we check
func IsPalindrome(x int) bool {
	if x < 0 || (x >= 10 && x % 10 == 0) {
		return false
	}

	y := 0
	for y < x {
		y = (y * 10) + (x % 10)

		x = x / 10
	}

	return y == x || (y / 10 == x)
}