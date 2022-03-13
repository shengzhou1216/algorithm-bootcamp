package week3

func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }
    inorderMap := map[int]int{}
    for i,v := range inorder {
        inorderMap[v] = i
    }
    var build func(inorderLeft,inorderRight int) *TreeNode 
    build =  func(inorderLeft,inorderRight int) *TreeNode {
        if inorderLeft > inorderRight {
            return nil
        }
        val := postorder[len(postorder)-1]
        postorder = postorder[:len(postorder)-1]
        root := &TreeNode{Val: val}
        inorderRootIndex := inorderMap[val]
        root.Right = build(inorderRootIndex+1,inorderRight)
        root.Left = build(inorderLeft,inorderRootIndex-1)
        return root
    }
    return build(0,len(inorder)-1)
}