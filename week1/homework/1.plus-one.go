package week1

func plusOne(digits []int) []int {
    for i := len(digits)-1; i >= 0; i-- {
        if digits[i] == 9 {
            digits[i] = 0
        }else {
            digits[i]++
            break
        }
    }
    // 第一个数为0,则添加一个位
    if digits[0] == 0 {
        digits = append([]int{1},digits...)
    }
    return digits
}