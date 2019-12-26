package top
import (
	"strings"
	"fmt"
	"sort"
)
// Top count top 10 words
func Top(str string) PairList {
	words := strings.Fields(str)
	counts := make(map[string]int, len(words))
	for _, word := range words {
		_, inlist := counts[word]
		if inlist {
			counts[word] ++
		} else {
			counts[word] = 1
		}
	}
  var sorted PairList = sortkv(counts);
	for i, v := range sorted {
		fmt.Printf("index %d, value %d \n", i, v)
  }
  
	return sorted
}
type keyValue struct {
	Key   string
	Value int
}
// PairList dddd
type PairList []keyValue

func sortkv(list map[string]int) PairList {
	var sorted PairList
	for k, v := range list {
		sorted = append(sorted, keyValue{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})
	return sorted
}
