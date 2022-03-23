// Copyright 2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package rack

import (
	"github.com/ljanyst/ghostscad/lib/shapes"

	. "github.com/go-gl/mathgl/mgl64"
	. "github.com/ljanyst/ghostscad/primitive"
)

// RaspberryPi drawer
type DrawerRPi struct {
	Primitive Primitive
	Cfg       Config
	Module    bool

	BaseScrew *Anchor
}

func NewDrawerRPi(cfg Config, module bool) *DrawerRPi {
	return &DrawerRPi{
		Cfg:    cfg,
		Module: module,
	}
}

func (o *DrawerRPi) Build() Primitive {
	base := NewDrawerBase(o.Cfg)
	base.Build()
	o.BaseScrew = base.BaseScrew

	sum := NewUnion()
	sum.Add(base.Primitive)
	sub := NewList()

	buildLine := func(pts []Vec2) {
		line := shapes.NewPolyline(pts, 6).SetRound(true)
		sum.Add(
			NewLinearExtrusion(
				o.Cfg.DrawerWidth,
				line.Build()))
	}

	buildPts := func(pts []Vec2, corn Vec2) {
		for i, p := range pts {
			if corn[0] != p[0] && corn[1] != p[1] {
				buildLine([]Vec2{base.Corners[i], p})
			}

			// Add the board pin
			sum.Add(
				NewTranslation(
					Vec3{p[0], p[1], (-o.Cfg.DrawerWidth + 5) / 2},
					NewCylinder(5, 2.5).SetFn(48)),
				NewTranslation(
					Vec3{p[0], p[1], 0},
					NewCylinder(o.Cfg.DrawerWidth, 3.5).SetFn(48)))

			// Pin holes
			sub.Add(
				NewTranslation(
					Vec3{p[0], p[1], 0},
					NewCylinder(14, 1).SetFn(48)),
				NewTranslation(
					Vec3{p[0], p[1], (-o.Cfg.DrawerWidth + 1.8) / 2},
					NewCylinder(2, 2.75).SetFn(48)))
		}
	}

	// RPi
	fu := base.Corners[0]
	pts := []Vec2{
		fu,
		{fu[0] + 48.5, fu[1]},
		{fu[0], fu[1] + 57.5},
		{fu[0] + 48.5, fu[1] + 57.5},
	}
	buildPts(pts, fu)
	buildLine([]Vec2{pts[1], pts[3]})
	buildLine([]Vec2{pts[3], pts[2]})

	// module
	if o.Module {
		bb := base.Corners[3]
		pts = []Vec2{
			{bb[0] - 22.5, bb[1] - 29.5},
			{bb[0] - 22.5, bb[1]},
			{bb[0], bb[1] - 29.5},
			bb,
		}
		buildPts(pts, bb)
		buildLine([]Vec2{pts[0], pts[1]})
		buildLine([]Vec2{pts[0], pts[2]})
	}

	o.Primitive = NewDifference(sum, sub)
	return o.Primitive
}
