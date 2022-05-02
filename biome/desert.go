package biome

import (
	"github.com/tristanmorgan/pm-gen/populate"
)

type Desert struct {
	sandy
}

func (d Desert) Populators() []populate.Populator {
	return []populate.Populator{populate.Cactus{Amount: 1}, populate.DeadBush{Amount: 2}}
}

func (d Desert) ID() uint8 {
	return IDDesert
}

func (d Desert) Elevation() (min, max int) {
	return 63, 74
}

func (d Desert) Temperature() float64 {
	return 2.0
}

func (d Desert) Rainfall() float64 {
	return 0.0
}
