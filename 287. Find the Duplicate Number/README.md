# [287. Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)

## Problem Description

Given an array of integers `nums` containing `n + 1` integers where each integer is in the range `[1, n]` inclusive.

There is only one repeated number in `nums`, return this repeated number.

You must solve the problem without modifying the array `nums` and uses only constant extra space.

### Example 1:

```text
Input: nums = [1,3,4,2,2]
Output: 2
```

### Example 2:

```text
Input: nums = [3,1,3,4,2]
Output: 3
```

### Example 3:

```text
Input: nums = [1,1]
Output: 1
```

### Example 4:

```text
Input: nums = [1,1,2]
Output: 1
```

## Constraints:

- `1 <= n <= 10^5`
- `nums.length == n + 1`
- `1 <= nums[i] <= n`
- All the integers in `nums` appear only once except for precisely one integer which appears two or more times.

## Follow up:

- How can we prove that at least one duplicate number must exist in `nums`?
- Can you solve the problem in linear runtime complexity?

```
func findDuplicate(nums []int) int {
    const size = 100000
    record := make([]int64, (size/64)+1)

    for _, num := range nums {
        index := num / 64
        bit := int64(1 << (num % 64))
        if record[index]&bit != 0 {
            return num
        }
        record[index] |= bit
    }

    return -1
}
```


思路：

這邊使用int 來紀錄, 並且把每個 bit 都當成一個 record 的 flag, 藉此來壓縮空間的使用，
透過把每個數字除以 64 來取得 index, 因為我們用的是 int64
並且取得餘數來取得 bit 的位置，透過這樣的方式來記錄是否有重複的數字。

時間複雜度：O(n)
空間複雜度：O(1)
