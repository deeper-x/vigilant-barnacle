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

// Filter stream data with maneuvering criteria
func Filter(in r.Res, sc []string) ([]string, int) {
	ret := make([]string, 0)
	counter := 0
	tot := 0
	for _, v := range in {
		// states := v.States
		// fmt.Println(sc, v.States)
		for _, w := range v.States {
			for _, z := range sc {
				if z == w {
					counter++
					if counter > 1 {
						ret = append(ret, v.ShipName)
						tot++
					}
				}
			}
		}
		counter = 0
	}
	//fmt.Println("TOT:", tot)
	return ret, tot
}

// ScanDoc Read document content
func ScanDoc(pf string) r.Res {
	f, err := os.Open(pf)

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
		sn := td.GetShipName(ss)

		res[t] = r.TripData{Tot: nc, States: st, ShipName: sn}
	}

	// for k, v := range res {
	// 	fmt.Println(k, v)
	// }

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return res
}
