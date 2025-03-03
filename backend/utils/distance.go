package utils

import (
	"math"
)

/*Distance calculates the distance between two points on the Earth's surface given their latitude and longitude
does not take into account the altitude infrastructure*/
func GetDistanceBetweenTwoPointsOnEarth(lat1, lon1, lat2, lon2 float64) float64 {	
	if lat1 == lat2 && lon1 == lon2 {
		return 0
	}
	// Haversine formula
	toRadians := func(degree float64) float64 {
		return degree * math.Pi / 180
	}
	const R = 6371; // Earth's radius in kilometers
	lat1 = toRadians(lat1)
	lon1 = toRadians(lon1)
	lat2 = toRadians(lat2)
	lon2 = toRadians(lon2)
	dLat := lat2 - lat1
	dLon := lon2 - lon1
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * c
    return distance; // Distance in kilometers
}