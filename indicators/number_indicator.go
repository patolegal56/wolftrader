package indicators

import (
	"cex-bot/types"
)

type NumberIndicatorConfig struct {
	Number float64
}

type NumberIndicator struct {
	Config NumberIndicatorConfig
}

func (r *NumberIndicator) Calculate(input []*types.Candle) []float64 {
	result := make([]float64, 0)

	for i := 0; i < len(input); i++ {
		result = append(result, r.Config.Number)
	}

	return result
}