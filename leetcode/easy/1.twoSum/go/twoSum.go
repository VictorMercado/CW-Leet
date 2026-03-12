package twosum

func TwoSum(nums []int, target int) []int {
	result := []int{}
	tempMap :=  make(map[int]int)
	for _, value := range nums {
		tempMap[value] = tempMap[value] + 1
	}
	for key, value := range nums {
		if len(result) == 2 {
			break
		}
		complement := (target - value)
		mapCount, ok := tempMap[complement]
		if !ok {
			continue
		}
		if mapCount == 1 && complement == value {
			continue
		}
		result = append(result, key)
	}
	return result
}

// 0.171
func MorePerfTwoSum(nums []int, target int) []int {
	tempMap :=  make(map[int]int)
	for i, value := range nums {
		tempMap[value] = i
	}

	for i, value := range nums {
		complement := (target - value)
		if cMapIndex, ok := tempMap[complement]; ok && i != cMapIndex {
			return []int{i, cMapIndex}
		}
	}
	return []int{}
}

func EvenMorePerfTwoSum(nums []int, target int) []int{
	tempMap := make(map[int]int)
	
	for index, value := range nums {
		complement := target - value
		if compIndex, ok := tempMap[complement]; ok {
			return []int{compIndex,index}
		}
		tempMap[value] = index
	}

	return []int{}
}