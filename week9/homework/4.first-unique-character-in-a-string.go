func firstUniqChar(s string) int {
    cntMap := map[rune]int{}
    for _,v := range s {
        cntMap[v]++
    }
    ans := len(s)
    for k,v := range cntMap {
        if v == 1 {
            ans = min(ans,strings.IndexRune(s,k))
        }
    }
    if ans == len(s) {
        return -1
    }
    return ans
}

func min(a,b int) int {
    if a <  b {
        return a
    }
    return b
}