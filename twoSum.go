package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indices := []int{}
	// for loop through nums
	// add each index and check if it equals the `target`
	for i := 0; i < len(nums); i++ {
		// add up each index and compare to target
		// if nums[0] + nums [2] == target ==> [0,2]
		// if not keep iterating
		fmt.Println("i: ", nums[i])
		for j := i + 1; j < len(nums); j++ {
			fmt.Println("j: ", nums[j])
			if nums[i]+nums[j] == target {
				indices = []int{i, j}
			}
		}
	}
	return indices
}

/*
Optimal solution to review:

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int) // A mapping to store numbers and their indices
    for i, num := range nums {
        complement := target - num // Find the required number to reach the target
        if index, found := numMap[complement]; found {
            return []int{index, i} // Return indices of the complement and current number
        }
        numMap[num] = i // Store the number with its index
    }
    return nil // This line is never reached due to the problem guarantee
}

*/
