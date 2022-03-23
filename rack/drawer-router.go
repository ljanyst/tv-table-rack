// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	"github.com/ljanyst/ghostscad/lib/shapes"

	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

// Router drawer
type DrawerRouter struct {
	Primitive Primitive
	Cfg       Config

	BaseScrew *Anchor
}

func NewDrawerRouter(cfg Config) *DrawerRouter {
	return &DrawerRouter{
		Cfg: cfg,
	}
}

func (o *DrawerRouter) Build() Primitive {
	base := NewDrawerBase(o.Cfg)
	base.Build()
	o.BaseScrew = base.BaseScrew

	lst := NewList(
		base.Primitive,
	)

	// Connector points to the board
	pts := []Vec2{
		{-45, -35},
		{45, -35},
		{-45, 35},
		{45, 35},
	}

	for i, p := range pts {
		// Connect the board pin to the frame
		line := shapes.NewPolyline([]Vec2{base.Corners[i], p}, 15).SetRound(true)
		lst.Add(
			NewLinearExtrusion(
				o.Cfg.DrawerWidth,
				line.Build()))

		// Add the board pin
		lst.Add(
			NewDifference(
				NewTranslation(
					Vec3{p[0], p[1], (o.Cfg.DrawerWidth + 5) / 2},
					NewCylinder(5, 8).SetFn(48)),
				NewTranslation(
					Vec3{p[0], p[1], o.Cfg.DrawerWidth/2 + 5},
					NewCylinder(4, 6.5).SetFn(48))))
	}
	o.Primitive = lst
	return o.Primitive
}
