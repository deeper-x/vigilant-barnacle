package main

import (
	p "github.com/deeper-x/at_roadstead_sw/parser"
)

func main() {
	res := p.ScanDoc()
	a := []string{"16", "20"}

	p.Filter(res, a)
}
