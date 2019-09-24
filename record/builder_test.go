package record

import (
	"fmt"
	"testing"
)

var l = []string{"3195", "KOZNITZA", "13925", "DA ORMEGGIO A ORMEGGIO", "18", "2018-01-05 22:15"}
var ms = []string{"10", "16"}
var o = TripData{Tot: 10, States: ms, ShipName: "SHIP1"}
var m = map[int]TripData{}

func TestGetTripID(t *testing.T) {
	res := o.GetTripID(l)
	if res < 0 {
		t.Fail()
	}
}

func TestGetShipName(t *testing.T) {
	res := o.GetShipName(l)
	if len(res) == 0 {
		t.Fail()
	}
}

func TestGetStateID(t *testing.T) {
	res := o.GetStateID(l)
	if len(res) == 0 {
		t.Fail()
	}
}

func TestGetCounter(t *testing.T) {
	m[3195] = o
	res := o.GetCounter(m, 3195)

	if res < 0 {
		t.Fail()
	}
}

func TestGetStates(t *testing.T) {
	m[10] = o
	res := o.GetStates(m, 3195)
	fmt.Println(res)
	if len(res) == 0 {
		t.Fail()
	}
}
