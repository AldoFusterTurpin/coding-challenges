#!/bin/python3
# ALDO FUSTER TURPIN

def rotate(n, d, a):
    if (d == 0):
        for i in a:
            print(i)

    # index of original_array
    i = 0
    final_array = [None] * n
    while (i < n):
        # we first do "i + d" to calculate the offset,
        # then we do "% d" to delete the "cicles" because sometimes is the same
        # rotate n positions than rotate m * n positions(is like do the same
        # rotation multiple times)
        index = i + d % n

        # now if we go furher than the end of the array we must recalculate
        # the new index (think like after the last item comes the first again)
        if (index >= n):
            index = index - n
        final_array[i] = a[index]
        i += 1

    [print(i, end=" ") for i in final_array]
    print()


if __name__ == '__main__':
    nd = input().split()

    # number of items in the array
    n = int(nd[0])

    # number of positions to rotate to left
    d = int(nd[1])

    # array of numbers
    a = list(map(int, input().rstrip().split()))

    rotate(n, d, a)
    
