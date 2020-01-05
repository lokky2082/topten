package top

import (
	"testing"
)

func TestTop(t *testing.T) {
	top10 := "cat and dog one dog two cats and one man"
	top10correct := map[string]int{
		"and":  2,
		"dog":  2,
		"one":  2,
		"two":  1,
		"cat":  1,
		"cats": 1,
		"man":  1,
	}
	list := Top(top10)
	if list != top10correct {
		t.Errorf("List was incorrect, got: %s, want: %s.", list, top10correct)
	}
}
