package main

import (
	"bufio"
	"fmt"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type line struct {
	hand string
	bid  int
}

func main() {
	content, err := os.Open("7.in")
	if err != nil {
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	var lines []line

	for scanner.Scan() {
		text := scanner.Text()
		var line line
		line.hand = strings.Split(text, " ")[0]
		line.bid, _ = strconv.Atoi(strings.Split(text, " ")[1])
		lines = append(lines, line)
	}
	sort.Slice(lines, func(i, j int) bool {
		if getHand(lines[i].hand) != getHand(lines[j].hand) {
			return getHand(lines[i].hand) < getHand(lines[j].hand)
		}
		return hand2pts(lines[i].hand) < hand2pts(lines[j].hand)
	})
	total := 0
	for i, j := range lines {
		total += day7(j.bid, i+1)
	}
	fmt.Println(total)
}

func day7(bid int, rank int) int {
	return bid * rank
}

func card2pts(card byte) int {
	if card == 'A' {
		return 14
	}
	if card == 'K' {
		return 13
	}
	if card == 'Q' {
		return 12
	}
	if card == 'J' {
		//return 11
		return 0
	}
	if card == 'T' {
		return 10
	}
	return int(card - '0')
}

func hand2pts(hand string) int {
	return ((((card2pts(hand[0]))*15+card2pts(hand[1]))*15+card2pts(hand[2]))*15+card2pts(hand[3]))*15 + card2pts(hand[4])
}

func sortHand(hand string) string {
	slice := []byte(hand)
	sort.Slice(slice, func(i, j int) bool {
		return card2pts(slice[i]) < card2pts(slice[j])
	})
	return string(slice)
}

func getHand(hand string) int {
	hand = replaceJoker(hand)
	hand = sortHand(hand)
	five := pcre.MustCompile(`([TJQKA\d])\1{4}`, 0).MatcherString(hand, 0)
	four := pcre.MustCompile(`([TJQKA\d])\1{3}`, 0).MatcherString(hand, 0)
	three := pcre.MustCompile(`([TJQKA\d])\1{2}`, 0).MatcherString(hand, 0)
	two := pcre.MustCompile(`([TJQKA\d])\1.*([TJQKA\d])\2`, 0).MatcherString(hand, 0)
	one := pcre.MustCompile(`([TJQKA\d])\1`, 0).MatcherString(hand, 0)
	if five.Matches() {
		return 6
	}

	if four.Matches() {
		return 5
	}

	if three.Matches() {
		if two.Matches() {
			return 4
		}
		return 3
	}

	if two.Matches() {
		return 2
	}

	if one.Matches() {
		return 1
	}

	return 0

}

func replaceJoker(hand string) string {
	count := make(map[rune]int)
	most := 0
	common := 'J'
	for _, j := range hand {
		if j != 'J' {
			count[j] += 1
			if count[j] > most {
				most = count[j]
				common = j
			}
		}
	}
	re := regexp.MustCompile("J")
	hand = re.ReplaceAllString(hand, string(common))
	return hand
}
