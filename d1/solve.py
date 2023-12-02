import re

# digits only
P1 = re.compile(r'(\d)')

# digits or spelled out numbers
# use lookahead to match overlapping numbers
P2 = re.compile(r'(?=(\d|one|two|three|four|five|six|seven|eight|nine))')

VALUE = {
    'one': 1,
    'two': 2,
    'three': 3,
    'four': 4,
    'five': 5,
    'six': 6,
    'seven': 7,
    'eight': 8,
    'nine': 9
}

def expand(s):
    if s in VALUE:
        return VALUE[s]
    else:
        return int(s)

def solve(s):
    s1 = 0
    s2 = 0
    for line in s.splitlines():
        do = P1.findall(line)
        ds = P2.findall(line)
        # debug: print line s1 (first and last digit) and s2 (first and last digit/number)
        # print(f"{line} {do[0]}{do[-1]} {expand(ds[0])}{expand(ds[-1])}")
        s1 += int(do[0]) * 10 + int(do[-1])
        s2 += expand(ds[0]) * 10 + expand(ds[-1])
    print(f"Part 1: {s1}")
    print(f"Part 2: {s2}")


if __name__ == '__main__':
    f = open('INPUT', 'r')
    solve(f.read())
    f.close()