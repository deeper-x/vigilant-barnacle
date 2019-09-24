package parser

import (
	"bufio"
	"log"
	"os"
	"strings"

	r "github.com/deeper-x/at_roadstead_sw/record"
)

var td r.TripData
var res = make(r.Res)

// ScannerDoc generic interface for scanner
type ScannerDoc interface {
	ScanDoc()
}

// ScanDoc Read document content
func ScanDoc() r.Res {
	f, err := os.Open("./assets/all_data.csv")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		ss := strings.Split(line, ",")

		t := td.GetTripID(ss)
		c := td.GetCounter(res, t)

		st := td.GetStates(res, t)
		st = append(st, td.GetStateID(ss))

		nc := c + 1

		res[t] = r.TripData{Tot: nc, States: st}
	}

	// for k, v := range res {
	// 	fmt.Println(k, v)
	// }

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return res
}
