import re

text=""
start = 0
total = 0
index = 0
with open("./Input/3.txt") as f:
    text = f.read()

do_str = ""

while True:
    try:
        index = text.index("don't()", index)
        do_str += text[start:index]
        start = text.index("do()", index)
        index += 4
    except:
        break

times = re.findall("mul\([0-9]+,[0-9]+\)", do_str)

for time in times:
    time = time.replace("mul(", "")
    time = time.replace(")", "")
    split = time.split(",")
    total += (int(split[0]) * int(split[1]))

print(total)