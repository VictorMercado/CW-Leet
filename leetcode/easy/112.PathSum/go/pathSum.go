package pathsum

import (
  // "fmt"
)
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
 }

// fmt.Printf("At Node: %d, Remaining: %d\n", root.Val, targetSum)
// fmt.Printf("Root left val: %d \n", root.Left.Val)
// fmt.Printf("Root right val: %d \n", root.Right.Val)

func hasPathSum(root *TreeNode, targetSum int) bool {
  if root == nil {
    return false
  }
 
  if root.Left == nil && root.Right == nil {
    return targetSum == root.Val
  }

  return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}