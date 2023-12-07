package d7

import (
	"bufio"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/itsbth/aoc-2023/runner"
)

type solver struct{}

var _ runner.Solver = solver{}

var (
	HAND_HIGH_CARD       handType = 0
	HAND_ONE_PAIR        handType = 1
	HAND_TWO_PAIR        handType = 2
	HAND_THREE_OF_A_KIND handType = 3
	HAND_FULL_HOUSE      handType = 4
	HAND_FOUR_OF_A_KIND  handType = 5
	HAND_FIVE_OF_A_KIND  handType = 6
)

type handType int

func (h handType) String() string {
	switch h {
	case HAND_HIGH_CARD:
		return "High Card"
	case HAND_ONE_PAIR:
		return "One Pair"
	case HAND_TWO_PAIR:
		return "Two Pair"
	case HAND_THREE_OF_A_KIND:
		return "Three of a Kind"
	case HAND_FULL_HOUSE:
		return "Full House"
	case HAND_FOUR_OF_A_KIND:
		return "Four of a Kind"
	case HAND_FIVE_OF_A_KIND:
		return "Five of a Kind"
	default:
		return "Unknown"
	}
}

// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
var cardValues = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

type hand struct {
	cards string
	value int
	// value part 2: J => 1
	value2 int
	best   handType
	best2  handType

	bid int
}

func calculateBest(counts map[int]int) handType {
	best := HAND_HIGH_CARD
	for _, count := range counts {
		if count == 2 {
			if best == HAND_ONE_PAIR {
				best = HAND_TWO_PAIR
			} else if best == HAND_THREE_OF_A_KIND {
				best = HAND_FULL_HOUSE
			} else {
				best = HAND_ONE_PAIR
			}
		} else if count == 3 {
			if best == HAND_ONE_PAIR {
				best = HAND_FULL_HOUSE
			} else {
				best = HAND_THREE_OF_A_KIND
			}
		} else if count == 4 {
			best = HAND_FOUR_OF_A_KIND
		} else if count == 5 {
			best = HAND_FIVE_OF_A_KIND
		}
	}
	return best
}

func NewHand(cards string, bid int) hand {
	value := 0
	value2 := 0
	for i := 0; i < len(cards); i += 1 {
		value *= 100
		value2 *= 100
		cardValue := cardValues[string(cards[i])]
		value += cardValue
		if cardValue == 11 {
			value2 += 1
		} else {
			value2 += cardValue
		}
	}
	counts := make(map[int]int)
	for i := 0; i < len(cards); i += 1 {
		counts[cardValues[string(cards[i])]] += 1
	}
	best := calculateBest(counts)
	jokers := counts[cardValues["J"]]
	counts[cardValues["J"]] = 0
	// try every card for J
	best2 := best
	// TODO: Should really only need to test cards already in hand
	for i := 2; i <= 14; i += 1 {
		counts[i] += jokers
		candidate := calculateBest(counts)
		if candidate > best2 {
			best2 = candidate
		}
		counts[i] -= jokers
	}

	return hand{
		cards:  cards,
		value:  value,
		value2: value2,
		best:   best,
		best2:  best2,
		bid:    bid,
	}
}

func (h *hand) Compare(other hand) int {
	if h.best > other.best {
		return 1
	} else if h.best < other.best {
		return -1
	} else {
		if h.value > other.value {
			return 1
		} else if h.value < other.value {
			return -1
		} else {
			return 0
		}
	}
}

func (h *hand) Compare2(other hand) int {
	if h.best2 > other.best2 {
		return 1
	} else if h.best2 < other.best2 {
		return -1
	} else {
		if h.value2 > other.value2 {
			return 1
		} else if h.value2 < other.value2 {
			return -1
		} else {
			return 0
		}
	}
}

func (solver) Solve(input io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(input)
	var hands []hand
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, 0, err
		}
		hands = append(hands, NewHand(parts[0], bid))
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Compare(hands[j]) < 0
	})
	// sum of hand * rank
	part1 := 0
	for i, hand := range hands {
		log.Printf("%d: %s", i, hand.cards)
		part1 += hand.bid * (i + 1)
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Compare2(hands[j]) < 0
	})
	log.Printf("hands: %+v", hands)
	part2 := 0
	for i, hand := range hands {
		log.Printf("%d: %s", i, hand.cards)
		part2 += hand.bid * (i + 1)
	}
	return part1, part2, nil
}

func init() {
	runner.Register(2023, 7, solver{})
}
