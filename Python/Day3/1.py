
text = ""
start = 0
total = 0
with open("./Input/3.txt") as f:
    text = f.read()

index = text.index("mul(")

def find_mul_content(start, text):
    pos = start
    content = ""
    while pos < len(text):
        if not text[pos].isdigit() and text[pos] != ',':
            break
        elif text[pos] != ")":
            content += text[pos]
            pos+=1
            continue
        else:
            break
    
    return (content, pos)

def process_content(content):
    split = content.split(",")
    if len(split) != 2:
        return 0
    if not split[0].isnumeric() or not split[1].isnumeric():
        return 0
    return int(split[0]) * int(split[1])

while start < len(text):
    print(f"Starting at {start}")
    print(f"Total at {total}")
    try:
        if start + 4 >= len(text):
            break
        
        content, start = find_mul_content(index + 4, text)
        total += process_content(content)
        index = text.index("mul(", start - 1)
    except:
        break



print(total)


