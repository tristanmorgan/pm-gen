package populate

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/chunk"
	"github.com/tristanmorgan/pm-gen/rand"
)

type Kelp struct {
	Amount int
}

func (t Kelp) Populate(w *world.World, pos world.ChunkPos, chunk *chunk.Chunk, r *rand.Random) {
	amount := r.Int31n(2) + int32(t.Amount)
	for i := int32(0); i < amount; i++ {
		x, z := int(r.Range(pos[0]*16, pos[0]*16+15)), int(r.Range(pos[1]*16, pos[1]*16+15))
		if y, ok := t.highestWorkableBlock(w, x, z); ok {
			kelpHeight := int(r.Int31n(3)) + 1
			for h := 0; h < kelpHeight; h++ {
				w.SetBlock(cube.Pos{x, y + h, z}, block.Kelp{Age: int(r.Int31n(15)) + 5}, &world.SetOpts{DisableBlockUpdates: true, DisableLiquidDisplacement: false})
			}
		}
	}
}

func (t Kelp) highestWorkableBlock(w *world.World, x, z int) (int, bool) {
	var next world.Block
	for y := 61; y >= 42; y-- {
		next = w.Block(cube.Pos{x, y - 1, z})
		if next == sand {
			return y, true
		}
	}
	return 0, false
}
