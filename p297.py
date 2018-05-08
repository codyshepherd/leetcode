# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        if root is None:
            return "[]"

        final = ""
        nulls = ""

        q = list()
        q.append(root) #enqueue

        while len(q) > 0:
            item = q[0]
            q = q[1:]

            if item is None:
                nulls += "null,"
            else:
                q.append(item.left)
                q.append(item.right)

                final += nulls + str(item.val) + ','
                nulls = ""

        final = final[:-1]
        final = "[" + final + "]"
        return final
        

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        if data == "[]"
            return None

        lst = data.strip('[]')
        lst = lst.split(',')
        root = TreeNode(lst[0])
        q = [root]
        lst = lst[1:]
        lastlevel = list()
        i = 0
        while i < len(lst):
            node = q[0]
            node.left = TreeNode(lst[i]) if lst[i] != "null" else None
            i += 1
            if i < len(lst):
                node.right = TreeNode(lst[i]) if lst[i] != "null" else None
                i += 1
            if node.left is not None:
                q.append(node.left)
            if node.right is not None:
                q.append(node.right)
            q = q[1:]

        return root

# Your Codec object will be instantiated and called as such:
codec = Codec()
print codec.serialize(codec.deserialize("[5,4,7,3,null,2,null,-1,null,9]"))
'''
root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
root.right.left = TreeNode(4)
root.right.right = TreeNode(5)
print codec.serialize(root)
print codec.serialize(codec.deserialize(codec.serialize(root)))
'''