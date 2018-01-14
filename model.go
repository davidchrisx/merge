package merge

type Line struct {
	Index int
	Text  string
}

type Change struct {
	ChangeType string
	Line
}

type LCSTable struct {
	lengths []int
	a, b    []Line
}

type Changes []Change
