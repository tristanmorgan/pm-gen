package biome

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/tristanmorgan/pm-gen/populate"
)

type Ocean struct{}

func (o Ocean) Populators() []populate.Populator {
	return []populate.Populator{populate.Kelp{Amount: 15}}
}

func (o Ocean) ID() uint8 {
	return IDOcean
}

func (o Ocean) Elevation() (min, max int) {
	return 46, 62
}

func (o Ocean) GroundCover() []world.Block {
	return []world.Block{
		block.Sand{},
		block.Sand{},
		block.Gravel{},
		block.Gravel{},
		block.Gravel{},
	}
}

func (o Ocean) Temperature() float64 {
	return 0.4
}

func (o Ocean) Rainfall() float64 {
	return 0.5
}
