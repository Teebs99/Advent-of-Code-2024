
left, right, raw = [], [], []

with open("./Python/Day1/1-input.txt") as f:
    raw = f.readlines()

count = {}
for line in raw:
    split_str = line.split("   ")
    left.append(int(split_str[0]))
    r_num = int(split_str[1])
    if r_num in count:
        count[r_num] += 1
    else:
        count[r_num] = 1

total = 0
for num in left:
    multiple = count[num] if num in count else 0
    total += (num * multiple)

print(total)