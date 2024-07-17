# [150. Candy](https://leetcode.com/problems/candy/description/?envType=study-plan-v2&envId=top-interview-150)

## Description

There are `n` children standing in a line. Each child is assigned a rating value given in the integer array `ratings`.

You are giving candies to these children subjected to the following requirements:

- Each child must have at least one candy.
- Children with a higher rating get more candies than their neighbors.

Return the minimum number of candies you need to have to distribute the candies to the children.

## Example 1:

**Input:**
```
ratings = [1,0,2]
```
**Output:**
```
5
```
**Explanation:**
- You can allocate to the first, second, and third child with 2, 1, and 2 candies respectively.

## Example 2:

**Input:**
```
ratings = [1,2,2]
```
**Output:**
```
4
```
**Explanation:**
- You can allocate to the first, second, and third child with 1, 2, and 1 candies respectively.
- The third child gets 1 candy because it satisfies the above two conditions.

## Constraints:

- `n == ratings.length`
- `1 <= n <= 2 * 10^4`
- `0 <= ratings[i] <= 2 * 10^4`

## 思路
一開始我先把範例題目畫成折線圖，然後發現可以簡單遍歷兩次, 第一次從左到右，第二次從右到左，這樣就可以找到每個小孩應該得到的糖果數, 這邊需要另外的空間去紀錄每個小孩應該得到的糖果數，最後再將這些糖果數加總起來就是答案。

## Solution

```golang
func candy(ratings []int) int {
    record := make([]int, len(ratings))
    for i := 0; i < len(ratings); i++ {
        if i == 0 {
            record[i] = 1
        } else {
            if ratings[i] > ratings[i-1] {
                record[i] = record[i-1] + 1
            } else {
                record[i] = 1
            }
        }
            
    }
    for i := len(ratings) - 1; i >= 0; i-- {
        if i != len(ratings) - 1 {
            if ratings[i] > ratings[i+1] {
                record[i] = max(record[i], record[i+1] + 1)
            }
        }
    }
    sum := 0
    for _,i := range record{
        sum += i
    }
    return sum
}
```

- Time Complexity:  O(3N) but we can say O(N)
- Space Complexity: O(N), the space used by the record array.
