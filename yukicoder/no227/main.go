package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type pokerHand int

const (
	fullHouse pokerHand = iota // ある数をちょうど3つと、別の数をちょうど2つ含む
	threeCard                  // ある数をちょうど3つ含む
	twoPair                    // ある数をちょうど2つと、別の数をちょうど2つ含む
	onePair                    // ある数をちょうど2つ含む
	noHand                     // 役無し
)

func (ph pokerHand) String() string {
	switch ph {
	case fullHouse:
		return "FULL HOUSE"
	case threeCard:
		return "THREE CARD"
	case twoPair:
		return "TWO PAIR"
	case onePair:
		return "ONE PAIR"
	case noHand:
		return "NO HAND"
	default:
		return "UNKOWN"
	}
}

func scanToSliceInt(in *os.File) ([]int, error) {
	var scanner = bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	var res []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}

	return res, scanner.Err()
}

func main() {

	myCards, err := scanToSliceInt(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println("%v", myCards)
}
