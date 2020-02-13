/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return validate(root, math.MinInt64, math.MaxInt64)
}

// real validate func
func validate(node *TreeNode, min int, max int) bool {
    // null is true
    if node == nil {
        return true
    }

    // check min and max
    if node.Val <= min || node.Val >= max {
        return false
    }

    // recursive call
    return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
}s