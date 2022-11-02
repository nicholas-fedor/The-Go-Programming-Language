// Package lenconv performs imperial and metric weight conversions.
package lenconv

import "fmt"

type Feet float64
type Meters float64

func (ft Feet) String() string       { return fmt.Sprintf("%gft", ft) }
func (meters Meters) String() string { return fmt.Sprintf("%gmeters", meters) }

// FtToMeters converts feet to meters.
func FtToMeters(ft Feet) Meters { return Meters(ft * 0.3048) }

// MetersToFt converts meters to feet.
func MetersToFt(meters Meters) Feet { return Feet(meters / 0.3048) }
