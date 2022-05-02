package populate

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/chunk"
	"github.com/tristanmorgan/pm-gen/rand"
)

type DeadBush struct {
	Amount int
}

var (
	deadBush = block.DeadBush{}
)

func (t DeadBush) Populate(w *world.World, pos world.ChunkPos, chunk *chunk.Chunk, r *rand.Random) {
	amount := r.Int31n(2) + int32(t.Amount)
	for i := int32(0); i < amount; i++ {
		x, z := int(r.Range(pos[0]*16, pos[0]*16+15)), int(r.Range(pos[1]*16, pos[1]*16+15))
		if y, ok := t.highestWorkableBlock(w, x, z); ok {
			w.SetBlock(cube.Pos{x, y, z}, deadBush, &world.SetOpts{DisableBlockUpdates: true, DisableLiquidDisplacement: true})
		}
	}
}

func (t DeadBush) highestWorkableBlock(w *world.World, x, z int) (int, bool) {
	next := w.Block(cube.Pos{x, 127, z})
	for y := 127; y >= 0; y-- {
		b := next
		next = w.Block(cube.Pos{x, y - 1, z})
		if b == air && supportsVegetation(deadBush, next) {
			return y, true
		}
	}
	return 0, false
}
