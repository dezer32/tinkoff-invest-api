package helpers

import (
	"math"

	"github.com/dezer32/tinkoff-invest-api/pkg/generated/investapi"
)

func ConvertQuotation(quotation *investapi.Quotation) float64 {
	if quotation == nil {
		return float64(0)
	}
	if quotation.Nano <= 0 {
		return float64(quotation.Units)
	}

	lenNano := int(math.Ceil(math.Log10(float64(quotation.Nano))))
	lenZero := int64(math.Pow10(lenNano))

	return float64(quotation.Units*lenZero+int64(quotation.Nano)) / float64(lenZero)
}
