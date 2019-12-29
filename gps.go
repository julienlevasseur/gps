package gps

import (
	"errors"
	"fmt"

	gpsd "github.com/stratoberry/go-gpsd"
)

// Speed is provided by TPVs
// Lat is provided by TPV
// Lon is provided by TPV
// Heading is provided by ATT

// TPV Time Position Velocity
var TPV *gpsd.TPVReport

// Init initialize the gps.
func Init() {
	var gps *gpsd.Session
	var err error

	if gps, err = gpsd.Dial(gpsd.DefaultAddress); err != nil {
		panic(fmt.Sprintf("Failed to connect to GPSD: %s", err))
	}

	gps.AddFilter("TPV", func(r interface{}) {
		TPV = r.(*gpsd.TPVReport)
	})

	done := gps.Watch()
	<-done
}

// TPV return the TPV (Time, Position, Velocity) object.
func Tpv() (tpv *gpsd.TPVReport, err error) {
	if TPV != nil {
		tpv.Speed = TPV.Speed
		tpv.Lat = TPV.Lat
		tpv.Lon = TPV.Lon
		tpv.Time = TPV.Time
		tpv.Alt = TPV.Alt
		tpv.Climb = TPV.Climb
		return tpv, nil
	}

	return nil, errors.New("TPVReport is not ready yet")
}

// Speed return the current gps speed.
func Speed() (speed float64) {
	if TPV != nil {
		return TPV.Speed
	}
}
//
//// Coordinates return the current gps coordinates.
//func Coordinates() (lat, lon float64, err error) {
//	if TPV != nil {
//		return TPV.Lat, TPV.Lon, nil
//	}
//
//	return 0, 0, errors.New("Coordinates not aquired yet.\n")
//}