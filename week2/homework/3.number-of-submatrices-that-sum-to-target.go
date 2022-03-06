package week2
func numSubmatrixSumTarget(matrix [][]int, target int) int {
    m := len(matrix)
    n := len(matrix[0])
    if m == 0 || n == 0 {
        return 0
    }
    ans := 0
    // 前缀和
    prefixSum := make([][]int,m+1)
    for i := 0; i < m+1; i++ {
        prefixSum[i] = make([]int,n+1)
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            // 每个位置的前缀和
            prefixSum[i][j] = prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1] + matrix[i-1][j-1]
        }
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            for p := 1; p <= i; p++ {
                for q := 1; q <= j; q++ {
                    if prefixSum[i][j] - prefixSum[i][q-1] - prefixSum[p-1][j] + prefixSum[p-1][q-1] == target {
                        ans++
                    }
                }
            }
        }
    }

    return ans
}