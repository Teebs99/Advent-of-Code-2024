import re

text = ""
total = 0
with open("./Input/3.txt") as f:
    text = f.read()

occurances = re.findall("mul\([0-9]+,[0-9]+\)", text)

for time in occurances:
    time = time.replace("mul(", "")
    time = time.replace(")", "")
    split = time.split(",")
    total += (int(split[0]) * int(split[1]))

print(total)