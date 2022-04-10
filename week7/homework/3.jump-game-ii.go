func jump(nums []int) int {
    // 状态：位置i
    // f[i] 表示达到位置i的最少跳跃次数
    n := len(nums)
    f := make([]int,n)
    INF := int(1e5)
    for i := 0; i < n; i++ {
        f[i] = INF
    }
    f[0] = 0
    for i := 0; i < n; i++ {
        // 从i能够跳到[i,...,i+nums[i]]
        for j := 0; j <= nums[i]; j++ {
            if i + j < n {
                f[i+j] = min(f[i+j],f[i] + 1)
            }
        }
    }
    return f[n-1]
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}