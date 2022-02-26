package week1

func maximalRectangle(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    heights := make([][]int,m)
    res := 0
    for row,vals := range matrix {
        t := make([]int,n)
        for col,v := range vals {
            if row == 0 {
                if v == '1' {
                    t[col] = 1
                }
                continue
            }
            if v == '1' {
                t[col] = heights[row-1][col] + 1
            }else {
                t[col] = 0
            }
        }
        heights[row] = t
        res = max(largestRectangleArea(t),res)
    }
    return res
}

// 求解柱状图中的最大矩形面积
func largestRectangleArea(heights []int) int {
    heights = append([]int{0},heights...) // 保证s[len(s)-1]不会越界
    heights = append(heights,0) // 如果是单调增的数组，保证不会一直append
    maxArea := 0
    s := []int{}
    for i := 0; i < len(heights); i++ {
        for len(s) > 0 && heights[s[len(s)-1]] > heights[i] {
            currH := heights[s[len(s)-1]]
            s = s[:len(s)-1]
            currW := i - s[len(s)-1] - 1 // 添加了前置0，因此需要-1
            maxArea = max(maxArea, currH * currW ) 
        }
        s = append(s,i)
    }
    return maxArea
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}