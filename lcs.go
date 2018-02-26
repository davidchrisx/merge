package merge

// InitLCSTable precompute len(a) * len(b)
// lengths of array for LCS
func InitLCSTable(a, b []Line) *LCSTable {
	table := &LCSTable{
		lengths: make([]int, (len(a)+1)*(len(b)+1)),
		a:       a,
		b:       b,
	}

	for i, _ := range a {
		for j, _ := range b {
			k := (i+1)*(len(b)+1) + (j + 1)
			if a[i].Text == b[j].Text {
				table.lengths[k] = table.getLength(i, j) + 1
			} else {
				nextA := table.getLength(i+1, j)
				nextB := table.getLength(i, j+1)
				if nextA > nextB {
					table.lengths[k] = nextA
				} else {
					table.lengths[k] = nextB
				}
			}
		}
	}
	return table
}

// getLength return index for LCSTable from
// index ai and index bi
func (t *LCSTable) getLength(ai, bi int) int {
	return t.lengths[ai*(len(t.b)+1)+bi]
}

// Diff returns a diff of the two sets of lines the LCSTable was created with,
// as determined by the LCS.
func (t *LCSTable) Diff() []Change {
	return t.recursiveDiff(len(t.a), len(t.b))
}

func (t *LCSTable) recursiveDiff(i, j int) []Change {
	if i == 0 && j == 0 {
		return nil
	}

	var toAdd Change
	if i == 0 {
		toAdd.ChangeType = "Insertion"
		toAdd.Line = t.b[j-1]
		j--
	} else if j == 0 {
		toAdd.ChangeType = "Deletion"
		toAdd.Line = t.a[i-1]
		i--
	} else if t.a[i-1].Text == t.b[j-1].Text {
		toAdd.ChangeType = "Unchanged"
		toAdd.Line = t.a[i-1]
		i--
		j--
	} else if t.getLength(i, j-1) > t.getLength(i-1, j) {
		toAdd.ChangeType = "Insertion"
		toAdd.Line = t.b[j-1]
		j--
	} else {
		toAdd.ChangeType = "Deletion"
		toAdd.Line = t.a[i-1]
		i--
	}

	return append(t.recursiveDiff(i, j), toAdd)
}
