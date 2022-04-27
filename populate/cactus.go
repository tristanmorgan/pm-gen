package populate

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/chunk"
	"github.com/tristanmorgan/pm-gen/rand"
)

type Cactus struct {
	Amount int
}

var (
	sand   = block.Sand{}
	cactus = block.Cactus{}
)

func (t Cactus) Populate(w *world.World, pos world.ChunkPos, chunk *chunk.Chunk, r *rand.Random) {
	amount := r.Int31n(2) + int32(t.Amount)
	for i := int32(0); i < amount; i++ {
		x, z := int(r.Range(pos[0]*16, pos[0]*16+15)), int(r.Range(pos[1]*16, pos[1]*16+15))
		if y, ok := t.highestWorkableBlock(w, x, z); ok {
			cactusHeight := int(r.Int31n(3)) + 1
			for h := 0; h < cactusHeight; h++ {
				w.SetBlock(cube.Pos{x, y + h, z}, cactus, &world.SetOpts{DisableBlockUpdates: true, DisableLiquidDisplacement: true})
			}
		}
	}
}

func (t Cactus) highestWorkableBlock(w *world.World, x, z int) (int, bool) {
	next := w.Block(cube.Pos{x, 127, z})
	for y := 127; y >= 0; y-- {
		b := next
		next = w.Block(cube.Pos{x, y - 1, z})
		if b == air && next == sand {
			return y, true
		}
	}
	return 0, false
}
