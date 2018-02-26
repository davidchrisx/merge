package merge

import (
//	"sort"
	"os"
    "bufio"
)

const (
	maxMoveDist   = 0.2
	minMoveLength = 10
)

var (
	memo = []Shift{}
)

func min(a, b, c int) int {
	ret := a
	if b > ret {
		ret = b
	}
	if c > ret {
		ret = c
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewShift(a int, b string) Shift {
	return Shift{
		Index:  a,
		Change: b,
	}
}

func NewPair(a, b int) Pair {
	return Pair{
		a: a,
		b: b,
	}
}

func InsertChange(text []string, a int, b Pair) Change {
	return Change{
		ChangeType: "Insertion",
		text:       text,
		posA:       a,
		rangeB:     b,
	}
}
func DeleteChange(text []string, a Pair, b int) Change {
	return Change{
		ChangeType: "Deletion",
		text:       text,
		rangeA:     a,
		posB:       b,
	}
}

func MoveChange(texta []string, rangea Pair, posa int, textb []string, rangeb Pair, posb int) Change {
	return Change{
		ChangeType: "Move",
		textA:      texta,
		textB:      textb,
		posA:       posa,
		posB:       posb,
		rangeA:     rangea,
		rangeB:     rangeb,
	}
}

func InSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func minDiff(si, fi int) []LD {
	return []LD{}
}

func strDiff(a, b []string) []Change {
	//ls := len(a)
	//lf := len(b)
	memo = []Shift{}
	diff := minDiff(0, 0)
	//sort.Sort(diff)
	changes := []Change{}
	posDiff := 0
	offsetB := 0
	for posDiff < len(diff) {
		length := 0
		posAOld := diff[posDiff].Index
		for posDiff < len(diff) && diff[posDiff].Change == "i" {
			if diff[posDiff].Index != posAOld {
				break
			}
			length += 1
			posDiff += 1
		}
		if length > 0 {
			posA := posAOld
			rangeB0 := posAOld + offsetB
			rangeB1 := posAOld + offsetB + length
			changes = append(changes, InsertChange(b[rangeB0:rangeB1], posA, NewPair(rangeB0, rangeB1)))
			offsetB += length
		}
		if posDiff >= len(diff) {
			break
		}

		length = 0
		posAOld = diff[posDiff].Index
		for posDiff < len(diff) && diff[posDiff].Change == "d" {
			if diff[posDiff].Index != posAOld+length {
				break
			}
			length += 1
			posDiff += 1
		}
		if length > 0 {
			rangeA0 := posAOld
			rangeA1 := posAOld + length
			posB := posAOld + offsetB
			changes = append(changes, DeleteChange(a[rangeA0:rangeA1], NewPair(rangeA0, rangeA1), posB))
		}
	}
	return changes
}

func levDistance(a, b string) int {
	d := [100][100]int{}
	for i := 0; i <= len(a); i++ {
		d[i][0] = i
	}
	for j := 0; j <= len(b); j++ {
		d[0][j] = j
	}
	for j := 1; j <= len(b); j++ {
		for i := 1; i <= len(a); i++ {
			if a[i-1] == b[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i-1][j], d[i][j-1], d[i-1][j-1]) + 1
			}
		}
	}
	return d[len(a)][len(b)]
}

func findMove(diff []Change, Head bool) []Change {
	deleteLine := []int{}
	for i, _ := range diff {
		if diff[i].ChangeType == "Deletion" {
			for j, _ := range diff {
				if diff[j].ChangeType == "Insertion" {
					if !(InSlice(i, deleteLine)) && !(InSlice(j, deleteLine)) {
						lenI := 1
						lenJ := 1
						normalDist := 1
						if normalDist <= 2 && max(lenI, lenJ) >= minMoveLength {
							deleteLine = append(deleteLine, i)
							deleteLine = append(deleteLine, j)
							diff = append(diff, MoveChange(diff[i].text, diff[i].rangeA, diff[j].posA, diff[j].text, diff[j].rangeB, diff[i].posB))
						}
					}
				}
			}
		}
	}
	returnLine := []Change{}
	for i, d := range diff {
		if !InSlice(i, deleteLine) {
			returnLine = append(returnLine, d)
		}
	}
	return returnLine
}

func Merge(a, b, c *os.File) *os.File {
	aLine := toLine(a)
	bLine := toLine(b)
	abTable := InitLCSTable(aLine, bLine)
	abDiff := abTable.Diff()
	if Unchange(abDiff) {
		return c
	}
	return b
}

func Unchange(changes []Change) bool{
	for _, c := range changes{
		if c.ChangeType == "Insertion" {
			return false
		}
		if c.ChangeType == "Deletion" {
			return false
		}
	}
	return true
}

func toLine(f *os.File) []Line {
	line := []Line{}
	scanner := bufio.NewScanner(f)
	c := 0
    for scanner.Scan() {
    	line = append(line, Line{
    		Index: c,
    		Text: scanner.Text(),
    		}) 
    	c += 1
    }
    return line
}
