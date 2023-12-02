package main

type matcher struct {
    one, two, three, four, five, six, seven, eight, nine int
}

func (m *matcher) reset() {
    m.one = 0
    m.two = 0
    m.three = 0
    m.four = 0
    m.five = 0
    m.six = 0
    m.seven = 0
    m.eight = 0
    m.nine = 0
}

func (m *matcher) resetBW() {
    m.one = 2
    m.two = 2
    m.three = 4
    m.four = 3
    m.five = 3
    m.six = 2
    m.seven = 4
    m.eight = 4
    m.nine = 3
}

func (m *matcher) match(b byte) (int, bool) {
    switch b {
    case 'e':
        if m.one == 2 { m.one++ } else { m.one = 0 }
        m.two = 0
        if m.three == 3 || m.three == 4 { m.three++ } else { m.three = 0 }
        m.four = 0
        if m.five == 3 { m.five++ } else { m.five = 0 }
        m.six = 0
        if m.seven == 1 || m.seven == 3 { m.seven++ } else { m.seven = 0 }
        if m.eight == 0 { m.eight++ } else { m.eight = 1 }
        if m.nine == 3 { m.nine++ } else { m.nine = 0 }
    case 'f':
        m.one = 0
        m.two = 0
        m.three = 0
        if m.four == 0 { m.four++ } else { m.four = 1 }
        if m.five == 0 { m.five++ } else { m.five = 1 }
        m.six = 0
        m.seven = 0
        m.eight = 0
        m.nine = 0
    case 'g':
        m.one = 0
        m.two = 0
        m.three = 0
        m.four = 0
        m.five = 0
        m.six = 0
        m.seven = 0
        if m.eight == 2 { m.eight++ } else { m.eight = 0 }
        m.nine = 0
    case 'h':
        m.one = 0
        m.two = 0
        if m.three == 1 { m.three++ } else { m.three = 0 }
        m.four = 0
        m.five = 0
        m.six = 0
        m.seven = 0
        if m.eight == 3 { m.eight++ } else { m.eight = 0 }
        m.nine = 0
    case 'i':
        m.one = 0
        m.two = 0
        m.three = 0
        m.four = 0
        if m.five == 1 { m.five++ } else { m.five = 0 }
        if m.six == 1 { m.six++ } else { m.six = 0 }
        m.seven = 0
        if m.eight == 1 { m.eight++ } else { m.eight = 0 }
        if m.nine == 1 { m.nine++ } else { m.nine = 0 }
    case 'n':
        if m.one == 1 { m.one++ } else { m.one = 0 }
        m.two = 0
        m.three = 0
        m.four = 0
        m.five = 0
        m.six = 0
        if m.seven == 4 { m.seven++ } else { m.seven = 0 }
        m.eight = 0
        if m.nine == 0 || m.nine == 2 { m.nine++ } else { m.nine = 1 }
    case 'o':
        if m.one == 0 { m.one++ } else { m.one = 1 }
        if m.two == 2 { m.two++ } else { m.two = 0 }
        m.three = 0
        if m.four == 1 { m.four++ } else { m.four = 0 }
        m.five = 0
        m.six = 0
        m.seven = 0
        m.eight = 0
        m.nine = 0
    case 'r':
        m.one = 0
        m.two = 0
        if m.three == 2 { m.three++ } else { m.three = 0 }
        if m.four == 3 { m.four++ } else { m.four = 0 }
        m.five = 0
        m.six = 0
        m.seven = 0
        m.eight = 0
        m.nine = 0
    case 's':
        m.one = 0
        m.two = 0
        m.three = 0
        m.four = 0
        m.five = 0
        if m.six == 0 { m.six++ } else { m.six = 1 }
        if m.seven == 0 { m.seven++ } else { m.seven = 1 }
        m.eight = 0
        m.nine = 0
    case 't':
        m.one = 0
        if m.two == 0 { m.two++ } else { m.two = 1 }
        if m.three == 0 { m.three++ } else { m.three = 1 }
        m.four = 0
        m.five = 0
        m.six = 0
        m.seven = 0
        if m.eight == 4 { m.eight++ } else { m.eight = 0 }
        m.nine = 0
    case 'u':
        m.one = 0
        m.two = 0
        m.three = 0
        if m.four == 2 { m.four++ } else { m.four = 0 }
        m.five = 0
        m.six = 0
        m.seven = 0
        m.eight = 0
        m.nine = 0
    case 'v':
        m.one = 0
        m.two = 0
        m.three = 0
        m.four = 0
        if m.five == 2 { m.five++ } else { m.five = 0 }
        m.six = 0
        if m.seven == 2 { m.seven++ } else { m.seven = 0 }
        m.eight = 0
        m.nine = 0
    case 'w':
        m.one = 0
        if m.two == 1 { m.two++ } else { m.two = 0 }
        m.three = 0
        m.four = 0
        m.five = 0
        m.six = 0
        m.seven = 0
        m.eight = 0
        m.nine = 0
    case 'x':
        m.one = 0
        m.two = 0
        m.three = 0
        m.four = 0
        m.five = 0
        if m.six == 2 { m.six++ } else { m.six = 0 }
        m.seven = 0
        m.eight = 0
        m.nine = 0
    default:
        m.one = 0
        m.two = 0
        m.three = 0
        m.four = 0
        m.five = 0
        m.six = 0
        m.seven = 0
        m.eight = 0
        m.nine = 0
    }
    if m.one == 3 { return 1, true }
    if m.two == 3 { return 2, true }
    if m.three == 5 { return 3, true }
    if m.four == 4 { return 4, true }
    if m.five == 4 { return 5, true }
    if m.six == 3 { return 6, true }
    if m.seven == 5 { return 7, true }
    if m.eight == 5 { return 8, true }
    if m.nine == 4 { return 9, true }
    return 0, false
}

func (m *matcher) matchBW(b byte) (int, bool) {
    switch b {
    case 'e':
        if m.one == 2 { m.one-- } else { m.one = 1 }
        m.two = 2
        if m.three == 3 || m.three == 4 { m.three-- } else { m.three = 3 }
        m.four = 3
        if m.five == 3 { m.five-- } else { m.five = 2 }
        m.six = 2
        if m.seven == 1 || m.seven == 3 { m.seven-- } else { m.seven = 4 }
        if m.eight == 0 { m.eight-- } else { m.eight = 4 }
        if m.nine == 3 { m.nine-- } else { m.nine = 2 }
    case 'f':
        m.one = 2
        m.two = 2
        m.three = 4
        if m.four == 0 { m.four-- } else { m.four = 3 }
        if m.five == 0 { m.five-- } else { m.five = 3 }
        m.six = 2
        m.seven = 4
        m.eight = 4
        m.nine = 3
    case 'g':
        m.one = 2
        m.two = 2
        m.three = 4
        m.four = 3
        m.five = 3
        m.six = 2
        m.seven = 4
        if m.eight == 2 { m.eight-- } else { m.eight = 4 }
        m.nine = 3
    case 'h':
        m.one = 2
        m.two = 2
        if m.three == 1 { m.three-- } else { m.three = 4 }
        m.four = 3
        m.five = 3
        m.six = 2
        m.seven = 4
        if m.eight == 3 { m.eight-- } else { m.eight = 4 }
        m.nine = 3
    case 'i':
        m.one = 2
        m.two = 2
        m.three = 4
        m.four = 3
        if m.five == 1 { m.five-- } else { m.five = 3 }
        if m.six == 1 { m.six-- } else { m.six = 2 }
        m.seven = 4
        if m.eight == 1 { m.eight-- } else { m.eight = 4 }
        if m.nine == 1 { m.nine-- } else { m.nine = 3 }
    case 'n':
        if m.one == 1 { m.one-- } else { m.one = 2 }
        m.two = 2
        m.three = 4
        m.four = 3
        m.five = 3
        m.six = 2
        if m.seven == 4 { m.seven-- } else { m.seven = 3 }
        m.eight = 4
        if m.nine == 0 || m.nine == 2 { m.nine-- } else { m.nine = 3 }
    case 'o':
        if m.one == 0 { m.one-- } else { m.one = 2 }
        if m.two == 2 { m.two-- } else { m.two = 1 }
        m.three = 4
        if m.four == 1 { m.four-- } else { m.four = 3 }
        m.five = 3
        m.six = 2
        m.seven = 4
        m.eight = 4
        m.nine = 3
    case 'r':
        m.one = 2
        m.two = 2
        if m.three == 2 { m.three-- } else { m.three = 4 }
        if m.four == 3 { m.four-- } else { m.four = 2 }
        m.five = 3
        m.six = 2
        m.seven = 4
        m.eight = 4
        m.nine = 3
    case 's':
        m.one = 2
        m.two = 2
        m.three = 4
        m.four = 3
        m.five = 3
        if m.six == 0 { m.six-- } else { m.six = 2 }
        if m.seven == 0 { m.seven-- } else { m.seven = 4 }
        m.eight = 4
        m.nine = 3
    case 't':
        m.one = 2
        if m.two == 0 { m.two-- } else { m.two = 2 }
        if m.three == 0 { m.three-- } else { m.three = 4 }
        m.four = 3
        m.five = 3
        m.six = 2
        m.seven = 4
        if m.eight == 4 { m.eight-- } else { m.eight = 3 }
        m.nine = 3
    case 'u':
        m.one = 2
        m.two = 2
        m.three = 4
        if m.four == 2 { m.four-- } else { m.four = 3 }
        m.five = 3
        m.six = 2
        m.seven = 4
        m.eight = 4
        m.nine = 3
    case 'v':
        m.one = 2
        m.two = 2
        m.three = 4
        m.four = 3
        if m.five == 2 { m.five-- } else { m.five = 3 }
        m.six = 2
        if m.seven == 2 { m.seven-- } else { m.seven = 4 }
        m.eight = 4
        m.nine = 3
    case 'w':
        m.one = 2
        if m.two == 1 { m.two-- } else { m.two = 2 }
        m.three = 4
        m.four = 3
        m.five = 3
        m.six = 2
        m.seven = 4
        m.eight = 4
        m.nine = 3
    case 'x':
        m.one = 2
        m.two = 2
        m.three = 4
        m.four = 3
        m.five = 3
        if m.six == 2 { m.six-- } else { m.six = 1 }
        m.seven = 4
        m.eight = 4
        m.nine = 3
    default:
        m.one = 2
        m.two = 2
        m.three = 4
        m.four = 3
        m.five = 3
        m.six = 2
        m.seven = 4
        m.eight = 4
        m.nine = 3
    }
    if m.one == -1 { return 1, true }
    if m.two == -1 { return 2, true }
    if m.three == -1 { return 3, true }
    if m.four == -1 { return 4, true }
    if m.five == -1 { return 5, true }
    if m.six == -1 { return 6, true }
    if m.seven == -1 { return 7, true }
    if m.eight == -1 { return 8, true }
    if m.nine == -1 { return 9, true }
    return 0, false
}
