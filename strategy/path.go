package strategy

import (
	"cex-bot/types"
)

type Path struct {
	Tiles []Tile
}

func (p *Path) HasSignal(candleCollection *types.CandleCollection, symbol types.Symbol) (bool, error) {
	for _, tile := range p.Tiles {
		signal, err := tile.HasSignal(candleCollection, symbol)
		if err != nil {
			return false, err
		}
		if !signal {
			return false, nil
		}
	}

	return true, nil
}
