func toLowerCase(s string) string {
    arr := []rune{}
    for _,v := range s {
        if v >= 'A' && v <= 'Z' {
            v = v - 'A' + 'a'
        }
        arr = append(arr,v)
    }
    return string(arr)
}