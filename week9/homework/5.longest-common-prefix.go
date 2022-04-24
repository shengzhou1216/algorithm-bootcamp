func longestCommonPrefix(strs []string) string {
    n := len(strs)
    if n == 0 {
        return ""
    }
    if n == 1 {
        return strs[0]
    }
    if n == 2 {
        return longestCommonPrefixOfTwo(strs[0],strs[1])
    }
    strs[1] = longestCommonPrefixOfTwo(strs[0],strs[1])
    strs = strs[1:]
    return longestCommonPrefix(strs)
}

func longestCommonPrefixOfTwo(str1,str2 string) string {
    i := 0
    for ; i < len(str1) && i < len(str2); i++ {
        if str1[i] != str2[i] {
            break
        }
    }
    return str1[:i]
}