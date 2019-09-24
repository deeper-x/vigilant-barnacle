package main

import (
	"fmt"

	p "github.com/deeper-x/at_roadstead_sw/parser"
)

func main() {
	pf := "./assets/all_data.csv"
	res := p.ScanDoc(pf)
	a := []string{"16", "20"}

	fmt.Println(p.Filter(res, a))
}
