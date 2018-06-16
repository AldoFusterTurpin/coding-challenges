#!/bin/python3
# ALDO FUSTER TURPIN

def findSuffix(queryString, count_dict):
    count = count_dict.get(queryString, 0)
    return count


if __name__ == '__main__':
    # number of strings that the collection will have
    strings_count = int(input())

    # the string collection
    strings = []

    # counter_dictionary
    count_dict = {}

    for _ in range(strings_count):
        strings_item = input()
        if (strings_item in count_dict):
            count_dict[strings_item] = count_dict[strings_item] + 1

        # first appearance of the element:
        else:
            count_dict[strings_item] = 1
        strings.append(strings_item)

    # print("The map:-> ", end="")
    # for k, v in count_dict.items():
    #     print("key:", k, ", value:", v)

    # number of querys that will be made
    q = int(input())

    for q_itr in range(q):
        # string that we want to look for
        queryString = input()

        # "res" is the number of times queryString appears in collections
        res = findSuffix(queryString, count_dict)
        print((str(res)))

