package fixture

import "time"

type Fluctuation struct {
	Date       time.Time
	Sign       int
	Multiplier uint // in percents
}
