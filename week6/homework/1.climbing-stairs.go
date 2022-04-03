package week6

func climbStairs(n int) int {
    // 状态： 剩余的楼梯数,
    // 
    dp := []int{0,1,2}
    for i := 3; i <= n; i++ {
        // 每次都可以选择爬1个或2个台阶
        dp = append(dp, dp[i-1] + dp[i-2])
    }
    return dp[n]
}