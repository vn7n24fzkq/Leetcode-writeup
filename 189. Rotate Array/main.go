func rotate(nums []int, k int)  {
    k = k % len(nums)
    copy_nums := make([]int, 0)
    copy(copy_nums, nums)
    for i := 0; i < len(nums); i++ {
        nums[(i+k)%len(nums)] = copy_nums[i]
    }
}
