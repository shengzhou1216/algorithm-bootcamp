package week2

func subarraySum(nums []int, k int) int {
    presum := make([]int,len(nums)+1)
    presum[0] = 0
    ans := 0
    h := map[int]int{0:1,}
    for i := 1; i < len(presum); i++ {
        presum[i] = presum[i-1] + nums[i-1]
        if cnt,ok := h[presum[i] - k]; ok {
            ans += cnt
        }
        h[presum[i]]++
    }
    return ans
}