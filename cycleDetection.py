# ALDO FUSTER TURPIN
"""
Detect a cycle in a linked list. Note that the head pointer may be 'None' if the list is empty.

A Node is defined as: 
 
    class Node(object):
        def __init__(self, data = None, next_node = None):
            self.data = data
            self.next = next_node
"""


def has_cycle(head):
    if (head is None):
        return False
    else:
        set_references = set()
        current_node = head
        while (current_node.next is not None):
            if (current_node in set_references):
                return True
            else:
                set_references.add(current_node)
            current_node = current_node.next
        return False
    
