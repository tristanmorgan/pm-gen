package populate

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/chunk"
	"github.com/tristanmorgan/pm-gen/rand"
)

var (
	air   = block.Air{}
	water = block.Water{}
)

type Populator interface {
	Populate(w *world.World, pos world.ChunkPos, chunk *chunk.Chunk, r *rand.Random)
}
