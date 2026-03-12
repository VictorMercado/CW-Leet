https://leetcode.com/problems/path-sum
# 112. Path Sum

## Problem Description
Given the `root` of a binary tree and an integer `targetSum`, return `true` if the tree has a **root-to-leaf** path such that adding up all the values along the path equals `targetSum`.

> **Note:** A leaf is a node with no children.

---

## Examples

**Example 1:**
![Tree Example 1](https://assets.leetcode.com/uploads/2021/01/18/pathsum1.jpg)
* **Input:** root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
* **Output:** true
* **Explanation:** The root-to-leaf path with the target sum is `5 -> 4 -> 11 -> 2`, which equals 22.

**Example 2:**
![Tree Example 2](https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg)
* **Input:** root = [1,2,3], targetSum = 5
* **Output:** false
* **Explanation:** * Path 1: `1 -> 2` (Sum = 3)
  * Path 2: `1 -> 3` (Sum = 4)
  * No path equals 5.

**Example 3:**
* **Input:** root = [], targetSum = 0
* **Output:** false
* **Explanation:** Since the tree is empty, there are no root-to-leaf paths.

---

## Constraints
* The number of nodes in the tree is in the range `[0, 5000]`.
* `-1000 <= Node.val <= 1000`
* `-1000 <= targetSum <= 1000`

---
