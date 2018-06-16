# ALDO FUSTER TURPIN

# For your reference:
#
# SinglyLinkedListNode:
#     int data
#     SinglyLinkedListNode next


#   Input:
#   3   n
#   16  a number of the linked list
#   13  a number of the linked list
#   7   a number of the linked list
#   1   new number to insert
#   2   position where to insert the new number
#


def insertNodeAtPosition(head, data, position):
    new_node = SinglyLinkedListNode(data)
    
    #if empty list
    if (head is None):
        head = new_node
        
    else:
        current_node = head
        i = 0
        while (i < position-1):
            current_node = current_node.next
            i += 1
        aux_node = current_node.next
        current_node.next = new_node
        new_node.next = aux_node
     
    return head
        

