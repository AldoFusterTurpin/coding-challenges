#!/bin/python3
# ALDO FUSTER TURPIN
import sys


def designerPdfViewer(h, word):
    # h is list of int's
    # word is string
    list_heights = []
    for i in word:
        index = ord(i) - 97
    #        print(index) # debug info
        letter_height = h[index]
        list_heights.append(letter_height)
    #    [print(i) for i in list_heights]# debug info
    max_height = max(list_heights)
    return len(word) * max_height


if __name__ == "__main__":
    h = list(map(int, input().strip().split(' ')))
    word = input().strip()
    result = designerPdfViewer(h, word)
    print(result)
    
