/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

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
