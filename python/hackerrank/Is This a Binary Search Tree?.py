""" Node is defined as
class node:
  def __init__(self, data):
      self.data = data
      self.left = None
      self.right = None
"""
MAX = 10001
MIN = -1
    
def check_binary_search_tree_(root):
    return check(root, MIN, MAX)

def check(node, minn, maxx):
    if node is None:
        return True
    
    if node.data < minn or node.data > maxx:
        return False
    
    return check(node.left, minn, node.data-1) and check(node.right, node.data+1, maxx)