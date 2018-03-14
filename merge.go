package merge

import (
	"bufio"
	"os"
)

func Merge3(o, a, b *os.File) (string, bool) {
	ret := ""
	oLine := toLine(o)
	aLine := toLine(a)
	bLine := toLine(b)
	lines, conflict := Intersect3(oLine, aLine, bLine)
	for _, r := range lines {
		ret += r.Text + "\n"
	}
	return ret, conflict
}

//return all the line that intersect on the three texts.
func Intersect3(o, a, b []Line) ([]Line, bool) {
	aDiff := Diff(o, a)
	bDiff := Diff(o, b)
	conflict := false
	stable := Intersect2(aDiff, bDiff)

	oIntersect := Diff(o, stable)
	aIntersect := Diff(a, stable)
	bIntersect := Diff(b, stable)

	conflictO := FindConflict(oIntersect)
	conflictA := FindConflict(aIntersect)
	conflictB := FindConflict(bIntersect)
	ret := []Line{}
	for i := 0; i <= len(stable); i++ {
		ca := conflictA[i]
		cb := conflictB[i]
		co := conflictO[i]
		if isEqual(ca, cb) {
			ret = appendl(ret, ca)
		} else {
			if isEqual(co, ca) {
				ret = appendl(ret, cb)
			}
			if isEqual(co, cb) {
				ret = appendl(ret, ca)
			}
			if !isEqual(co, ca) && !isEqual(co, cb) {
				conflict = true
				ret = append(ret, Line{Text: "<<<<<<< A"})
				ret = appendl(ret, ca)
				ret = append(ret, Line{Text: "<<<<<<< B"})
				ret = appendl(ret, cb)
				ret = append(ret, Line{Text: "======="})
			}
		}

		if i < len(stable) {
			ret = append(ret, stable[i])
		}
	}

	return ret, conflict

}

func appendl(base, ext []Line) []Line {
	for _, x := range ext {
		base = append(base, x)
	}
	return base
}

func FindConflict(a []Change) []Lines {
	conflicts := []Lines{}
	temp := []Line{}
	for _, c := range a {
		if c.ChangeType == "Unchanged" {
			conflicts = append(conflicts, temp)
			temp = []Line{}
		} else {
			temp = append(temp, c.Line)
		}
	}
	conflicts = append(conflicts, temp)
	return conflicts
}

func Diff(a, b []Line) []Change {
	abTable := InitLCSTable(a, b)
	return abTable.Diff()
}

func Intersect2(a, b []Change) []Line {
	aIntersect := Intersect1(a)
	bIntersect := Intersect1(b)
	abDiff := Diff(aIntersect, bIntersect)
	return Intersect1(abDiff)
}

func Intersect1(a []Change) []Line {
	ret := []Line{}
	for _, C := range a {
		if C.ChangeType == "Unchanged" {
			ret = append(ret, C.Line)
		}
	}
	return ret
}

func Unchange(changes []Change) bool {
	for _, c := range changes {
		if c.ChangeType == "Insertion" {
			return false
		}
		if c.ChangeType == "Deletion" {
			return false
		}
	}
	return true
}

func isEqual(a, b []Line) bool {

	if len(a) != len(b) {
		return false
	}

	for in, _ := range a {
		if a[in].Text != b[in].Text {
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
			Text:  scanner.Text(),
		})
		c += 1
	}
	return line
}
