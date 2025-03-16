
import os, sys

with open("file.txt", 'r') as f:
    content = f.read()
    print(content)

with open("file.txt", 'w') as f:
    for line in f:
        f.write(line)
        