#!/bin/python3
# ALDO FUSTER TURPIN

import math
import os
import random
import re
import sys

# constraints:
# 1 <= n <= 10^3
# 1 <= |s| <= 10^3 where |s| is the length of the sequence.


def complementary(myCharacter):
    # print("myCharacter: {}".format(myCharacter)) # debug info
    if myCharacter == "]":
        return "["

    elif myCharacter == "}":
        return "{"

    elif myCharacter == ")":
        return "("
    else:
        raise BaseException("Not supported")



def isBalanced(s):
    if len(s) % 2 != 0 :
        return "NO"

    else:
        # stack to save opening brackets
        stack = []
        for element in s:
            # print("element: {}".format(element)) # debug info

            if element == "(" or element == "[" or element == "{":
                stack.append(element)
            else:
                # element is "]" or ")" or "}"

                # if stack is empty then closing bracket does not match any
                # opening bracket
                if len(stack) == 0:
                    return "NO"

                elif stack.pop() != complementary(element):
                    return "NO"

        if len(stack) == 0:
            return "YES"

        else:
            return "NO"
    
    
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    t = int(input())

    for t_itr in range(t):
        s = input()

        result = isBalanced(s)

        fptr.write(result + '\n')

    fptr.close()

