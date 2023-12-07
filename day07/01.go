package day07

import (
	"sort"
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

type Ordering int

const (
	Less Ordering = iota
	Equal
	Greater
)

type Card byte

var cardOrder = map[Card]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

func (c Card) Compare(other Card) Ordering {
	var diff = cardOrder[c] - cardOrder[other]
	if diff > 0 {
		return Greater
	} else if diff < 0 {
		return Less
	} else {
		return Equal
	}
}

const handSize = 5

type Hand struct {
	cards []Card
}

func NewHand(s string) Hand {
	var cards = make([]Card, handSize)
	for i, c := range s {
		cards[i] = Card(c)
	}
	return Hand{cards}
}

func makeHandSet(cards []Card) map[Card]int {
	var s = make(map[Card]int)
	for _, card := range cards {
		s[card]++
	}
	return s
}

func checkHandType(cards []Card, size, size1, size2 int) bool {
	var s = makeHandSet(cards)
	if len(s) != size {
		return false
	}
	for _, c := range s {
		if c != size1 && c != size2 {
			return false
		}
	}
	return true
}

func (h Hand) IsFiveOfAKind() bool {
	var s = makeHandSet(h.cards)
	return len(s) == 1
}

func (h Hand) IsFourOfAKind() bool {
	return checkHandType(h.cards, 2, 1, 4)
}

func (h Hand) IsFullHouse() bool {
	return checkHandType(h.cards, 2, 2, 3)
}

func (h Hand) IsThreeOfAKind() bool {
	return checkHandType(h.cards, 3, 1, 3)
}

func (h Hand) IsTwoPair() bool {
	return checkHandType(h.cards, 3, 1, 2)
}

func (h Hand) IsOnePair() bool {
	return checkHandType(h.cards, 4, 1, 2)
}

func (h Hand) IsHighCard() bool {
	var s = makeHandSet(h.cards)
	return len(s) == handSize
}

func (h Hand) Kind() int {
	if h.IsFiveOfAKind() {
		return 7
	} else if h.IsFourOfAKind() {
		return 6
	} else if h.IsFullHouse() {
		return 5
	} else if h.IsThreeOfAKind() {
		return 4
	} else if h.IsTwoPair() {
		return 3
	} else if h.IsOnePair() {
		return 2
	} else if h.IsHighCard() {
		return 1
	}
	return 0
}

func (h Hand) compareCardByCard(other Hand) Ordering {
	for i := 0; i < handSize; i++ {
		var ordering = h.cards[i].Compare(other.cards[i])
		if ordering != Equal {
			return ordering
		}
	}
	return Equal
}

func (h Hand) Compare(other Hand) Ordering {
	var diff = h.Kind() - other.Kind()
	if diff > 0 {
		return Greater
	} else if diff < 0 {
		return Less
	}
	return h.compareCardByCard(other)
}

type HandWithBid struct {
	hand Hand
	bid  int
}

type Game struct {
	hwb       []HandWithBid
	withJoker bool
}

func parseGame(s string, withJoker bool) Game {
	var lines = strings.Split(s, "\n")
	var hwb = make([]HandWithBid, len(lines))
	for i, line := range lines {
		var chunks = strings.Split(line, " ")
		var hand = NewHand(chunks[0])
		var bid = utils.Atoi(chunks[1])
		hwb[i] = HandWithBid{hand, bid}
	}
	return Game{hwb, withJoker}
}

func (g Game) Len() int {
	return len(g.hwb)
}

func (g Game) Less(i, j int) bool {
	var handi = g.hwb[i].hand
	var handj = g.hwb[j].hand
	if !g.withJoker {
		return handi.Compare(handj) == Less
	} else {
		return handi.Compare2(handj) == Less
	}
}

func (g Game) Swap(i, j int) {
	g.hwb[i], g.hwb[j] = g.hwb[j], g.hwb[i]
}

func (s Solver) Solve1(in string) interface{} {
	var game = parseGame(in, false)
	sort.Sort(game)
	var answer = 0
	for i, handWithBid := range game.hwb {
		answer += (i + 1) * handWithBid.bid
	}
	return answer
}
