package week5
// 875. 爱吃香蕉的珂珂
func minEatingSpeed(piles []int, h int) int {
    // 根据h求速度k比较难；反向思考，假设k固定，能否在h小时内吃完
    left,right := 1,0
    for _,p := range piles {
        right = max(right,p)
    }
    for left < right {
        mid := (left + right) / 2
        if validate(piles,h,mid) {
            right = mid
        }else {
            left = mid + 1
        }
    }
    return right
}

// 以速度k能否在h小时内吃完所有的香蕉
func validate(piles []int,h int,k int) bool {
    cnt := 0 // 吃完香蕉所需要的小时数
    for _,p := range piles {
        if p % k == 0 {
            cnt += p / k
        }else {
            cnt += p / k + 1
        }
    }
    return cnt <= h
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}