package main

import (
	"fmt"

	p "github.com/deeper-x/at_roadstead_sw/parser"
)

func main() {
	res := p.ScanDoc()

	fmt.Println(res)
}
