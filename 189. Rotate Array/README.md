# [Rotate Array](https://leetcode.com/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150)

## Problem Description

Given an integer array `nums`, rotate the array to the right by `k` steps, where `k` is non-negative.

### Example 1:

**Input**: `nums = [1,2,3,4,5,6,7]`, `k = 3`  
**Output**: `[5,6,7,1,2,3,4]`  
**Explanation**:  
rotate 1 steps to the right: `[7,1,2,3,4,5,6]`  
rotate 2 steps to the right: `[6,7,1,2,3,4,5]`  
rotate 3 steps to the right: `[5,6,7,1,2,3,4]`

### Example 2:

**Input**: `nums = [-1,-100,3,99]`, `k = 2`  
**Output**: `[3,99,-1,-100]`  
**Explanation**:  
rotate 1 steps to the right: `[99,-1,-100,3]`  
rotate 2 steps to the right: `[3,99,-1,-100]`

### Constraints:

- `1 <= nums.length <= 10^5`
- `-2^31 <= nums[i] <= 2^31 - 1`
- `0 <= k <= 10^5`

## Follow up:

- Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
- Could you do it in-place with `O(1)` extra space?


## 思路：
因為最終會有一個 rotate 的固定值, 所以我們先將 k 取選轉次數的餘數, 這樣可以找到最終的陣列位移的位置。
接著我們將原本的陣列複製一份, 並將複製的陣列的值依照位移的規則放回原本的陣列中。


## Solution

```golang
func rotate(nums []int, k int)  {
    k = k % len(nums)
    copy_nums := make([]int, 0)
    copy(copy_nums, nums)
    for i := 0; i < len(nums); i++ {
        nums[(i+k)%len(nums)] = copy_nums[i]
    }
}
```

Time complexity: O(n), we need to copy the original array to a new array, and then copy the new array to the original array, so the time complexity is O(n).
Space complexity: O(n), we need to create a new array to store the original array.

## 延伸討論
Space Complexity: O(1) 的解法, 可以預先開一個足夠大的陣列來放 copy 的值, 這樣對於輸入來說, 會是 O(1) 的空間複雜度, 或是直接在原本的陣列上進行位移, 這樣也是 O(1) 的空間複雜度。
