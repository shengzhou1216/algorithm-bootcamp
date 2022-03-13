func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int,numCourses)
    inDeg := make([]int,numCourses)
    for _,p := range prerequisites {
        a := p[0]
        b := p[1]
        graph[b] = append(graph[b],a)
        inDeg[a]++
    }
    q := []int{}
    for i,v := range inDeg {
        if v == 0 {
            q = append(q,i)
        }
    }
    ans := []int{}
    for len(q) > 0 {
        n := q[0]
        q = q[1:]
        ans = append(ans,n)
        for _,v := range graph[n] {
            inDeg[v]--
            if inDeg[v] == 0 {
                q = append(q,v)
            }
        }
    }
    if len(ans) == numCourses {
        return ans
    }
    return []int{}
}