func findNumberOfLIS(nums []int) int {
    // 1. 求最长递增子序列的长度
    // 状态：以nums[i]结尾的最长子序列的长度
    n := len(nums)
    dp := make([]int,n)
    cnt := make([]int, n) //  状态：以nums[i]结尾的最长子序列的数量
    maxLen := 0
    ans := 0
    for i := 0; i < n; i++ {
        dp[i] = 1
        cnt[i] = 1
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] { // 满足递增条件
              if dp[j]+1 > dp[i] { 
                    dp[i] = dp[j] + 1
                    cnt[i] = cnt[j] // 重置计数
                } else if dp[j]+1 == dp[i] {
                    cnt[i] += cnt[j]
                }
            }
        }
        if dp[i] > maxLen {
            maxLen = dp[i]
            ans = cnt[i]
        } else if dp[i] == maxLen {
            ans += cnt[i]
        }
    }
    return ans
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}