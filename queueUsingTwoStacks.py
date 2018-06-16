#!/bin/python3
# ALDO FUSTER TURPIN


# move elements from "stackFrom" to "stackTo"
def moveFromTo(stackFrom, stackTo):
    while len(stackFrom) > 0:
        stackTo.append(stackFrom.pop())


if __name__ == '__main__':

    #when I say "queue" I refer to a imaginary queue where elements are in the correct order(appended chronologically)

    # where elements are pushed the first time a push occurs
    stack1 = []

    # temporary (for moving from stack1 to stack2 when we want to delete or print the element at the beginning
    #if elements are on stack2-> they are in reversed order(comparing to a chronological insert on a queue)
    stack2 = []

    # number of queries
    q = int(input())

    for i in range(q):
        line_list = [int(i) for i in input().split()]

        # print("the list: ", line_list) # debug info

        # Enqueue element into the end of the queue.
        if line_list[0] == 1:
            stack1.append(line_list[1])

        # Dequeue the element at the front of the queue.
        elif line_list[0] == 2:
            if len(stack1) == len(stack2) == 0:
                # do nothing
                pass

            # if elements are on the stack1 we must move it to stack2
            elif len(stack2) == 0:
                moveFromTo(stack1, stack2)

            stack2.pop()

        # Print the element at the front of the queue.
        elif line_list[0] == 3:
            if len(stack1) == len(stack2) == 0:
                # do nothing
                pass

            # if elements are on the stack1 we must move it to stack2
            elif len(stack2) == 0:
                moveFromTo(stack1, stack2)
            print(stack2[len(stack2) - 1])

        # debug info:
        # print("stack1:", stack1)
        # print("stack2:", stack2)
        # print("-" * 15)

