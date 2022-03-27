// 1011. 在 D 天内送达包裹的能力
package week5

func shipWithinDays(weights []int, days int) int {
    left,right := 0,0
    for _,w := range weights {
        left = max(left,w)
        right += w
    }
    for left < right {
        mid := (left + right) / 2
        if validate(weights,days,mid) {
            // 求最低运载能力，要小的
            right = mid
        }else {
            left = mid + 1
        }
    }
    return right
}

// 运载能力为cap，能不能在D天内运送完
func validate(weights []int,days,cap int) bool {
    cnt := 0 // 需要的天数
    cur_weight := 0 // 当前装载重量
    for _,w := range weights {
        if cur_weight + w > cap { // 当前装不下了，放到下一天装
            cnt++
            cur_weight = 0
        }
        cur_weight += w
    }
    // 装载剩余的货物
    cnt++
    return cnt <= days
}


func max(a,b int)int {
    if a > b {
        return a
    }
    return b
}