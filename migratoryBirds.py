# ALDO FUSTER TURPIN
#!/bin/python3


import os

def migratoryBirds(ar):
    my_map = {}
    for element in ar:
        # it is better to ask for forgiveness than permission
        try:
            my_map[element] += 1
        except KeyError:
            my_map[element] = 1
    # print("my_map: ", my_map)
    sorted_keys = sorted(my_map, key=my_map.get, reverse=True)
    #print(sorted_keys)
    return sorted_keys[0]
        

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    ar_count = int(input())

    ar = list(map(int, input().rstrip().split()))

    result = migratoryBirds(ar)

    fptr.write(str(result) + '\n')

    fptr.close()

