package fixture

import (
	"github.com/mixo/gorand"
	"math"
	"time"
)

type RowCountGenerator struct{}

func (this RowCountGenerator) generate(
	rowCountFrom, rowCountTo int,
	dateTime time.Time,
	fluctuations []Fluctuation) int {

	fluctuation, isFluctuationFound := this.findFluctuation(dateTime, fluctuations)
	if isFluctuationFound {
		return this.getFluctuationRowCount(fluctuation, rowCountFrom, rowCountTo)
	}

	return gorand.GetRandBetween(rowCountFrom, rowCountTo)
}

func (this RowCountGenerator) findFluctuation(
	dateTime time.Time,
	fluctuations []Fluctuation) (Fluctuation, bool) {
	for _, fluctuation := range fluctuations {
		if fluctuation.Date.Truncate(time.Hour * 24).Equal(dateTime.Truncate(time.Hour * 24)) {
			return fluctuation, true
		}
	}

	return Fluctuation{}, false
}

func (this RowCountGenerator) getFluctuationRowCount(
	fluctuation Fluctuation,
	rowCountFrom, rowCountTo int) int {

	rowCountBase := rowCountFrom
	if fluctuation.Sign > 0 {
		rowCountBase = rowCountTo
	}

	return int(math.Abs(float64(fluctuation.Sign*rowCountBase))) + fluctuation.Sign*int(fluctuation.Multiplier)
}
