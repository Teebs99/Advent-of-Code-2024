
safe = 0

def validate_arr(arr) -> bool:
    prev = arr[0]
    check = None
    for i in arr[1:]:
        if check is None:
            if prev < i:
                check = "increasing"
            if prev > i:
                check = "decreasing"
        if abs(i - prev) > 3 or abs(i - prev) < 1:
            return False
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
        result = validate_arr(arr)
        if not result:
            for i in range(len(arr)):
                result = validate_arr(arr[:i] + arr[i + 1:])
                if result:
                    break
        if result:
            safe += 1
    print(safe)
        

