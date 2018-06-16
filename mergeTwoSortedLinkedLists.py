# ALDO FUSTER TURPIN
"""
 Merge two linked lists
 head could be None as well for empty list
 Node is defined as
 
 class Node(object):
 
   def __init__(self, data=None, next_node=None):
       self.data = data
       self.next = next_node

 return back the head of the linked list in the below method.
"""

def MergeLists(headA, headB):
    newHead = None
    currentNode = None
    
    #easy end
    if headA is None:
        return headB
    
    #easy end
    elif headB is None:
        return headA
    
    #decide the new head
    if (headA and headB):
        if (headA.data <= headB.data):
            currentNode = headA
            headA = headA.next
            
        else:
            currentNode = headB
            headB = headB.next
    
    newHead = currentNode
    
    #create the new sorted linked list
    while((headA is not None) and (headB is not None)):
        if (headA.data <= headB.data):
            currentNode.next = headA
            currentNode = headA
            headA = headA.next
        else:
            currentNode.next = headB
            currentNode = headB
            headB = headB.next
    
    if (headA is None):
        currentNode.next = headB
    if (headB is None):
        currentNode.next = headA
    
    return newHead
