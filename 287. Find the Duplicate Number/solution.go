package main

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
