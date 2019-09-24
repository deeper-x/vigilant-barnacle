package parser

import (
	"testing"
)

func TestFilter(t *testing.T) {
	res := ScanDoc("../assets/all_data.csv")
	in := []string{"16", "20"}

	_, tot := Filter(res, in)

	if tot == 0 {
		t.Fail()
	}
}
