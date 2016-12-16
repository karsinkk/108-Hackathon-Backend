package main

import (
	"fmt"
	"sort"
)

func SortMapByValue(wordFrequencies map[int]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	return pl
}

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	a := map[int]int{1: 17, 30: 20, 32: 10}
	b := SortMapByValue(a)
	for _, v := range b {
		fmt.Println(v.Key, v.Value)
	}
}
