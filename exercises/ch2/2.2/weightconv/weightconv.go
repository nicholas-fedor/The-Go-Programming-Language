// Package weightconv performs imperial and metric weight conversions.
package weightconv

import "fmt"

type Pounds float64
type Kilograms float64

func (lbs Pounds) String() string    { return fmt.Sprintf("%glbs", lbs) }
func (kgs Kilograms) String() string { return fmt.Sprintf("%gkg", kgs) }

// LbsToKgs converts pounds to kilograms.
func LbsToKgs(lbs Pounds) Kilograms { return Kilograms(lbs / 2.2046) }

// KgsToLbs converts kilograms to pounds.
func KgsToLbs(kgs Kilograms) Pounds { return Pounds(kgs * 2.2046) }
