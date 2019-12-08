package fixture

import (
	"github.com/mixo/gorand"
	"github.com/mixo/godt"
	"time"
	"fmt"
	"math/rand"
)

type RowGenerator struct{}

func (this RowGenerator) Generate(rowCount int, date time.Time, valueDefinitions [][]interface{}) (rows [][]interface{}) {
	for i := 0; i < rowCount; i++ {
		rows = append(rows, this.generateRow(date, this.generateRandValues(valueDefinitions)))
	}
	return
}

func (this RowGenerator) generateRow(date time.Time, values []interface{}) []interface{} {
    row := []interface{}{godt.ToString(date)}
	return append(row, values...)
}

func (this RowGenerator) generateRandValues(valueDefinitions [][]interface{}) (values []interface{}) {
	var newValue interface{}
    for _, valueDefinition := range valueDefinitions {
        valueType := valueDefinition[0]
        rand.Seed(time.Now().UnixNano())
        switch (valueType) {
            case "int":
                newValue = gorand.GetIntRandBetween(valueDefinition[1].(int), valueDefinition[2].(int))
            case "float":
                newValue = gorand.GetFloatRandBetween(valueDefinition[1].(float64), valueDefinition[2].(float64))
			case "string":
				variants := valueDefinition[1:]
				variantIndex := gorand.GetIntRandBetween(0, len(variants) - 1)
				newValue = variants[variantIndex]
            default:
                panic(fmt.Sprintf("Undefined type %s", valueType))
        }
		values = append(values, newValue)
    }
    return
}
