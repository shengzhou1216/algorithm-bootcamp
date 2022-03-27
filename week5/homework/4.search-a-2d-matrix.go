package week5
// 74. 搜索二维矩阵

func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    row_start,row_end := 0,m-1
    // 1. 找到target所在的行
    for row_start < row_end {
        row_mid := (row_start + row_end + 1) / 2
        start,end := matrix[row_mid][0],matrix[row_mid][n-1]
        if target == start || target == end {
            return true
        }
        if target < start {
            row_end = row_mid - 1
        }else if target > end {
            row_start = row_mid + 1
        }else {
            // start < target < end
            row_end = row_mid
            break
        }
    }
    // 2. 找到target 所在的列
    col_start,col_end := 0,n-1
    for col_start <= col_end {
        col_mid := (col_start + col_end) / 2
        if matrix[row_end][col_mid] == target {
            return true
        }
        if matrix[row_end][col_mid] < target {
            col_start = col_mid + 1
        }else {
            col_end = col_mid - 1
        }
    }
    return false
}