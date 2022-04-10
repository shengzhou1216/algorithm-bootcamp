func canJump(nums []int) bool {
    // 状态: 位置i,
    // f[i] 表示位置i能否到达
    n := len(nums)
    nums = append([]int{1},nums...)
    f := make([]bool,n+1)
    f[1] = true
    for i := 1; i <= n; i++ {
        for j := 1; j <= nums[i]; j++ {
            if i + j <= n {
                // 看从i开始，能够走到哪里，比较容易想
                f[i+j] = f[i+j] || f[i]
            }
        }
    }
    return f[n]
}