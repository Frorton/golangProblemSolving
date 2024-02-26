package TwoSum

func twoSum(nums []int, target int) []int {

	for i, v := range nums {
		for j, k := range nums {
			if i == j {
				continue
			}
			if (v + k) == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
