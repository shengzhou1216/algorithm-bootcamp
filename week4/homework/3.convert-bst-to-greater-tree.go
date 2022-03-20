func convertBST(root *TreeNode) *TreeNode {
    var sum = 0
    var convert func(root *TreeNode) *TreeNode 
    convert = func(root *TreeNode) *TreeNode {
        if root == nil {
            return nil
        }
        root.Right = convert(root.Right)
        sum += root.Val
        root.Val = sum
        root.Left = convert(root.Left)
        return root
    }
    return convert(root)
}