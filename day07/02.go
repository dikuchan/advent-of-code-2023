package day07

import (
	"sort"
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

const joker = Card('J')

var cardOrder2 = map[Card]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}

func (c Card) Compare2(other Card) Ordering {
	var diff = cardOrder2[c] - cardOrder2[other]
	if diff > 0 {
		return Greater
	} else if diff < 0 {
		return Less
	} else {
		return Equal
	}
}

func (h Hand) compareCardByCard2(other Hand) Ordering {
	for i := 0; i < handSize; i++ {
		var ordering = h.cards[i].Compare2(other.cards[i])
		if ordering != Equal {
			return ordering
		}
	}
	return Equal
}

func replaceCard(hand Hand, old, new Card) Hand {
	var replaced = strings.ReplaceAll(string(hand.cards), string(new), string(old))
	return Hand{[]Card(replaced)}
}

func (h Hand) Kind2() int {
	var kinds = make([]int, 0)
	for card := range makeHandSet(h.cards) {
		var kind = replaceCard(h, card, joker).Kind()
		kinds = append(kinds, kind)
	}
	return utils.FindMax(kinds)
}

func (h Hand) Compare2(other Hand) Ordering {
	var diff = h.Kind2() - other.Kind2()
	if diff > 0 {
		return Greater
	} else if diff < 0 {
		return Less
	}
	return h.compareCardByCard2(other)
}

func (s Solver) Solve2(in string) interface{} {
	var game = parseGame(in, true)
	sort.Sort(game)
	var answer = 0
	for i, handWithBid := range game.hwb {
		answer += (i + 1) * handWithBid.bid
	}
	return answer
}
