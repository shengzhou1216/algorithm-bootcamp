package week3

func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    ans := [][]int{}
    choosen := []int{}
    used := make([]bool,n)
    var dfs func()
    dfs = func() {
        if len(choosen) == n {
            ans = append(ans,append([]int{},choosen...))
            return
        }
        for i := 0; i < n; i++ {
            if used[i] {
                continue
            }
            if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
                continue
            }
            choosen = append(choosen,nums[i])
            used[i] = true
            dfs()
            choosen = choosen[:len(choosen)-1]
            used[i] = false
        }
    }
    dfs()
    return ans
}