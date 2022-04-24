func numJewelsInStones(jewels string, stones string) int {
    h := map[rune]int{}
    for _,v := range jewels {
        h[v]++
    }
    ans := 0
    for _,v := range stones {
        if _,ok := h[v]; ok {
            ans++
        }
    }
    return ans
}