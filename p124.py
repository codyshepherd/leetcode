#Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def maxPathSum(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        # for each node get largest sum of children
        return max(self.largestSum(root))

    def largestSum(self, root):
        if root is None:
            return -float("inf"), -float("inf")

        if root.left is None and root.right is None:
            return root.val, root.val

        # if calling function cannot use this value in summation 
        # bc it comes from a child or a "crook", return false

        lu, ln = self.largestSum(root.left)
        ru, rn = self.largestSum(root.right)

        usable = max(root.val, lu+root.val, ru+root.val)
        nonusable = max(lu, ru, ln,rn,lu+ru+root.val)

        return usable, nonusable

        """
        m = 0
        usable = True

        if lusable and rusable:
            m = max(l, r, root.val, l+root.val, r+root.val, l+r+root.val)
            if m == l or m == r or m == l+r+root.val:
                usable = False
            return m, usable

        if lusable:
            m = max(l, r, root.val, l+root.val)
            if m == l or m == r:
                usable = False
            return m, usable

        if rusable:
            m = max(l, r, root.val, r+root.val)
            if m == l or m == r:
                usable = False
            return m, usable

        m = max(l,r,root.val)
        if m == l or m == r:
            return m, False
        return m, True
        """





root = TreeNode(-10)
root.left = TreeNode(9)
root.right = TreeNode(20)
root.right.left = TreeNode(15)
root.right.right = TreeNode(7)
sol = Solution()
print sol.maxPathSum(root)


root = TreeNode(1)
root.left = TreeNode(-2)
root.right = TreeNode(-3)
root.left.left = TreeNode(1)
root.left.left = TreeNode(3)
root.right.left = TreeNode(-2)
root.left.left.left = TreeNode(-1)
print sol.maxPathSum(root)