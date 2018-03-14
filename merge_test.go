package merge

import (
	"os"
	"testing"
)

func TestMerge(t *testing.T) {
	expectedString := "tambahheader\n1\n2\n3\n4\ntambaha\n5\ntambahb\n6\n"
	t.Run("testMerge", func(t *testing.T) {
		a, _ := os.Open("a.txt")
		o, _ := os.Open("o.txt")
		b, _ := os.Open("b.txt")
		ret, _ := Merge3(o, a, b)
		if expectedString != ret {
			t.Fatal(ret)
		}
	})
}
