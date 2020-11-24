#!/bin/python3
# ALDO FUSTER TURPIN

import math
import os
import random
import re
import sys


def createFrequencyMap(arr):
    freq = {}
    for num in arr:
        if num in freq:
            freq[num] += 1
        else:
            freq[num] = 1

    return freq


def pickingNumbers(a):
    freq = createFrequencyMap(a)
    items = sorted(freq.items())

    if len(items) == 1:
        return items[0][1]

    i = 1
    maxLength = 0

    while(i < len(items)):
        count1 = items[i][1]
        maxLength = max(maxLength, count1)
        
        count2 = items[i-1][1]
        maxLength = max(maxLength, count2)

        num1 = items[i][0]
        num2 = items[i-1][0]
        if abs(num1 - num2) <= 1:
            sum = count1 + count2
            maxLength = max(maxLength, sum)
            
        i += 1

    return maxLength


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input().strip())

    a = list(map(int, input().rstrip().split()))

    result = pickingNumbers(a)

    fptr.write(str(result) + '\n')

    fptr.close()
