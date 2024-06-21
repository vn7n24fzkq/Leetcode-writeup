package main

func removeDuplicates(nums []int) int {
	count := 0
	j := 2
	for i := 0; i < len(nums); i++ {
        // We check j >= len(nums) first to avoid index out of range
		if (j >= len(nums) || nums[i] != nums[j]) {
			nums[count] = nums[i]
			count += 1
        }
		j++
	}
	return count
}

