package romannumeral

var romanNums = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToIntSlow(s string) int {
  var sum int = 0
	var lastNum int = 9999
	for _, value := range(s) {

		sum += romanNums[value]

		if lastNum < romanNums[value] {
			sum -= (2 * lastNum)
		}
		lastNum = romanNums[value]
	}

	return sum
}


func RomanToInt(s string) int {
	getRomanNum := func (c rune) int {
		switch c {
		case 'I': 
			return 1
		case 'V':
			return 5
		case 'X':
			return 10
		case 'L':
			return 50
		case 'C':
			return 100
		case 'D':
			return 500
		case 'M':
			return 1000
		}
		return 0
	}
  var sum int = 0
	var lastNum int = 9999

	for _, value := range(s) {
		num := getRomanNum(value)
		sum += num

		if lastNum < num {
			sum -= (2 * lastNum)
		}
		lastNum = num
	}

	return sum
}