package merge

type Line struct {
	Index int
	Text  string
}

type LD struct{
	Index int
	Change string
}

type Change struct {
	ChangeType string
	Line
	text	[]string
	textA	[]string
	textB	[]string
	posA	int
	posB	int
	rangeA	Pair
	rangeB	Pair
}

type Pair struct{
	a,b int
}

type Shift struct {
	Index int 
	Change string
}

type LCSTable struct {
	lengths []int
	a, b    []Line
}

type Changes []Change
