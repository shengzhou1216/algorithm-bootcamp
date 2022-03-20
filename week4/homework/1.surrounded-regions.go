func solve(board [][]byte)  {
    m := len(board)
    n := len(board[0])
    dx := []int{-1,0,0,1}
    dy := []int{0,-1,1,0}
    visited := [][]bool{}
    for i := 0; i < m; i++ {
        visited = append(visited,make([]bool,n))
    }
    var bfs = func(sx,sy int) {
        q := [][]int{}
        q = append(q,[]int{sx,sy})
        visited[sx][sy] = true
        valid := true // 是否为合法区域
        candidate := [][]int{} // 候选坐标
        candidate = append(candidate,[]int{sx,sy})
        for len(q) > 0 {
            now := q[0]
            q = q[1:]
            x,y := now[0],now[1]
            for i := 0; i < 4; i++ {
                nx := x + dx[i]
                ny := y + dy[i]
                if nx < 0 || nx >= m || ny < 0 || ny >= n { // (sx,sy)周边的点越界了，说明(sx,sy)在边界上，与之相连的区域都不满足条件
                    valid = false
                    continue    
                }
                if board[nx][ny] == 'O' && !visited[nx][ny] {
                    q = append(q,[]int{nx,ny})
                    visited[nx][ny] = true
                    candidate = append(candidate,[]int{nx,ny})
                }
            }
        }
        if valid {
            for _,p := range candidate {
                board[p[0]][p[1]] = 'X'
            }
        }
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'O' && !visited[i][j]{
                bfs(i,j)
            }
        }
    }
}