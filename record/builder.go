package record

import (
	"log"
	"strconv"
)

// Res defines results container
type Res map[int]TripData

// TripData contain each trip data
type TripData struct {
	Tot    int
	States []string
}

// GetTripID reads record trip id
func (TripData) GetTripID(record []string) int {
	n, err := strconv.Atoi(record[0])
	if err != nil {
		log.Fatal(err)
	}

	return n
}

// GetStateID reads record state id
func (TripData) GetStateID(record []string) string {
	return record[4]
}

// GetCounter read current totale for record
func (td TripData) GetCounter(res map[int]TripData, tripID int) int {
	return res[tripID].Tot
}

// GetStates return list with trip states
func (td TripData) GetStates(res map[int]TripData, tripID int) []string {
	return res[tripID].States
}
