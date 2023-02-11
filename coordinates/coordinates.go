package coordinates

import (
	"math"

	"github.com/umahmood/haversine"
)

func RoundTo1Decimal(coords haversine.Coord) haversine.Coord {
	return haversine.Coord{
		lat: math.Round(coords.Lat*10) / 10,
		lon: math.Round(coords.Lon*10) / 10,
	}
}

func RoundTo2Decimals(coords haversine.Coord) haversine.Coord {
	return haversine.Coord{
		lat: math.Round(coords.Lat*100) / 100,
		lon: math.Round(coords.Lon*100) / 100,
	}
}

func RoundTo3Decimals(coords haversine.Coord) haversine.Coord {
	return haversine.Coord{
		lat: math.Round(coords.Lat*1000) / 1000,
		lon: math.Round(coords.Lon*1000) / 1000,
	}
}

func RoundTo4Decimals(coords haversine.Coord) haversine.Coord {
	return haversine.Coord{
		lat: math.Round(coords.Lat*10000) / 10000,
		lon: math.Round(coords.Lon*10000) / 10000,
	}
}
