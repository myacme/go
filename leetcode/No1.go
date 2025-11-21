package main

func towSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v := target - nums[i]
		if index, ok := m[v]; ok {

			return []int{index, i}
		}
		m[nums[i]] = i
	}
	return []int{}
}

//func main() {
//	nums := []int{2, 7, 11, 15}
//	target := 17
//	fmt.Println(towSum(nums, target))
//}
