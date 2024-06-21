[80. Remove Duplicates from Sorted Array II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii)

Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.

Custom Judge:

The judge will test your solution with the following code:

```
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```
If all assertions pass, then your solution will be accepted.

 

Example 1:

```
Input: nums = [1,1,1,2,2,3]
Output: 5, nums = [1,1,2,2,3,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 1, 1, 2, 2 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:

Input: nums = [0,0,1,1,1,1,2,3,3]
Output: 7, nums = [0,0,1,1,2,3,3,_,_]
Explanation: Your function should return k = 7, with the first seven elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```
 

Constraints:

1 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums is sorted in non-decreasing order.


------

以下為解題心得：

思路：題目要求不能另外 allocate 一個新的 array，所以只能在原本的 array 上面做修改。這題的解法是用兩個指針，一個指針用來遍歷整個 array，另一個指針用來指向最後一個不重複的元素。當遍歷到的元素與指向的元素不同時，就把遍歷到的元素放到指向的元素的下一個位置。這樣就可以達到只保留每個元素最多出現兩次的效果。

依照這個思路，我們可以用一個變數 count 來記錄目前遍歷到的元素出現的次數，當 count 大於等於 2 時，就不要把元素放到指向的元素的下一個位置。這樣就可以達到只保留每個元素最多出現兩次的效果。

時間複雜度：因為至少會需要遍歷一次整個 array，所以預期的時間複雜度至少是 O(n)。


以下為 array 比較的方式
```
[1,1,1,2,2,3]
    [1,1,1,2,2,3]
```

```go
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
```

tips: 這邊先檢查 j >= len(nums) 是為了避免 index out of range 的問題，因為我們在檢查 nums[i] != nums[j] 之前，j 有可能已經超過了 nums 的長度。
