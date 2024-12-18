
safe = 0

def validate_arr(arr) -> bool:
    prev = arr[0]
    check = None
    for i in arr[1:]:
        if abs(i - prev) > 3 or abs(i - prev) < 1:
            return False
        if check is None:
            check = "increasing" if prev < i else "decreasing"
        if check == "increasing" and not prev < i:
            return False
        if check == "decreasing" and not prev > i:
            return False
        prev = i
    return True

with open("./Input/2.txt") as f:
    lines = f.readlines()
    for line in lines:
        split = line.split(" ")
        arr = [int(x) for x in split]
        if validate_arr(arr):
            safe += 1
    print(safe)
        

