package main

import (
	// "github.com/ljanyst/ghostscad/lib/shapes"
	"github.com/ljanyst/ghostscad/sys"

	. "github.com/ljanyst/tv-table-rack/rack"

	// . "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

func main() {
	cfg := NewDefaultConfig()
	rBlind := NewRightConnector(cfg)
	lBlind := NewLeftConnector(cfg)
	column := NewColumn(cfg)

	sys.RenderMultiple(map[string]Primitive{
		"right-blind-down": rBlind.Build(),
		"left-blind-down":  lBlind.Build(),
		"column":           column.Build(),
	}, "column")
}
