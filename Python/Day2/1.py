
safe = 0

def validate_arr(arr) -> bool:
    prev = arr[0]
    check = None
    for i in arr[1:]:
        if abs(i - prev) > 3 or abs(i - prev) < 1:
            return False
        if check is None:
            check = lambda i, j: i < j if prev < i else lambda i, j: i > j
        if not check(prev, i):
            return False
    return True

with open("./Input/2.txt") as f:
    lines = f.readlines()
    for line in lines:
        split = line.split(" ")
        arr = [int(x) for x in split]
        print(safe)
        if validate_arr(arr):
            safe += 1
    print(safe)
        

