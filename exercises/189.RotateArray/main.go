package main

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(nums, k)

}
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	temp := make([]int, 0, len(nums))
	temp = append(temp, nums[len(nums)-k:]...)
	temp = append(temp, nums[:len(nums)-k]...)

	copy(nums, temp)
}
