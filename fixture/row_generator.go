package fixture

import (
	"github.com/mixo/godt"
	"time"
)

type RowGenerator struct{}

func (this RowGenerator) Generate(rowCount int, date time.Time) (rows [][]interface{}) {
	for i := 0; i < rowCount; i++ {
		rows = append(rows, this.generateRow(date))
	}
	return
}

func (this RowGenerator) generateRow(date time.Time) []interface{} {
	return []interface{}{godt.ToString(date)}
}
