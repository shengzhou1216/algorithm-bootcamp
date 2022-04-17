func numIslands(grid [][]byte) int {
    m := len(grid)
    n := len(grid[0])
    fa := make([]int,m * n + 1)
    for i := 0; i <= m * n; i++ {
        fa[i] = i
    }
    var find func(x int) int 
    find = func(x int) int {
        if x == fa[x] {
            return x
        }
        fa[x] = find(fa[x])
        return fa[x]
    }
    unionSet := func(x,y int) {
        x = find(x)
        y = find(y)
        if x != y {
            fa[x] = y
        }
    }
    
    num := func(i,j int) int {
        return i * n + j
    }

    dx := []int{-1,0,0,1}
    dy := []int{0,-1,1,0}

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '0' {
                continue
            }
            for k := 0; k < 4; k++ {
                ni := i + dx[k]
                nj := j + dy[k]
                if ni < 0 || nj < 0 || ni >= m || nj >= n {
                    continue
                }
                if grid[ni][nj] == '1' { // 将两个点连通起来
                    unionSet(num(ni,nj),num(i,j))
                }
            }
        }
    }
    ans := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j ++ {
            // 计算连通块的数量，即为岛屿数量
            if grid[i][j] == '1' && find(num(i,j)) == num(i,j){
                ans++
            }
        }
    }
    return ans
}