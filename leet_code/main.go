package main

func main() {

	// // 1. Two Sum
	// nums := []int{3, 2, 4}
	// target := 6
	// fmt.Println(twoSum(nums, target))

	// 2. Add Two Numbers

}

func twoSum(nums []int, target int) []int {

	resultado := []int{}

	for i := 0; i < len(nums); i++ {
		for k := i + 1; k < len(nums); k++ {
			if nums[i]+nums[k] == target {
				resultado = append(resultado, i, k)
				return resultado
			}
		}

	}

	return []int{}
}

func twoSum1(nums []int, target int) []int {
	record := make(map[int]int)

	for i, j := range nums {
		complement := target - j
		if res, ok := record[complement]; ok {
			return []int{res, i}
		}
		record[j] = i
	}
	return []int{}
}
