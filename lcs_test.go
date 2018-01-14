package merge

import (
	"fmt"
	"strings"
	"testing"
)

func diffToString(changes []Change) string {
	s := make([]string, len(changes))
	for i, c := range changes {
		pre := " "
		switch c.ChangeType {
		case "Insertion":
			pre = "+"
		case "Deletion":
			pre = "-"
		}
		s[i] = fmt.Sprintf("%s %s", pre, c.Text)
	}
	return strings.Join(s, "\n")
}

func splitLines(s string) (out []Line) {
	for i, l := range strings.Split(s, "\n") {
		out = append(out, Line{i, strings.TrimSpace(l)})
	}
	return
}

func TestLCSTable(t *testing.T) {
	cases := []struct {
		a, b            []Line
		expectedLengths []int
		expectedItems   string
	}{
		{
			splitLines("a\nx\nb"),
			splitLines("a\nb\nc"),
			[]int{0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 2, 2},
			`  a
- x
  b
+ c`,
		},
		{
			splitLines("g\na\nc"),
			splitLines("a\ng\nc\na\nt"),
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 1, 1, 2, 2, 2},
			`+ a
  g
+ c
  a
+ t
- c`,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			table := InitLCSTable(c.a, c.b)
			if len(table.lengths) != len(c.expectedLengths) {
				t.Fatalf("LCSTable lengths were not correct:\nexpected: %v\ngot:      %v", c.expectedLengths, table.lengths)
			}
			for i, l := range table.lengths {
				if l != c.expectedLengths[i] {
					t.Errorf("LCSTable lengths were not correct:\nexpected: %v\ngot:      %v", c.expectedLengths, table.lengths)
					break
				}
			}
			diff := table.Diff()
			diffStr := diffToString(diff)
			if diffStr != c.expectedItems {
				t.Fatalf("Diff was not correct:\nexpected:\n%v\ngot:\n%v", c.expectedItems, diffStr)
			}
		})
	}
}
