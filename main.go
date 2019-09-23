package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	r "github.com/deeper-x/at_roadstead_sw/record"
)

var td r.TripData
var res = make(r.Res)

func main() {

	f, err := os.Open("assets/all_data.csv")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		ss := strings.Split(line, ",")

		// res[10] = r.TripData{Tot: 10, Valid: true}
		// res[11] = r.TripData{Tot: 20, Valid: false}
		trip := td.GetTripID(ss)
		tot := td.GetCounter(res, trip)

		res[trip] = r.TripData{Tot: td.GetCounter(res, trip)}

		nv := tot + 1

		res[trip] = r.TripData{Tot: nv}
		for _, v := range res {
			fmt.Println(v)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
