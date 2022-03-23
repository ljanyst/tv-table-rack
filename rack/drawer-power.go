// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	"github.com/ljanyst/ghostscad/lib/shapes"

	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

// Power drawer
type DrawerPower struct {
	Primitive Primitive
	Cfg       Config

	BaseScrew *Anchor
}

func NewDrawerPower(cfg Config) *DrawerPower {
	return &DrawerPower{
		Cfg: cfg,
	}
}

func (o *DrawerPower) Build() Primitive {
	base := NewDrawerBase(o.Cfg)
	base.Build()
	o.BaseScrew = base.BaseScrew

	lst := NewList(
		base.Primitive,
	)

	// Back-down
	bd := Vec2{
		base.Corners[3][0] - 10,
		base.Corners[3][1] - 10,
	}

	// Connector points to the board
	pts := []Vec2{
		{bd[0] - 65.5, bd[1] - 45},
		{bd[0], bd[1] - 45},
		{bd[0] - 65.5, bd[1]},
		bd,
	}

	for i, p := range pts {
		// Connect the board pin to the frame
		line := shapes.NewPolyline([]Vec2{base.Corners[i], p}, 4).SetRound(true)
		lst.Add(
			NewLinearExtrusion(
				o.Cfg.DrawerWidth,
				line.Build()))

		// Add the board pin
		lst.Add(
			NewTranslation(
				Vec3{p[0], p[1], (-o.Cfg.DrawerWidth + 10) / 2},
				NewCylinder(10, 2).SetFn(48)),
			NewTranslation(
				Vec3{p[0], p[1], (-o.Cfg.DrawerWidth + 14) / 2},
				NewCylinder(14, 1).SetFn(48)))
	}
	o.Primitive = lst
	return o.Primitive
}
