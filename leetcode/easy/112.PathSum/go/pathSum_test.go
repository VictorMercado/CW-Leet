package pathsum

import "testing"

// Helper function to build a tree from a slice (LeetCode style)
func buildTree(nodes []interface{}) *TreeNode {
	if len(nodes) == 0 || nodes[0] == nil {
		return nil
	}

	root := &TreeNode{Val: nodes[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nodes) {
		curr := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(nodes) && nodes[i] != nil {
			curr.Left = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, curr.Left)
		}
		i++

		// Right child
		if i < len(nodes) && nodes[i] != nil {
			curr.Right = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, curr.Right)
		}
		i++
	}

	return root
}

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		root      []interface{}
		targetSum int
		want      bool
	}{
		{
			name:      "Example 1",
			root:      []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1},
			targetSum: 22,
			want:      true,
		},
		{
			name:      "Example 2",
			root:      []interface{}{1, 2, 3},
			targetSum: 5,
			want:      false,
		},
		{
			name:      "Example 3",
			root:      []interface{}{},
			targetSum: 0,
			want:      false,
		},
		{
			name:      "Single Node Match",
			root:      []interface{}{1},
			targetSum: 1,
			want:      true,
		},
		{
			name:      "Single Node Mismatch",
			root:      []interface{}{1},
			targetSum: 2,
			want:      false,
		},
		{
			name:      "Target Sum Zero with Zero Nodes",
			root:      []interface{}{0, 0, 0},
			targetSum: 0,
			want:      true,
		},
		{
			name:      "Left Skewed Success",
			root:      []interface{}{1, 2, nil, 3, nil, 4, nil},
			targetSum: 10,
			want:      true,
		},
		{
			name:      "Negative Path",
			root:      []interface{}{5, -2, nil, 1, nil},
			targetSum: 4,
			want:      true, // 5 + -2 + 1 = 4
		},
		{
			name:      "Not a Leaf Trap",
			root:      []interface{}{1, 9, nil, 5, nil},
			targetSum: 10,
			want:      false, // 1+9 is 10, but 9 is not a leaf
		},
		{
			name: "Some negatives",
			root: []interface{}{1,-2,-3,1,3,-2,nil,-1},
			// 1, -2, 1 3, -1
			targetSum: -1,
			want: true,
		},
		{
			name: "Some negatives and positive",
			root: []interface{}{1,-2,-3,1,3,-2,nil,-1},
			// 1, -2, 1 3, -1
			targetSum: -4,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.root)
			got := hasPathSum(root, tt.targetSum)
			if got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}