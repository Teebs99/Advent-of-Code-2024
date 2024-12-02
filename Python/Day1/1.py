import os

left, right, raw = [], [], []

with open("./Python/Day1/1-input.txt", "r") as f:
    raw = f.readlines()

for line in raw:
    split_str = line.split("   ")
    left.append(int(split_str[0]))
    right.append(int(split_str[1]))

left.sort()
right.sort()

total = 0
for i in range(len(left)):
    total += abs(left[i] - right[i])

print(total)