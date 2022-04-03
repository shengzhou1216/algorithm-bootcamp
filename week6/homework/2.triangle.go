func minimumTotal(triangle [][]int) int {
    // 状态: 当前所处的位置（level ， index）
    // 目标： 当前所处位置的最小路径
    // 最优子结构： 当前所处位置的最小路径 f[i][j] = min(f[i-1][j],f[i-1][j-1]) + triangle[i][j]
    n := len(triangle)
    f := [][]int{}
    for  i := 0; i < n; i++ {
        t := []int{}
        for j := 0; j < n; j++ {
            t = append(t,math.MaxInt)
        }
        f = append(f,t)
    }
    f[0][0] = triangle[0][0]
    for i := 1; i < n; i++ {
       for j := 0; j < len(triangle[i]); j++ {
           // 状态转移
           if j - 1 >= 0 {
            f[i][j] = min(f[i-1][j],f[i-1][j-1]) + triangle[i][j]
           }else {
            f[i][j] = f[i-1][j] + triangle[i][j]
           }
       }
    }
    ans := math.MaxInt
    // 答案在最后一层
    for _,v := range f[len(f)-1] {
        ans = min(ans,v)
    }
    return ans
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}