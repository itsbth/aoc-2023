words = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

values = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

# generate a go struct that matches all the words
# methods: reset(), match(byte) (int, bool)
#          resetBackwards(), matchBackwards(byte) (int, bool)
# struct contains one uint for each word
# reset sets them all to 0
# resetBW sets them all to len(word)
# match is a giant switch statement that increments/resets counters. likewise matchBW decrements/resets counters


def print_case(ch, backwards=False):
    mod = "--" if backwards else "++"
    for word in words:
        idx = word.find(ch)
        # assume at most 2 occurrences of ch in word
        lidx = word.rfind(ch)
        default = len(word) - 1 if backwards else 0
        if idx == 0 and not backwards:
            default = 1
        if lidx == len(word) - 1 and backwards:
            default = len(word) - 2
        if idx == -1:
            print(f"        m.{word} = {default}")
        elif lidx != idx:
            print(
                f"        if m.{word} == {idx} || m.{word} == {lidx} {{ m.{word}{mod} }} else {{ m.{word} = {default} }}"
            )
        else:
            print(
                f"        if m.{word} == {idx} {{ m.{word}{mod} }} else {{ m.{word} = {default} }}"
            )

print("package main")
print()
print("type matcher struct {")
print(f'    {", ".join(words)} int')
print("}")
print()
print("func (m *matcher) reset() {")
for word in words:
    print(f"    m.{word} = 0")
print("}")
print()
print("func (m *matcher) resetBW() {")
for word in words:
    print(f"    m.{word} = {len(word) - 1}")
print("}")
print()
print("func (m *matcher) match(b byte) (int, bool) {")
print("    switch b {")
all_chars = sorted(set("".join(words)))
for ch in all_chars:
    print(f"    case '{ch}':")
    print_case(ch)
print("    default:")
for word in words:
    print(f"        m.{word} = 0")
print("    }")
# todo: integrate this into the switch
for word in words:
    print(f"    if m.{word} == {len(word)} {{ return {values[word]}, true }}")
print("    return 0, false")
print("}")
print()
print("func (m *matcher) matchBW(b byte) (int, bool) {")
print("    switch b {")
for ch in all_chars:
    print(f"    case '{ch}':")
    print_case(ch, backwards=True)
print("    default:")
for word in words:
    print(f"        m.{word} = {len(word) - 1}")
print("    }")
for word in words:
    print(f"    if m.{word} == {-1} {{ return {values[word]}, true }}")
print("    return 0, false")
print("}")
