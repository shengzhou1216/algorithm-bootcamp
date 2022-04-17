func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	fa := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if x == fa[x] {
			return x
		}
		fa[x] = find(fa[x])
		return fa[x]
	}
	unionSet := func(x, y int) {
		x = find(x)
		y = find(y)
		if x != y {
			fa[x] = y
		}
	}
	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		fa_a := find(a)
		fa_b := find(b)
		if fa_a != fa_b { // 当前的边[a,b]没有连通，那么添加[a,b]不会导致环出现
			unionSet(a, b) // 连通a,b
		} else { // 当前的[a,b]连通了，添加[a,b]会导致环出现
			return []int{a, b}
		}
	}
	return nil
}