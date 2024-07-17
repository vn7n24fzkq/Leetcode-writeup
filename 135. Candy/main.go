package main

func main() {
    ratings := []int{1, 0, 2}
    print(candy(ratings))
}

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

