package week2

func findShortestSubArray(nums []int) int {
    horder := map[int]int{} // 统计各个数字出现的次数
    hstart := map[int]int{} // 总计各个数组出现开始的位置
    hend := map[int]int{} // 统计各个数组出现结束的位置
    // 度数
    order := 0
    for i,num := range nums {
        horder[num]++
        if horder[num] > order {
            order = horder[num]
        }
        if _,ok := hstart[num]; !ok { // 记录数字第一次出现的位置
            hstart[num] = i
        }
        hend[num] = i // 记录数字最后一次出现的位置
    }
    shortest := len(nums) // 最短子数组的长度
    for k,v := range horder { 
        if v == order { // 找到出现次数最多的数字
            if hend[k] - hstart[k] < shortest { // 计算子数组的长度
                shortest = hend[k] - hstart[k] + 1
            }
        }
    }
    return shortest
}